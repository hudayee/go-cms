package app

import (
	"cms/controller/role"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var App *iris.Application

func init() {
	App = iris.New()
	//App.Logger().SetLevel("warn")
	App.Use(logger.New())
	App.Use(recover.New())
	//App.OnAnyErrorCode(internalServerError)
	//role

	roleParty := App.Party("role")
	roleParty.Get("/", role.GetAll)
	roleParty.Post("/", role.Create)
	roleParty.Get("/{id:int}", role.GetById)
	roleParty.Put("/{id:int}", role.Update)

	//user

	//permissions

	//
}

func Start(addr string) {
	App.Run(iris.Addr(addr))
}
