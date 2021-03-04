package genshin_public_cdkey

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type NGA struct {
	lastVisit   *http.Cookie
	lastPath    *http.Cookie
	guestJS     *http.Cookie
	passportUid *http.Cookie
	check       bool
	random      string
	pageId      string
	resp        []byte
}

var callback interface{}

func NewNGA(pageId string) NGA {

	return NGA{
		random:    strconv.Itoa(RandInt(200, 600)),
		pageId:    pageId,
		lastVisit: &http.Cookie{Name: "lastVisit", Value: "0"},
	}
}

func (ngaContent *NGA) Parser() (*Message, bool) {
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(ngaContent.resp))
	if err != nil {
		log.Fatalln(err)
	}

	var msg *Message
	var ret = false
	dom.Find("#postcontent0").Each(func(i int, selection *goquery.Selection) {
		str := selection.Text()
		str = ConvertToString(str, "gbk", "utf-8")
		msg, err = sort(str)
		if err == nil {
			ret = true
		}
	})
	return msg, ret
}

type Message struct {
	Title   string
	Content string
	MD5     string
}

type Push struct {
	FT       string
	FTSwitch bool
}

func sort(str string) (*Message, error) {
	k, err := FindInString(str, "[quote]", "[/quote]")
	str1 := strings.ToUpper(string(k))
	if err == nil {
		s := strings.Index(str1, " ")
		e := strings.Index(str1[s+1:], " ")
		return &Message{
			Title:   str1[s+1 : s+e+1],
			Content: str1,
			MD5:     MD5(str1[s+1 : s+e+1]),
		}, nil

	}
	return nil, errors.New("not found")
}

func (ngaContent *NGA) VisitPage() error {
	res, err := GET("https://ngabbs.com/read.php?tid="+ngaContent.pageId+"&rand="+ngaContent.random, func(reqPoint *http.Request) {
		ngaContent.baseHeader(reqPoint)
		if ngaContent.check {
			if ngaContent.lastVisit.Value == "0" {
				reqPoint.AddCookie(ngaContent.guestJS)
			}
			reqPoint.AddCookie(ngaContent.guestJS)
			reqPoint.AddCookie(ngaContent.passportUid)
		}
	})

	if err != nil {
		return errors.New("connected failed")
	} else if res == nil || res.StatusCode == http.StatusForbidden {
		result, _ := ioutil.ReadAll(res.Body)
		ind := strings.Index(string(result), "document.cookie = '")
		js := string(result)[ind+19 : ind+37]
		ngaContent.guestJS = &http.Cookie{Name: "guestJs", Value: js[8:]}
		println("[INFO] guestjs " + js[8:])
		ngaContent.check = true

		for _, v := range res.Cookies() {
			if v.Name == "lastVisit" {
				ngaContent.lastVisit = v
			} else if v.Name == "lastpath" {
				ngaContent.lastPath = v
			} else if v.Name == "ngaPassportUid" {
				ngaContent.passportUid = v
			}
		}
		return errors.New("check/recheck open")
	}

	ngaContent.resp, err = ioutil.ReadAll(res.Body)
	return err
}

func (ngaContent *NGA) baseHeader(reqPoint *http.Request) {
	reqPoint.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	reqPoint.Header.Add("Accept-Encoding", "")
	reqPoint.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	reqPoint.Header.Add("Connection", "keep-alive")
	reqPoint.Header.Add("Host", "ngabbs.com")
	reqPoint.Header.Add("Referer", "https://ngabbs.com/read.php?tid="+ngaContent.pageId+"&rand="+ngaContent.random)
	reqPoint.Header.Add("Sec-Fetch-Dest", "document")
	reqPoint.Header.Add("Sec-Fetch-Mode", "navigate")
	reqPoint.Header.Add("Sec-Fetch-Site", "same-origin")
	reqPoint.Header.Add("Upgrade-Insecure-Requests", "1")
}
