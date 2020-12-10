package until

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbUser     string
	DbPort     string
	DbName     string
	DbPassWord string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("%s", err)
	}
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	fmt.Println(AppMode)
	fmt.Println(HttpPort)
}

func LoadDB(file *ini.File) {

}
