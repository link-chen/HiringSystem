package Utils

import (
	"HiringSystem/Config"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendEmail(Target string, Subject string, EmailMessage string) {
	em := email.NewEmail()
	em.From = Config.GetEmail() //发送者
	em.To = []string{Target}
	// 设置主题
	em.Subject = Subject
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(EmailMessage)
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", Config.GetEmail(), Config.GetEmailCode(), "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}

func SendSuccessEmail(Target string) {
	EmailMessage := "我们很高兴您选择加入我们的团队，热烈欢迎您的到来，请添加团队qq群1036421449"
	SendEmail(Target, "来自档案组Archivare的Offer", EmailMessage)
}

func SendFailedEmail(Target string) {
	SendEmail(Target, "来自档案组Archivare的感谢", "感谢您对于档案组的信任，经过谨慎的评估，我们认为您目前不适合对应的岗位，希望未来与您相会")
}
