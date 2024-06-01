package main

import (
	"fmt"
	"runtime"
)

func getMessageType(msg []byte) []byte {
	return msg[:5]
 }
 
 func getMessageTypeWithCopy(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
 }
 
 func receiveMessage() []byte {
	return make([]byte, 1024*1024*1024)
 }
 
 func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
 }
 
 func SliceGC() {
	printAlloc()
	msg := receiveMessage()
	five := getMessageType(msg)
	printAlloc()
	runtime.GC()
	runtime.KeepAlive(five) // 保持对five的引用
	printAlloc()
 }