package otp

import (
	"errors"
	"fmt"
)

type DefaultImplement struct {
}

func (d DefaultImplement) genRandomOTP() string {
	randomOTP := "1234"
	fmt.Printf("生成随机opt: %s\n", randomOTP)
	return randomOTP
}

func (d DefaultImplement) saveOTPCache(otp string) {
	fmt.Printf("保存opt: %s 到缓存中\n", otp)
}

func (d DefaultImplement) getMessage(otp string) string {
	return "你的 OTP 是 " + otp
}

func (d DefaultImplement) sendNotification(s string) error {
	return errors.New("未实现具体的发送信息接口")
}
