package controller

import (
	"net/http"
	"lib"
	"html/template"
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
