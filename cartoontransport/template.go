package cartoontransport

import (
	"github.com/gobuffalo/packr/v2"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	Path string
}

func (T Template) Read(name string) ([]byte, error) {
	box := packr.New(name, T.Path)
	b, e := box.Find(name)
	if e != nil {
		log.Println("template read fail", e.Error())
		return nil, e
	}
	return b, nil
}

func (T Template) Execute(w http.ResponseWriter, r *http.Request, v interface{}) {
	t := template.New("")
	text, e := T.Read("plot.html")
	if e != nil {
		log.Println("template read fail")
		return
	}
	tt, e := t.Parse(string(text))
	if e != nil {
		log.Println("template parse fail")
		return
	}
	tt.Execute(w, v)
}
