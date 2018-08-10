package main

import (
	"github.com/go-martini/martini"
	)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("hugoSite/public"))
	m.RunOnAddr(":80")
	//m.RunOnAddr(":8080")
}