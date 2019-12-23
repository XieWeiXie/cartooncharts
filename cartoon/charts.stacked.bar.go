package cartoon

import "net/http"

func newStackedBar(t Transport) StackedBar {
	return func(title string, o ...func(request *StackedBarRequest)) func(w http.ResponseWriter, r *http.Request) {
		r := StackedBarRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Plot(t)

	}
}

type StackedBar func(title string, o ...func(request *StackedBarRequest)) func(w http.ResponseWriter, r *http.Request)

type StackedBarRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (sta StackedBarRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface StackedBarRequest
		}{
			Type:      stackedBarStackedType,
			Interface: sta,
		}
		t.Execute(w, r, v)
	}
}

func (s *StackedBar) WithTitle(title string) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		request.Title = title
	}
}
func (s *StackedBar) WithXLabel(xLabel string) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		request.XLabel = xLabel
	}
}

func (s *StackedBar) WithYLabel(yLabel string) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		request.YLabel = yLabel
	}
}

func (s *StackedBar) WithOptions(key string, value interface{}) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		if request.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (s *StackedBar) WithDataLabels(value []interface{}) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (s *StackedBar) WithDataDataSets(label string, data []interface{}) func(request *StackedBarRequest) {
	return func(request *StackedBarRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
