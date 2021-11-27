package main

import (
	"encoding/json"
)

type Config struct {
	App     ConfigApp                  `json:"App"`
	Default RoomConfig                 `json:"Default"`
	Rooms   map[string]json.RawMessage `json:"Rooms"`
}

type ConfigApp struct {
	Listen   string `json:"Listen"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RoomConfig struct {
	Enabled           bool `json:"Enabled"`
	FreshInterval     int  `json:"FreshInterval"`
	FreshLiveInterval int  `json:"FreshLiveInterval"`
}

var globalConfig Config
var globalRooms map[string]*Room

func (c *Config) Reset() {
	c.App.Listen = ":8080"
	c.App.Username = "admin"
	c.App.Password = "admin"
	c.Default.Reset()
}

func (c *Config) Load(data []byte) error {
	c.Rooms = make(map[string]json.RawMessage)
	return json.Unmarshal(data, c)
}

func (c *Config) Save() []byte {
	b, _ := json.MarshalIndent(c, "", "\t")
	return b
}

func (c *RoomConfig) Reset() {
	c.Enabled = true
	c.FreshInterval = 60 * 60
	c.FreshLiveInterval = 10 * 60
}

func (c *RoomConfig) Load(data json.RawMessage) error {
	var conf map[string]interface{}
	if err := json.Unmarshal(data, &conf); err != nil {
		return err
	}
	c.Reset()
	if v, ok := conf["Enabled"]; ok {
		c.Enabled = v.(bool)
	}
	if v, ok := conf["FreshInterval"]; ok {
		c.FreshInterval = int(v.(float64))
	}
	if v, ok := conf["FreshLiveInterval"]; ok {
		c.FreshLiveInterval = int(v.(float64))
	}
	return nil
}
