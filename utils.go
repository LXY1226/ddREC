package main

import "log"

const DEBUG = true

func DebugPrint(v ...interface{}) {
	if DEBUG {
		log.Println(v...)
	}
}
