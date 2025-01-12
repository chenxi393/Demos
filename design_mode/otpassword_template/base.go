package otp

type Template interface {
	genRandomOTP() string          // 生成随机密码
	saveOTPCache(string)           // 保存生成的密码
	getMessage(string) string      // 准备发送内容
	sendNotification(string) error // 发送信息
}

type TemplateExecutor struct {
	IOtp Template
}

func (t TemplateExecutor) Execute() error {
	otp := t.IOtp.genRandomOTP()
	t.IOtp.saveOTPCache(otp)
	message := t.IOtp.getMessage(otp)
	err := t.IOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
