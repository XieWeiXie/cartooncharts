package cartoon

import (
	"net/http"
)

func newLine(t Transport) Line {
	return func(title string, o ...func(request *LineRequest)) func(w http.ResponseWriter, r *http.Request) {
		line := LineRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&line)
		}
		return line.Plot(t)
	}
}

type Line func(title string, o ...func(request *LineRequest)) func(w http.ResponseWriter, r *http.Request)

type LineRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (L LineRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface LineRequest
		}{
			Type:      lineStackedType,
			Interface: L,
		}
		t.Execute(w, r, v)
	}
}

func (L LineRequest) Save(name string, transport Transport) bool {
	return true
}

func (L LineRequest) Render(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return L.Plot(t)
}

func (L *Line) WithTitle(title string) func(request *LineRequest) {
	return func(request *LineRequest) {
		request.Title = title
	}
}

func (L *Line) WithXLabel(xLabel string) func(request *LineRequest) {
	return func(request *LineRequest) {
		request.XLabel = xLabel
	}
}

func (L *Line) WithYLabel(yLabel string) func(request *LineRequest) {
	return func(request *LineRequest) {
		request.YLabel = yLabel
	}
}

func (L *Line) WithOptions(key string, value interface{}) func(request *LineRequest) {
	return func(request *LineRequest) {
		if request.WithOption.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (L *Line) WithDataLabels(value []interface{}) func(request *LineRequest) {
	return func(request *LineRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (L *Line) WithDataDataSets(label string, data []interface{}) func(request *LineRequest) {
	return func(request *LineRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
