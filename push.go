package genshin_public_cdkey

import "strings"

// 方糖
func (p *Push) SendMessage2WeChat(msg *Message) {
	if OK(p.FT, msg.MD5) {
		return
	}
	ft := p.FT
	msg.Content = strings.Replace(msg.Content, " ", "%20", -1)
	if p.FTSwitch && ft != "" {
		if len(msg.Content) > 0 {
			GetRequest("https://sc.ftqq.com/" + ft + ".send?desp=" + msg.Content + "&text=" + msg.Title)
		} else {
			GetRequest("https://sc.ftqq.com/" + ft + ".send?text=" + msg.Title)
		}
	}
	PushOK(p.FT, msg.MD5)
	println("[INFO] push ok -> " + p.FT)
}
