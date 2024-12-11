package main

import (
	"log"
	"time"

	"github.com/yaoyaochil/wcf-client-go/client"
)

func main() {
	// 创建客户端
	wcf, err := client.NewClient("tcp://192.168.2.111:10086")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer wcf.Close()

	// 检查登录状态
	isLogin, err := wcf.IsLogin()
	if err != nil {
		log.Fatalf("Failed to check login status: %v", err)
	}
	log.Printf("Login status: %v", isLogin)

	// 获取自己的微信ID
	wxid, err := wcf.GetSelfWxid()
	if err != nil {
		log.Fatalf("Failed to get self wxid: %v", err)
	}
	log.Printf("Self wxid: %s", wxid)

	// 获取通讯录
	// contacts, err := wcf.GetContacts()
	// if err != nil {
	// 	log.Fatalf("Failed to get contacts: %v", err)
	// }
	// log.Printf("Contacts: %v", contacts)

	// 获取个人信息
	userInfo, err := wcf.GetSelfContactInfo()
	if err != nil {
		log.Fatalf("Failed to get self contact info: %v", err)
	}
	log.Printf("Self contact info: %v", userInfo)

	// 获取联系人信息
	contact, err := wcf.GetContactInfo("wxid_lyp7rlu2ind121")
	if err != nil {
		log.Fatalf("Failed to get contact info: %v", err)
	}
	log.Printf("Contact info: %v", contact)

	// 执行SQL查询
	rows, err := wcf.ExecDbQuery("MSG0.db", "SELECT * FROM MSG LIMIT 20")
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	for _, row := range rows {
		// IsSender 是 int64 类型，需要进行类型断言
		if isSender, ok := row["IsSender"].(int64); ok && isSender == 1 {
			log.Printf("发送时间: %s, 发送人: %s, 消息内容: %s", time.Unix(row["CreateTime"].(int64), 0).Format("2006-01-02 15:04:05"), "本人", row["StrContent"])
		} else {
			log.Printf("发送时间: %s, 发送人: %s, 消息内容: %s", time.Unix(row["CreateTime"].(int64), 0).Format("2006-01-02 15:04:05"), row["StrTalker"], row["StrContent"])
		}
	}
}
