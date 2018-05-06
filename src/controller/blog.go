package controller

import (
	"html/template"
	"lib"
	"net/http"
)

func HandlerList(w http.ResponseWriter, r *http.Request) {
	posts := lib.GetList()
	t := template.New("list.html")
	t, _ = t.ParseFiles("source/tpl/list.html")
	t.Execute(w, posts)
}

func HandlerDoc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	posts := lib.GetDoc(name)
	t := template.New("doc.html")
	t, _ = t.ParseFiles("source/tpl/doc.html")
	t.Execute(w, posts)
}
