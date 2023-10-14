package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HellTom()
	expect := "Tom"
	// if output != expect {
	// 	t.Errorf("not match")
	// }
	assert.Equal(t, expect, output)
}

// go test P_test.go Print.go --cover
// 可以计算覆盖率 以代码行数计算 
// 覆盖率100% 说明原函数所有行都用上了