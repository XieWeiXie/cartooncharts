package cartoon

import (
	"net/http"
)

type ChartsInterface interface {
	Plot(t Transport) func(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Render(t Transport) func(w http.ResponseWriter, r *http.Request)
}

type Transport interface {
	Execute(w http.ResponseWriter, r *http.Request, v interface{})
	Read(name string) ([]byte, error)
}

var (
	lineStackedType       string
	barStackedType        string
	stackedBarStackedType string
	pieStackedType        string
	radarStackedType      string
)

func init() {
	lineStackedType = "Line"
	barStackedType = "Bar"
	stackedBarStackedType = "StackedBar"
	pieStackedType = "Pie"
	radarStackedType = "Radar"
}

type Charts struct {
	Line       Line
	Bar        Bar
	Pie        Pie
	Radar      Radar
	StackedBar StackedBar
	XY         XY
}

func NewCharts(t Transport) *Charts {
	return &Charts{
		Line: newLine(t),
	}
}
