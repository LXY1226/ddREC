package main

import (
	"log"
	"os"
)

const configFile = "config.json"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf, err := os.ReadFile(configFile)
	if err != nil {
		os.WriteFile(configFile, globalConfig.Save(), 0644)
		panic(err)
	}
	err = globalConfig.Load(conf)
	if err != nil {
		os.WriteFile(configFile, globalConfig.Save(), 0644)
		panic(err)
	}
	for name, room := range globalConfig.Rooms {
		r := new(Room)
		r.Init(name, room)
		go r.Connect()
		globalConfig.rooms[name] = r
	}
	select {}
}
