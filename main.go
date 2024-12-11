package main

import (
	"log"

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

	// SendXMlTest("测试", wcf)

	// GetChatMsgForDB(wcf)

	// GetContactInfo(wcf)
	GetAllContacts(wcf)
}

func SendXMlTest(msg string, wcf *client.ClientWCF) {
	xml := `<emoji fromusername="wxid_413p6ngt3cq612" tousername="48988100861@chatroom" type="2" androidmd5="5fcecbdfccc5b0ad99b2f6ca8f78f5dd" androidlen="22475" aeskey="82cdf8fc5ead45aaaca79aa0fb1dd024" encrypturl="http://vweixinf.tc.qq.com/110/20402/stodownload?m=213be8b8c6eecddd639272e5e525f1b9&amp;filekey=30440201010430302e02016e0402534804203231336265386238633665656364646436333932373265356535323566316239020300a8e0040d00000004627466730000000132&amp;hy=SH&amp;storeid=2641b1ab90008d62b83ef304b0000006e02004fb253481806db40b643c55b3&amp;ef=2&amp;bizid=1022" externurl="http://vweixinf.tc.qq.com/110/20403/stodownload?m=32f634e7ae4c067067ca4e063e28ae47&amp;filekey=3043020101042f302d02016e0402534804203332663633346537616534633036373036376361346530363365323861653437020257d0040d00000004627466730000000132&amp;hy=SH&amp;storeid=2641b1ab90009588183ef304b0000006e03004fb353481806db40b643c55c0&amp;ef=3&amp;bizid=1022" externmd5="2505370c2c52e01cef30861b669175ad"/></msg>
2024/12/11 19:06:09 发送时间: 2024-12-11 18:40:54, 发送人: 48988100861@chatroom, 消息内容: <msg><emoji fromusername = "wxid_lyp7rlu2ind121" tousername = "48988100861@chatroom" type="2" idbuffer="media:0_0" md5="8d8c05846e9ad2c9c480c052b0ddf9c3" len = "39962" productid="" androidmd5="8d8c05846e9ad2c9c480c052b0ddf9c3" androidlen="39962" s60v3md5 = "8d8c05846e9ad2c9c480c052b0ddf9c3" s60v3len="39962" s60v5md5 = "8d8c05846e9ad2c9c480c052b0ddf9c3" s60v5len="39962" cdnurl = "http://wxapp.tc.qq.com/262/20304/stodownload?m=8d8c05846e9ad2c9c480c052b0ddf9c3&amp;filekey=30350201010421301f020201060402534804108d8c05846e9ad2c9c480c052b0ddf9c30203009c1a040d00000004627466730000000132&amp;hy=SH&amp;storeid=263178c22000dd5ae000000000000010600004f50534818e67b40b6483bfbe&amp;bizid=1023" designerid = "" thumburl = "" encrypturl = "http://wxapp.tc.qq.com/262/20304/stodownload?m=b62780be5e8d2c14cf2f8c7885801276&amp;filekey=30350201010421301f02020106040253480410b62780be5e8d2c14cf2f8c78858012760203009c20040d00000004627466730000000132&amp;hy=SH&amp;storeid=263178c2300014673000000000000010600004f5053481d167b40b64833766&amp;bizid=1023" aeskey= "2e678a0307166b8165c3cd71882a684e" externurl = "http://wxapp.tc.qq.com/262/20304/stodownload?m=1a51c0cb470cb61ce921d8e86540edb0&amp;filekey=30340201010420301e020201060402535a04101a51c0cb470cb61ce921d8e86540edb002020fe0040d00000004627466730000000132&amp;hy=SZ&amp;storeid=263178c2300045618000000000000010600004f50535a00e278809635ca3e2&amp;bizid=1023" externmd5 = "8a8650c3f25408826e0620c62e1e9dc0" width= "427" height= "172" tpurl= "" tpauthkey= "" attachedtext= "" attachedtextcolor= "" lensid= "" emojiattr= "Cijmk43kvaDlpojnmoTlgrvpgLzkuJzopb8g5aSn5a625Yir6aqC5Lq6" linkid= "" desc= "" ></emoji> <gameext type="0" content="0" ></gameext>`
	// wxid_lyp7rlu2ind121 48988100861@chatroom
	if err := wcf.SendXml("48988100861@chatroom", xml, "", 49); err != nil {
		log.Fatalf("Failed to send xml: %v", err)
	}
}

// 获取聊天记录
func GetChatMsgForDB(wcf *client.ClientWCF) {
	// 执行SQL查询
	// rows, err := wcf.ExecDbQuery("MSG0.db", "SELECT * FROM MSG")
	// if err != nil {
	// 	log.Fatalf("Failed to execute query: %v", err)
	// }
	// for _, row := range rows {
	// 	// IsSender 是 int64 类型，需要进行类型断言
	// 	if isSender, ok := row["IsSender"].(int64); ok && isSender == 1 {
	// 		log.Printf("发送时间: %s, 发送人: %s, 消息内容: %s", time.Unix(row["CreateTime"].(int64), 0).Format("2006-01-02 15:04:05"), "本人", row["StrContent"])
	// 	} else {
	// 		log.Printf("发送时间: %s, 发送人: %s, 消息内容: %s", time.Unix(row["CreateTime"].(int64), 0).Format("2006-01-02 15:04:05"), row["StrTalker"], row["StrContent"])
	// 	}
	// 	// log.Printf("Row: %v", row)
	// }

	rows, err := wcf.ExecDbQuery("MSG0.db", "SELECT * FROM MSG WHERE StrTalker = 'wxid_lyp7rlu2ind121'")
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	for _, row := range rows {
		log.Printf("Row: %v", row)
	}
}

func GetContacts(wcf *client.ClientWCF) {
	// 获取联系人信息
	contacts, err := wcf.GetContacts()
	if err != nil {
		log.Fatalf("Failed to get contact info: %v", err)
	}
	for _, contact := range contacts {
		log.Printf("Contact: %v", contact)
	}
}

func GetContactInfo(wcf *client.ClientWCF) {
	// 获取联系人信息
	contact, err := wcf.GetContactInfo("wxid_lyp7rlu2ind121")
	if err != nil {
		log.Fatalf("Failed to get contact info: %v", err)
	}
	log.Printf("Contact info: %v", contact)
}

func GetAllContacts(wcf *client.ClientWCF) {
	// 获取所有联系人信息
	contacts, err := wcf.GetContacts()
	if err != nil {
		log.Fatalf("Failed to get all contact info: %v", err)
	}

	for _, contact := range contacts {
		log.Printf("wxid: %s, 名称: %s, 备注: %s", contact.Wxid, contact.GetName(), contact.GetRemark())
	}
}
