package genshin_public_cdkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"time"
)

/**
 * local config struct
 */
type LocalConfig struct {
	FT        string    `json:"ft"`
	CacheTime time.Time `json:"-"`
}

var config = new(LocalConfig)
var configPath = "config.json"

func GetConfig() *LocalConfig {
	if config.CacheTime.Before(time.Now()) {
		if err := LoadConfig(configPath, config); err != nil {
			println("loading file failed")
			return nil
		}
		config.CacheTime = time.Now().Add(time.Second * 60)
	}
	return config
}

/**
 * save cnf/config.json
 */
func (lc *LocalConfig) SetConfig() error {
	fp, err := os.Create(configPath)
	if err != nil {
		println("loading file failed")
	}
	defer fp.Close()
	data, err := json.Marshal(lc)
	if err != nil {
		println("marshal file failed")
	}
	n, err := fp.Write(data)
	if err != nil {
		println("write file failed", n)
	}
	println("already update config file")
	return nil
}

const configFileSizeLimit = 10 << 20

/**
 * Load File
 * @param path 文件路径
 * @param dist 存放目标
 */
func LoadConfig(path string, dist interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		println("Failed to open config file.")
		return err
	}

	fi, _ := configFile.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		println("Config file size exceeds reasonable limited")
		return errors.New("limited")
	}

	if fi.Size() == 0 {
		println("Config file is empty, skipping")
		return errors.New("empty")
	}

	buffer := make([]byte, fi.Size())
	_, err = configFile.Read(buffer)
	buffer, err = StripComments(buffer)
	if err != nil {
		println("Failed to strip comments from json")
		return err
	}

	buffer = []byte(os.ExpandEnv(string(buffer)))

	err = json.Unmarshal(buffer, &dist)
	if err != nil {
		println("Failed unmarshalling json")
		return err
	}
	return nil
}

/**
 * 注释清除
 */
func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0)
	lines := bytes.Split(data, []byte("\n"))
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}
	return bytes.Join(filtered, []byte("\n")), nil
}
