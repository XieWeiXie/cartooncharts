package cartoon

import "net/http"

func newPie(t Transport) Pie {
	return func(title string, o ...func(request *PieRequest)) func(w http.ResponseWriter, r *http.Request) {
		r := PieRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Plot(t)

	}
}

type Pie func(title string, o ...func(request *PieRequest)) func(w http.ResponseWriter, r *http.Request)

type PieRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (bar PieRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface PieRequest
		}{
			Type:      pieStackedType,
			Interface: bar,
		}
		t.Execute(w, r, v)
	}
}

func (p *Pie) WithTitle(title string) func(request *PieRequest) {
	return func(request *PieRequest) {
		request.Title = title
	}
}
func (p *Pie) WithXLabel(xLabel string) func(request *PieRequest) {
	return func(request *PieRequest) {
		request.XLabel = xLabel
	}
}

func (p *Pie) WithYLabel(yLabel string) func(request *PieRequest) {
	return func(request *PieRequest) {
		request.YLabel = yLabel
	}
}

func (p *Pie) WithOptions(key string, value interface{}) func(request *PieRequest) {
	return func(request *PieRequest) {
		if request.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (p *Pie) WithDataLabels(value []interface{}) func(request *PieRequest) {
	return func(request *PieRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (p *Pie) WithDataDataSets(label string, data []interface{}) func(request *PieRequest) {
	return func(request *PieRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
