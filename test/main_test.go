package test

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {

	str := "和崩坏3一样，本帖会将看到的兑换码都收集起来，方便大家查找使用。同样有一个兑换码通知QQ群，永久禁言，只有发兑换码的时候我会艾特全体成员。有需要的可以加群，QQ群号628059351。原神这边兑换码的发布周期和模式不是特别清楚，如果有遗漏或者更新缓慢请见谅。[size=150%][color=red][b][align=center]↓最新兑换码看这里↓[/align][/b][/color][/size][quote]2021/02/04 b站拜年纪兑换码HTNA2GXGMMFN 23333摩拉 2个树脂[/quote][size=150%][color=red][b][align=center]↑最新兑换码看这里↑[/align][/b][/color][/size][b]往期礼包[/b][collapse=2021/02/04 b站拜年纪兑换码]HTNA2GXGMMFN 23333摩拉 2个树脂[/collapse][collapse=2021/01/22 1.3版本直播兑换码]HT7A24APV9VS 100原石 10个紫矿8SNB34SNVQDW 100原石 5本紫书 NSPT3MBNV9C2 100原石 50000摩拉[/collapse][collapse=2020/12/11 1.2版本直播兑换码]O6V2UN25 100原石 10个紫矿B6F4SS33G 100原石 5本书 6GS45K3KS 100原石 50000摩拉[/collapse][collapse=2020/10/30 1.1版本直播兑换码]TF42JK44K 100原石TJ23D45HG 150原石B2G257MFA 50000摩拉 5本紫书[/collapse][collapse=2020/09/19]FS6AW6J8B4C6 30原石KT6AX6JQS5UW 50原石DTNSF7K8A5D2 50原石[/collapse]\n"
	k, e := findInString(str, "[quote]", "[/quote]")
	println(e)
	println(string(bytes.ToUpper(k)))
}

func TestFT(t *testing.T) {
	//p := Push{
	//	FT:       "SCU95761Te034c249fe732b1b5cb523f8effaf9d05ea6ece34e3d9",
	//	FTSwitch: true,
	//}
	//msg.title = "213"
	//msg.content = "234"
	//p.SendMessage2WeChat(msg)
}

func findInString(str, start, end string) ([]byte, error) {
	var match []byte
	index := strings.Index(str, start)

	if index == -1 {
		return match, errors.New("not found")
	}

	index += len(start)

	for {
		char := str[index]

		if strings.HasPrefix(str[index:index+len(match)], end) {
			break
		}

		match = append(match, char)
		index++
	}

	return match, nil
}