package cartoon

import "net/http"

func newRadar(t Transport) Radar {
	return func(title string, o ...func(request *RadarRequest)) func(w http.ResponseWriter, r *http.Request) {
		r := RadarRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Plot(t)

	}
}

type Radar func(title string, o ...func(request *RadarRequest)) func(w http.ResponseWriter, r *http.Request)

type RadarRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (radar RadarRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface RadarRequest
		}{
			Type:      radarStackedType,
			Interface: radar,
		}
		t.Execute(w, r, v)
	}
}

func (r *Radar) WithTitle(title string) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		request.Title = title
	}
}
func (r *Radar) WithXLabel(xLabel string) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		request.XLabel = xLabel
	}
}

func (r *Radar) WithYLabel(yLabel string) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		request.YLabel = yLabel
	}
}

func (r *Radar) WithOptions(key string, value interface{}) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		if request.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (r *Radar) WithDataLabels(value []interface{}) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (r *Radar) WithDataDataSets(label string, data []interface{}) func(request *RadarRequest) {
	return func(request *RadarRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
