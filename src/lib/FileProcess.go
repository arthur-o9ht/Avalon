package lib

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Doc struct {
	Title   string
	Date    string
	Summary string
	Body    template.HTML
	File    string
	Name    string
}

func GetDoc(name string) []Doc {
	a := []Doc{}
	a = Md2Html(name, a)
	return a
}

func GetList() []Doc {
	a := []Doc{}
	files, _ := filepath.Glob("source/mds/*")
	for _, f := range files {
		file := strings.Replace(f, "source/mds/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		a = Md2Html(file, a)
	}
	return a
}

func Md2Html(name string, docs []Doc) []Doc {
	f := "source/mds/" + name + ".md"
	fileRead, _ := ioutil.ReadFile(f)
	lines := strings.Split(string(fileRead), "\n")
	title := string(lines[0])
	date := string(lines[1])
	summary := string(lines[2])
	bodyTemp := strings.Join(lines[3:len(lines)], "\n")
	bodyTemp = string(blackfriday.MarkdownCommon([]byte(bodyTemp)))
	body := template.HTML(bodyTemp)
	docs = append(docs, Doc{title, date, summary, body, f, name})
	return docs
}
