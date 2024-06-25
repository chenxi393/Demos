package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func Test_Time() {
	type Event1 struct {
		ID   int
		Time time.Time
	}

	type Event2 struct {
		ID        int
		time.Time // time 结构体实现了  MarshalJSON 所以会只marshal time
	}

	event := Event1{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return
	}

	fmt.Println(string(b))
}

// TODO  这个好像不是很常见
// wall和ext共同记录了时间，但是分为两种情况，一种是没有记录单调时钟（比如是通过字符串解析得到的时间）
// 另一种是记录了单调时钟（比如通过Now）。
func Time_Unmarshal() {
	type Event struct {
		Time time.Time
	}
	// 通过time.Now()
	t := time.Now()
	event1 := Event{
		Time: t,
	}

	b, err := json.Marshal(event1)
	if err != nil {
		return
	}

	var event2 Event
	err = json.Unmarshal(b, &event2) // 通过序列号生成
	if err != nil {
		return
	}

	fmt.Println(event1 == event2)

	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
}
