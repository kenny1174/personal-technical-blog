package utils

import (
	"fmt"

	ini "gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone       int
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8090")
	JwtKey = file.Section("server").Key("JwtKey").MustString("mykey")
}
func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("homeblog")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("homeblog")
}
func LoadQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}