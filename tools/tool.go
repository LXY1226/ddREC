package tools

import "unsafe"

// BytesToU32 仅用于比较
func BytesToU32(b []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

// StringToU32 仅用于比较
func StringToU32(s string) uint32 {
	return BytesToU32([]byte(s))
}
