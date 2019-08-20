package models

import (
	"cms/config"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

var Engine *xorm.Engine

func init() {
	var err error
	var dataSource string
	if dataSource, err = config.Config.GetValue("DATABASE", "dataSource"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataSource)
	Engine, err = xorm.NewEngine("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	Engine.ShowSQL(true)
}
