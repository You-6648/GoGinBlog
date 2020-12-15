package until

import (
	"GoGinBlog/until/log"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DbData     string
	DbHost     string
	DbUser     string
	DbPort     string
	DbName     string
	DbPassWord string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		//fmt.Printf("%s", err)
		log.Info("读取配置信息出错:%v",err)
	}
	LoadServer(file)
	LoadDB(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadDB(file *ini.File) {
	DbName = file.Section("database").Key("DbName").String()
	DbData = file.Section("database").Key("DbData").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
}
