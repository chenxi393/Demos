package main

// 77 json 任何数值，当将它通过JSON反序列化到一个map中时，无论数值是否包含小数，都将被转化为float64类型
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
	str := `{"id": 32,    "name": "foo"}`
	return []byte(str)
}
