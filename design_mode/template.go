package main

import (
	"design_mode/otpassword_template"
	"fmt"
)

// 模板方法设计模式

func template_test() {
	fmt.Println()
	fmt.Println("sms 发送 opt 的步骤:")
	smsImplement := otp.SmsImplement{}
	smsExecutor := otp.TemplateExecutor{
		IOtp: smsImplement,
	}
	_ = smsExecutor.Execute()

	fmt.Println()
	fmt.Println("email 发送 otp 的步骤:")
	emailImplement := otp.EmailImplement{}
	emailExcutor := otp.TemplateExecutor{
		IOtp: emailImplement,
	}
	_ = emailExcutor.Execute()
}

/*
sms 发送 opt 的步骤:
生成随机opt: 1234
保存opt: 1234 到缓存中
调用【短信】接口发送信息: 你的 OTP 是 1234

email 发送 otp 的步骤:
生成随机opt: 1234
保存opt: 1234 到缓存中
调用【邮件】接口发送信息: 你的 OTP 是 1234
*/
