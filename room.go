package main

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Room struct {
	RoomID string
	Title  string
	Desc   string
	Config *RoomConfig

	isLive    bool
	rdBuf     []byte
	rawBuf    []byte
	conn      *net.TCPConn
	heartbeat bool // heartbeat 是否在发送心跳包
}

func (r *Room) Init(id string, config json.RawMessage) error {
	r.RoomID = id
	r.Config = &RoomConfig{}
	return r.Config.Load(config)
}

const (
	ProtocolPlain = iota
	ProtocolHeartbeat
	ProtocolDeflate
	ProtocolBrotli
)

// Read data在下次调用Read前有效
func (r *Room) Read() (proto uint16, act uint32, data []byte, err error) {
reParse:
	data, err = r.read(4)
	if err != nil {
		panic(err)
	}
	packLen := binary.BigEndian.Uint32(data)
	data, err = r.read(int(packLen - 4))
	if err != nil {
		panic(err)
	}
	if binary.BigEndian.Uint16(data) != 16 {
		panic("packet length != 16?")
	}
	proto = binary.BigEndian.Uint16(data[2:])
	switch proto {
	case ProtocolDeflate:
		var zr io.ReadCloser
		zr, err = zlib.NewReader(bytes.NewReader(data[12:]))
		r.rdBuf, err = io.ReadAll(zr)
		if err != nil {
			return 0, 0, nil, err
		}
		goto reParse
	case ProtocolBrotli:
		panic("br is not implemented")
	}
	act = binary.BigEndian.Uint32(data[4:])
	data = data[12:]
	return
}

func (r *Room) read(size int) (buf []byte, err error) {
	if len(r.rdBuf) >= size {
		buf = r.rdBuf[:size]
		r.rdBuf = r.rdBuf[size:]
		return
	}

	if cap(r.rawBuf) < size {
		DebugPrint(size, "> buffer, consider increase buffer")
		buf = make([]byte, size)
		copy(buf, r.rdBuf)
		_, err = io.ReadFull(r.conn, buf[len(r.rdBuf):])
		if err != nil {
			return nil, err
		}
		r.rdBuf = r.rawBuf[:0]
		return
	}
	var n int
	copy(r.rawBuf, r.rdBuf)
	n, err = io.ReadAtLeast(r.conn, r.rawBuf[len(r.rdBuf):], size-len(r.rdBuf))
	if err != nil {
		return nil, err
	}
	buf = r.rawBuf[:size]
	n += len(r.rdBuf)
	r.rdBuf = r.rawBuf[size:n]
	return
}

var dataHeartBeat = []byte{0x00, 0x00, 0x00, 0x10,
	0x00, 0x10,
	0x00, 0x01,
	0x00, 0x00, 0x00, 0x02,
	0x00, 0x00, 0x00, 0x01}

func (r *Room) Heartbeat() {
	time.AfterFunc(45*time.Second, r.Heartbeat)
	if r.conn != nil {
		_, err := r.conn.Write(dataHeartBeat)
		if err != nil {
			r.conn.Close()
			return
		}
	}
}

func (r *Room) Record() {
	var err error
fErr:
	for r.isLive {
		if err != nil {
			time.Sleep(5 * time.Second)
		}
		resp := getAPI("/getRoomPlayInfo", "room_id=%s&protocol=0,1&format=0,1,2&codec=0,1&qn=10000&platform=web&ptype=8", r.RoomID)
		var info RoomPlayInfo
		err = json.Unmarshal(resp, &info)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, stream := range info.PlayurlInfo.Playurl.Stream {
			for _, format := range stream.Format {
				for _, codec := range format.Codec {
					if codec.CurrentQn == 10000 {
						for _, host := range codec.URLInfo {
							resp, err := http.Get(host.Host + codec.BaseURL + host.Extra)
							if err != nil {
								log.Println(err)
								continue
							}
							f, err := os.OpenFile(r.RecordName(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
							if err != nil {
								log.Println(err)
								continue fErr
							}
							_, err = io.Copy(f, resp.Body)
							if err != nil {
								log.Println(err)
								continue fErr
							}
						}
					}
				}
			}
		}
	}
}

func (r *Room) RefreshInfo() {

}

func (r *Room) RecordName() string {
	r.RefreshInfo()
	return fmt.Sprintf("%s/录制-%s-%s-%s.flv", r.RoomID, r.RoomID, time.Now().Format("2006-01-02-15-04-05"), r.Title)
}

func (r *Room) Connect() {
	goto start
retry:
	time.Sleep(2 * time.Second)
start:
	if r.conn != nil {
		r.conn.Close()
		r.conn = nil
	}
	resp := getAPI("/xlive/web-room/v1/index/getDanmuInfo", "id=%s&type=0", r.RoomID)
	var info DanmuInfo
	err := json.Unmarshal(resp, &info)
	if err != nil {
		log.Println("解析弹幕姬信息出错", err)
		goto retry
	}

	const bufSize = 8192
	if r.rawBuf == nil {
		r.rawBuf = make([]byte, 16, bufSize)
	}
	r.rawBuf[5] = 0x10
	r.rawBuf[7] = 0x01
	r.rawBuf[11] = 0x07
	/*	copy(r.rawBuf, []byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x10,
		0x00, 0x01,
		0x00, 0x00, 0x00, 0x07})*/
	r.rawBuf = append(r.rawBuf, `{"uid":0,"protover":0,"platform":"web","clientver":"2.6.25","type":2,"key":`...)
	r.rawBuf = append(r.rawBuf, info.Token...)
	r.rawBuf = append(r.rawBuf, `,"roomid":`...)
	r.rawBuf = append(r.rawBuf, r.RoomID...)
	r.rawBuf = append(r.rawBuf, '}')

	binary.BigEndian.PutUint32(r.rawBuf, uint32(len(r.rawBuf)))

	for _, host := range info.HostList {
		addrs, err := net.DefaultResolver.LookupIPAddr(context.Background(), host.Host)
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			r.conn, err = net.DialTCP("tcp", nil, &net.TCPAddr{IP: addr.IP, Port: 2243})
			if err != nil {
				log.Println("连接弹幕姬失败", err)
				continue
			}
			r.conn.Write(r.rawBuf)
			r.rawBuf = r.rawBuf[:cap(r.rawBuf)]
			proto, _, _, err := r.Read()
			if err != nil {
				log.Println("与弹幕姬交流失败")
				continue
			}
			if proto != ProtocolHeartbeat {
				panic("?")
			}
			r.Heartbeat()
			for {
				const (
					ActHeartBeatResp = 3
				)
				_, act, data, err := r.Read()
				if act == ActHeartBeatResp {
					continue
				}
				if err == nil {
					fmt.Printf("%x %s\n", act, string(data))
					if len(data) > 8 {
						fmt.Printf("%x\n", binary.LittleEndian.Uint32(data[8:]))
						switch binary.LittleEndian.Uint32(data[8:]) {
						case binary.LittleEndian.Uint32([]byte("LIVE")):
							println("LIVE!")
							r.isLive = true
							r.Record()
						case binary.LittleEndian.Uint32([]byte("PREP")):
							println("STOP!")
							r.isLive = false
						}
					}
				}
			}
		}
	}
	log.Println("无可用弹幕姬")
	goto retry
}
