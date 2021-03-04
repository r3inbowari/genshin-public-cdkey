package genshin_public_cdkey

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/axgle/mahonia"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func SendMessage2WeChat(title string, content ...string) {
	//ft := biu.FT
	//if biu.FTSwitch && ft != "" {
	//	if len(content) > 0 {
	//		GetRequest("https://sc.ftqq.com/" + ft + ".send?desp=" + content[0] + "&text=" + title)
	//	} else {
	//		GetRequest("https://sc.ftqq.com/" + ft + ".send?text=" + title)
	//	}
	//}
}

func GET(url string, interceptor func(reqPoint *http.Request)) (*http.Response, error) {
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if interceptor != nil {
		interceptor(req)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	res, err := client.Do(req)
	return res, nil
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func FindInString(str, start, end string) ([]byte, error) {
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

func GetRequest(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Println("[FAIL] GET错误")
	}
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
