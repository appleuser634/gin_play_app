package main

import (
	"local.package/controller"
	"local.package/model"
	"local.package/router"
)

func main() {
	m := model.NewModel()
	c := controller.NewController(m)
	r := router.NewRouter(c)

	s := r.SetupRouter()
	s.Run(":3000")
}
