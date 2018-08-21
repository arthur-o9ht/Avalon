package main

import (
	"github.com/go-martini/martini"
	"os"
	"fmt"
)

var port string

func init(){
	name,_ := os.Hostname()
	if name == "ArthurdeMacBook-Pro.local" || name == "bogon"{
		port = ":8080"
	}else{
		port = ":80"
	}
	fmt.Println(name)
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("hugoSite/public"))
	m.RunOnAddr(port)
}