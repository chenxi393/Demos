package main

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestProcessFirstLine(t *testing.T) {

	line := ProcessFirstLine()
	assert.Equal(t, "line00", line)
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	//注意这里patch unpatch 是函数名 不然会出错
	//运行时替换地址 为一个函数打桩
	//也就是替换下面函数的依赖函数

	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line0000", line)
}
