package main

import "encoding/json"

type (
	BiliResp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		//TTL     int    `json:"ttl"`
		Data json.RawMessage `json:"data"`
	}

	// UserInfo /live_user/v1/UserInfo/get_anchor_in_room?roomid=%d
	UserInfo struct {
		Info struct {
			//UID   int                 `json:"uid"`
			Uname json.RawMessage `json:"uname"`
			Face  string          `json:"face"`
			//Rank              string `json:"rank"`
			//PlatformUserLevel int    `json:"platform_user_level"`
			//MobileVerify      int    `json:"mobile_verify"`
			//Identification    int    `json:"identification"`
			//OfficialVerify    struct {
			//	Type int    `json:"type"`
			//	Desc string `json:"desc"`
			//	Role int    `json:"role"`
			//} `json:"official_verify"`
			//VipType int `json:"vip_type"`
			//Gender  int `json:"gender"`
		} `json:"info"`
		//Level struct {
		//	UID         int    `json:"uid"`
		//	Cost        int    `json:"cost"`
		//	Rcost       int    `json:"rcost"`
		//	UserScore   string `json:"user_score"`
		//	Vip         int    `json:"vip"`
		//	VipTime     string `json:"vip_time"`
		//	Svip        int    `json:"svip"`
		//	SvipTime    string `json:"svip_time"`
		//	UpdateTime  string `json:"update_time"`
		//	MasterLevel struct {
		//		Level            int    `json:"level"`
		//		Color            int    `json:"color"`
		//		Current          []int  `json:"current"`
		//		Next             []int  `json:"next"`
		//		AnchorScore      int    `json:"anchor_score"`
		//		UpgradeScore     int    `json:"upgrade_score"`
		//		MasterLevelColor int    `json:"master_level_color"`
		//		Sort             string `json:"sort"`
		//	} `json:"master_level"`
		//	UserLevel   int `json:"user_level"`
		//	Color       int `json:"color"`
		//	AnchorScore int `json:"anchor_score"`
		//} `json:"level"`
		//San int `json:"san"`
	}

	// RoomInfo /room/v1/Room/get_info?RoomID=%d
	RoomInfo struct {
		UID    int `json:"uid"`
		RoomID int `json:"room_id"`
		//ShortID          int      `json:"short_id"`
		//Attention        int      `json:"attention"`
		//Online int `json:"online"`
		//IsPortrait bool `json:"is_portrait"`
		//Description      string   `json:"description"`
		LiveStatus int `json:"live_status"`
		//AreaID           int      `json:"area_id"`
		//ParentAreaID     int      `json:"parent_area_id"`
		//OldAreaID        int      `json:"old_area_id"`
		//Background string `json:"background"`
		Title     string `json:"title"`
		UserCover string `json:"user_cover"`
		//Keyframe         string   `json:"keyframe"`
		//IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime string `json:"live_time"`
		//Tags             string   `json:"tags"`
		//IsAnchor         int      `json:"is_anchor"`
		//RoomSilentType   string   `json:"room_silent_type"`
		//RoomSilentLevel  int      `json:"room_silent_level"`
		//RoomSilentSecond int      `json:"room_silent_second"`
		AreaName       string `json:"area_name"`
		ParentAreaName string `json:"parent_area_name"`
		//Pendants         string   `json:"pendants"`
		//AreaPendants     string   `json:"area_pendants"`
		//HotWords         []string `json:"hot_words"`
		//HotWordsStatus   int      `json:"hot_words_status"`
		//Verify           string   `json:"verify"`
		//NewPendants      struct {
		//	Frame struct {
		//		Name       string `json:"name"`
		//		Value      string `json:"value"`
		//		Position   int    `json:"position"`
		//		Desc       string `json:"desc"`
		//		Area       int    `json:"area"`
		//		AreaOld    int    `json:"area_old"`
		//		BgColor    string `json:"bg_color"`
		//		BgPic      string `json:"bg_pic"`
		//		UseOldArea bool   `json:"use_old_area"`
		//	} `json:"frame"`
		//	Badge struct {
		//		Name     string `json:"name"`
		//		Position int    `json:"position"`
		//		Value    string `json:"value"`
		//		Desc     string `json:"desc"`
		//	} `json:"badge"`
		//	MobileFrame struct {
		//		Name       string `json:"name"`
		//		Value      string `json:"value"`
		//		Position   int    `json:"position"`
		//		Desc       string `json:"desc"`
		//		Area       int    `json:"area"`
		//		AreaOld    int    `json:"area_old"`
		//		BgColor    string `json:"bg_color"`
		//		BgPic      string `json:"bg_pic"`
		//		UseOldArea bool   `json:"use_old_area"`
		//	} `json:"mobile_frame"`
		//	MobileBadge interface{} `json:"mobile_badge"`
		//} `json:"new_pendants"`
		//UpSession            string `json:"up_session"`
		//PkStatus             int    `json:"pk_status"`
		//PkID                 int    `json:"pk_id"`
		//BattleID             int    `json:"battle_id"`
		//AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		//AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		//StudioInfo           struct {
		//	Status     int           `json:"status"`
		//	MasterList []interface{} `json:"master_list"`
		//} `json:"studio_info"`
	}

	// DanmuInfo /xlive/web-room/v1/index/getDanmuInfo?RoomID=%d&type=0
	// where RoomID=RoomID
	DanmuInfo struct {
		//Group            string  `json:"group"`
		//BusinessID       int     `json:"business_id"`
		//RefreshRowFactor float64 `json:"refresh_row_factor"`
		//RefreshRate      int     `json:"refresh_rate"`
		//MaxDelay         int     `json:"max_delay"`
		Token    json.RawMessage `json:"token"`
		HostList []HostList      `json:"host_list"`
	}

	HostList struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		WssPort int    `json:"wss_port"`
		WsPort  int    `json:"ws_port"`
	}

	// RoomPlayInfo /xlive/web-room/v2/index/getRoomPlayInfo?room_id=%d&protocol=0,1&format=0,1,2&codec=0,1&qn=10000&platform=web&ptype=8
	// TODO 新协议 移动端 (br压缩级别)
	RoomPlayInfo struct {
		//RoomID          int           `json:"room_id"`
		//ShortID         int           `json:"short_id"`
		//UID             int           `json:"uid"`
		//IsHidden        bool          `json:"is_hidden"`
		//IsLocked        bool          `json:"is_locked"`
		//IsPortrait      bool          `json:"is_portrait"`
		LiveStatus int `json:"live_status"`
		//HiddenTill      int           `json:"hidden_till"`
		//LockTill        int           `json:"lock_till"`
		//Encrypted       bool          `json:"encrypted"`
		//PwdVerified     bool          `json:"pwd_verified"`
		LiveTime int `json:"live_time"`
		//RoomShield      int           `json:"room_shield"`
		//AllSpecialTypes []interface{} `json:"all_special_types"`
		PlayurlInfo struct {
			//ConfJSON string `json:"conf_json"`
			Playurl struct {
				//Cid     int `json:"cid"`
				GQnDesc []struct {
					Qn   int    `json:"qn"`
					Desc string `json:"desc"`
				} `json:"g_qn_desc"`
				Stream []struct {
					ProtocolName string `json:"protocol_name"`
					Format       []struct {
						FormatName string `json:"format_name"`
						Codec      []struct {
							CodecName string `json:"codec_name"`
							CurrentQn int    `json:"current_qn"`
							//AcceptQn  []int  `json:"accept_qn"`
							BaseURL string `json:"base_url"`
							URLInfo []struct {
								Host      string `json:"host"`
								Extra     string `json:"extra"`
								StreamTTL int    `json:"stream_ttl"`
							} `json:"url_info"`
						} `json:"codec"`
					} `json:"format"`
				} `json:"stream"`
				//P2PData struct {
				//	P2P      bool        `json:"p2p"`
				//	P2PType  int         `json:"p2p_type"`
				//	MP2P     bool        `json:"m_p2p"`
				//	MServers interface{} `json:"m_servers"`
				//} `json:"p2p_data"`
				//DolbyQn interface{} `json:"dolby_qn"`
			} `json:"playurl"`
		} `json:"playurl_info"`
	}
	Stream struct {
	}
)
