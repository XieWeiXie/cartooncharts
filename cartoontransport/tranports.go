package cartoontransport

import (
	"github.com/wuxiaoxiaoshen/cartoonycharts/cartoon"
	"net/http"
)

type Interface interface {
	Execute(w http.ResponseWriter, r *http.Request, v interface{})
	Read(name string) ([]byte, error)
}

type ChartsTransport struct {
	Template Interface
	Charts   *cartoon.Charts
}

func (C ChartsTransport) Execute(w http.ResponseWriter, r *http.Request, v interface{}) {
	C.Template.Execute(w, r, v)
}
func (C ChartsTransport) Read(name string) ([]byte, error) {
	return C.Template.Read(name)
}

func NewChartsTransport() *ChartsTransport {
	t := Template{Path: "./template"}
	return &ChartsTransport{
		Template: t,
		Charts:   cartoon.NewCharts(t),
	}
}
