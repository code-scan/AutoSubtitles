package config

import (
	"gopkg.in/ini.v1"
)

var AccessKeyID = ""
var AccessKey = ""
var EndPoint = ""
var BucketName = ""
var Appkey = ""

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	AccessKeyID = cfg.Section("").Key("AccessKeyID").String()
	AccessKey = cfg.Section("").Key("AccessKey").String()
	EndPoint = cfg.Section("").Key("EndPoint").String()
	BucketName = cfg.Section("").Key("BucketName").String()
	Appkey = cfg.Section("").Key("Appkey").String()
}
