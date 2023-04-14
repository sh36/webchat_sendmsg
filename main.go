package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	Run()
}

func Run() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	err := bot.Login()
	if err != nil {
		log.Printf("login error: %v \n", err)
		return
	}

	fmt.Println("hello")
	//异步执行一个定时发送消息的函数，周期为8小时
	go sendToPerson(bot, "金科小兴", "在吗")
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

// 每隔8小时向指定人发送一条消息
func sendToPerson(bot *openwechat.Bot, userID string, message string) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Println("Error getting current user:", err)
		return
	}

	friends, err := self.Friends()
	if err != nil {
		log.Println("Error getting friends list:", err)
		return
	}

	for {
		// 生成发送消息的随机时间点
		interval := time.Duration(rand.Intn(5)) * time.Minute
		duration := 8*time.Hour + interval

		// 等待随机时间点
		log.Printf("Waiting for %v before sending next message to %s\n", duration, userID)
		<-time.After(duration)

		result := friends.SearchByRemarkName(1, userID)

		// 在好友列表中查找指定用户
		if len(result) == 0 {
			log.Printf("Could not find user with remark name '%s'\n", userID)
			continue
		}

		// 向指定用户发送一条消息
		if err := result.SendText(message); err != nil {
			log.Printf("Error sending message to '%s': %v\n", userID, err)
		} else {
			log.Printf("Sent message '%s' to %s at %s\n", message, userID, time.Now().Format("2023-04-14 16:13:36"))
		}
	}
}

// 向指定人发送一条消息
func sendmessageToPerson(bot *openwechat.Bot, userID string, message string) {
	self, err := bot.GetCurrentUser()

	if err != nil {
		return
	}

	firends, err := self.Friends()

	sult := firends.SearchByRemarkName(1, userID)
	// 向指定人发送一条消息
	sult.SendText(message)

}
