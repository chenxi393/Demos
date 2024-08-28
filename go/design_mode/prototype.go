package main

import (
	"encoding/json"
	"fmt"
)

// 原型（Prototype）模式
type Member struct {
	StringBody string
}

type ProtoType struct {
	MemberStruct *Member
}

// 深拷贝类型
// 克隆方法，复制一份和之前完全一致的结构体（这里使用序列化，防止浅拷贝问题）
func (p *ProtoType) ProtoTypeClone() (error, *ProtoType) {
	marshal, err := json.Marshal(p)
	if err != nil {
		return err, nil
	}
	copyProtoType := &ProtoType{}
	err = json.Unmarshal(marshal, copyProtoType)
	if err != nil {
		return err, nil
	}
	return nil, copyProtoType
}

func (p *ProtoType) Show() {
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))
}
