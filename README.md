# WeChatFerry Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/yaoyaochil/wcf-client-go.svg)](https://pkg.go.dev/github.com/yaoyaochil/wcf-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/yaoyaochil/wcf-client-go)](https://goreportcard.com/report/github.com/yaoyaochil/wcf-client-go)
[![License](https://img.shields.io/github/license/yaoyaochil/wcf-client-go)](LICENSE)

基于 [WeChatFerry](https://github.com/lich0821/WeChatFerry) 开发的 Go 客户端 SDK，提供简单易用的 API 接口来操作微信客户端。

## 功能特性

- 获取登录二维码和登录状态
- 发送文本、图片、文件等多种类型消息
- 群组管理（添加/删除成员、邀请等）
- 接收并处理微信消息
- 获取联系人和群组信息
- 支持数据库查询操作
- 支持语音、图片等多媒体消息处理

## 安装 

```bash
go get github.com/yaoyaochil/wcf-client-go
```

## 使用

```go

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
	contacts, err := wcf.GetContacts()
	if err != nil {
		log.Fatalf("Failed to get contacts: %v", err)
	}
	for _, contact := range contacts {
		log.Printf("Contact: %v", contact)
	}

	// 获取个人信息
	userInfo, err := wcf.GetSelfContactInfo()
	if err != nil {
		log.Fatalf("Failed to get self contact info: %v", err)
	}
	log.Printf("Self contact info: %v", userInfo)
}
```

## 感谢

- [lich0821](https://github.com/lich0821/WeChatFerry) 提供的WeChatFerry项目