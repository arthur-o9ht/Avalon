package main

import (
	"controller"
	"github.com/go-martini/martini"
)

func serverInit() {
	m := martini.Classic()
	m.Group("/blog/", func(r martini.Router) {
		r.Get("list", controller.HandlerList)
		r.Get("doc", controller.HandlerDoc)
	})
	m.RunOnAddr(":8000")
}

func main() {
	serverInit()
}
