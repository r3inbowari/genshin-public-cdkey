package main

import (
	"fmt"
	genshin_public_cdkey "genshin-public-cdkey"
	"github.com/jasonlvhit/gocron"
	"time"
)

var nga genshin_public_cdkey.NGA
var msg genshin_public_cdkey.Message

func taskReduce() {
	fmt.Println("[INFO] check expire")
	nga = genshin_public_cdkey.NewNGA("23406619")
	err := nga.VisitPage()
	if err != nil {
		fmt.Println("[INFO] " + err.Error())
	}
}

func taskGET() {
	err1 := nga.VisitPage()
	if err1 != nil {
		fmt.Println("[INFO] " + err1.Error())
	} else {
		msgRet, ok := nga.Parser()
		// println("diff")
		if ok {
			// if msg.Title != msgRet.Title {
			// msg = *msgRet
			// println(msg.Title)
			// println(msg.Content)
			// println(msg.MD5)

			p := genshin_public_cdkey.Push{
				FT:       genshin_public_cdkey.GetConfig().FT,
				FTSwitch: true,
			}
			p.SendMessage2WeChat(msgRet)

			// }
		}
	}
}

func main() {
	genshin_public_cdkey.InitRedis()
	taskReduce()
	time.Sleep(time.Second)
	s := gocron.NewScheduler()
	_ = s.Every(1).Hour().Do(taskReduce)
	_ = s.Every(5).Seconds().Do(taskGET)
	s.Start()

	select {}
}
