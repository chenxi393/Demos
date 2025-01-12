package otp

import "fmt"

type SmsImplement struct {
	DefaultImplement
}

func (s SmsImplement) sendNotification(message string) error {
	fmt.Printf("调用【短信】接口发送信息: %s\n", message)
	return nil
}
