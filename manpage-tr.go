package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

const (
	ConfigFile    = "config.ini"
	ConfigSection = "ManTrConfig"
)

type ConfigSetting struct {
	// 说明、注释
	Comment   string
	TrComment string
	// 日期
	TrData string
	// 头部说明
	Title   string
	TrTitle string

	// 说明标题
	SubTitle   string
	TrSubTitle string

	// 说明项
	SubSection   string
	TrSubSection string

	//说明内容
	SubContent   string
	TrSubContent string

	//内容部分: 起始-结束
	ContentBegin   string
	ContentEnd     string
	TrContentBegin string
	TrContentEnd   string
}

var config = &ConfigSetting{}
var cfg *ini.File

func init() {
	cfg, _ := ini.Load(ConfigFile)
	cfgSection := cfg.Section(ConfigSection)
	cfgSection.MapTo(config)

	// fmt.Println(cfg.Section(ConfigSection).KeyStrings())

	// fmt.Println("config-Comment:", config.Comment)
	// fmt.Println("config-Title:", config.Title, "---->", config.TrTitle)
	// fmt.Println("config-SubTitle:", config.SubTitle, "---->", config.TrSubTitle)
	// fmt.Println("config-SubSection:", config.SubSection, "---->", config.TrSubSection)
	// fmt.Println("config-SubSectionContent:", config.SubContent, "---->", config.TrSubContent)
	// fmt.Println("config-SubSectionContent:", config.ContentBegin, "---->", config.TrContentBegin)
	// fmt.Println("config-SubSectionContent:", config.ContentEnd, "---->", config.TrContentEnd)
}

func format(source string, old string, now string) string {
	now = strings.Replace(now, "\\n", "\n", 1)
	// now = strings.Replace(now, "\\t", "\t", 1)
	return strings.ReplaceAll(source, old, now)
}

func formatAll(source string) string {
	temp := source
	now := time.Now()                  //获取当前时间
	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	temps := strings.Split(temp, "\n")
	temp = ""
	for _, v := range temps {
		temp = temp + strings.TrimLeft(v, " ") + "\n"
	}
	temp = format(temp, config.TrData, timeObj.String())
	temp = format(temp, config.TrComment, config.Comment)
	temp = format(temp, config.TrTitle, config.Title)
	temp = format(temp, config.TrSubTitle, config.SubTitle)
	temp = format(temp, config.TrSubSection, config.SubSection)
	temp = format(temp, config.TrSubContent, config.SubContent)
	temp = format(temp, config.TrContentBegin, config.ContentBegin)
	temp = format(temp, config.TrContentEnd, config.ContentEnd)

	fmt.Println("生成时间:", timeObj)
	return temp
}

func checkFileExist() bool {
	_, err := os.Stat(os.Args[1])
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func usage() {
	if len(os.Args) == 1 {
		// Usage:  docker [OPTIONS] COMMAND
		fmt.Println(fmt.Sprintf("-- Usage:  %s <FILE>.tr ", path.Base(os.Args[0])))
		os.Exit(0)
	} else {
		if !strings.HasSuffix(os.Args[1], "tr") {
			fmt.Println(fmt.Sprintf("%s: %s 不是有效的文件扩展", path.Base(os.Args[0]), os.Args[1]))
			fmt.Println(fmt.Sprintf("Usage:  %s <FILE>.tr ", path.Base(os.Args[0])))
			os.Exit(0)
		}

		if !checkFileExist() {
			fmt.Println(fmt.Sprintf("%s: %s 不是有效的文件", path.Base(os.Args[0]), os.Args[1]))
			fmt.Println(fmt.Sprintf("Usage:  %s <FILE>.tr ", path.Base(os.Args[0])))
			os.Exit(0)
		}
	}
}

func main() {
	usage()
	fmt.Println("加载配置: config.ini")

	sf := os.Args[1]
	tf := strings.Replace(path.Base(os.Args[1]), ".tr", "", 1) + ".1"
	fmt.Println("读取文件:", sf)
	fmt.Println("生成文件:", tf)
	data, _ := ioutil.ReadFile(sf)
	out := formatAll(string(data))
	ioutil.WriteFile(tf, []byte(out), 0666)
}
