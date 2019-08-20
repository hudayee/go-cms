package config

import (
	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
)

var Config *goconfig.ConfigFile

func init() {
	var c = flag.String("c", "default", "config文件名称，默认default")
	flag.Parse()
	filePath := fmt.Sprintf("./config/%s.ini", *c)
	var err error
	if Config, err = goconfig.LoadConfigFile(filePath); err != nil {
		panic(err.Error())
	}
}
