package otp

import "fmt"

type EmailImplement struct {
	DefaultImplement
}

func (e EmailImplement) sendNotification(message string) error {
	fmt.Printf("调用【邮件】接口发送信息: %s\n", message)
	return nil
}
