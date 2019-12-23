package cartoon

import "net/http"

func newBar(t Transport) Bar {
	return func(title string, o ...func(request *BarRequest)) func(w http.ResponseWriter, r *http.Request) {
		r := BarRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Plot(t)

	}
}

type Bar func(title string, o ...func(request *BarRequest)) func(w http.ResponseWriter, r *http.Request)

type BarRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (bar BarRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface BarRequest
		}{
			Type:      barStackedType,
			Interface: bar,
		}
		t.Execute(w, r, v)
	}
}

func (b *Bar) WithTitle(title string) func(request *BarRequest) {
	return func(request *BarRequest) {
		request.Title = title
	}
}
func (b *Bar) WithXLabel(xLabel string) func(request *BarRequest) {
	return func(request *BarRequest) {
		request.XLabel = xLabel
	}
}

func (b *Bar) WithYLabel(yLabel string) func(request *BarRequest) {
	return func(request *BarRequest) {
		request.YLabel = yLabel
	}
}

func (b *Bar) WithOptions(key string, value interface{}) func(request *BarRequest) {
	return func(request *BarRequest) {
		if request.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (b *Bar) WithDataLabels(value []interface{}) func(request *BarRequest) {
	return func(request *BarRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (b *Bar) WithDataDataSets(label string, data []interface{}) func(request *BarRequest) {
	return func(request *BarRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
