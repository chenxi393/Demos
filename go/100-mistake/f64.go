package main

import (
	"encoding/json"
	"fmt"
)

func listing1() error {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
	   return err
	}
	fmt.Printf("%T\n", m["id"]) // float64
	if m["id"] == 32 {
	   fmt.Println("id is int && id == 32") // 不输出
	} else if m["id"] == float64(32) {
	   fmt.Println("id is float && id == float64(32)") // 输出此项  id is float && id == float64(32)
 
	}
	return nil
 }
 
 func getMessage() []byte {
	str := "{\n        \"id\": 32,\n        \"name\": \"foo\"\n}"
	return []byte(str)
 }
 
 func main() {
	listing1()
 }