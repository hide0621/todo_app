package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

//フィールドがconfig.iniと対応する
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

//グローバルで宣言
var Config ConfigList

//main関数より前にLoadConfig関数を実行する
func init() {
	LoadConfig()
}

func LoadConfig() {
	//config.iniの読み込み
	cfg, err := ini.Load("config.ini")

	if err != nil {
		log.Fatalln(err)
	}

	//各種の設定を打ち込む
	Config = ConfigList{

		Port: cfg.Section("web").Key("port").MustString("8080"),

		SQLDriver: cfg.Section("db").Key("driver").String(),

		DbName: cfg.Section("db").Key("name").String(),

		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}
