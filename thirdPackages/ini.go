package main
// iniはconfigファイルを読み込むもの

import (
  "fmt"

  "github.com/go-ini/ini"
)

type ConfigList struct {
  Port      int
  DbName    string
  DbDriver  string
}


var Config ConfigList

func init() {
  cfg, _ := ini.Load("config.ini")
  Config = ConfigList{
    Port: cfg.Section("web").Key("port").MustInt(),
    DbName: cfg.Section("db").Key("name").MustString("example.sql"),
    DbDriver: cfg.Section("db").Key("driver").String(),
  }
}

func main() {
  fmt.Printf("%T %v\n", Config.Port, Config.Port)
  fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
  fmt.Printf("%T %v\n", Config.DbDriver, Config.DbDriver)
}
