package cartoon

import "net/http"

func newXY(t Transport) XY {
	return func(title string, o ...func(request *XYRequest)) func(w http.ResponseWriter, r *http.Request) {
		r := XYRequest{
			WithTitle: WithTitle{Title: title},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Plot(t)

	}
}

type XY func(title string, o ...func(request *XYRequest)) func(w http.ResponseWriter, r *http.Request)

type XYRequest struct {
	WithTitle
	WithXLabel
	WithYLabel
	WithDataCollection
	WithOption
}

func (XY XYRequest) Plot(t Transport) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Type      string
			Interface XYRequest
		}{
			Type:      xyStackedType,
			Interface: XY,
		}
		t.Execute(w, r, v)
	}
}

func (XY XY) WithTitle(title string) func(request *XYRequest) {
	return func(request *XYRequest) {
		request.Title = title
	}
}
func (XY XY) WithXLabel(xLabel string) func(request *XYRequest) {
	return func(request *XYRequest) {
		request.XLabel = xLabel
	}
}

func (XY XY) WithYLabel(yLabel string) func(request *XYRequest) {
	return func(request *XYRequest) {
		request.YLabel = yLabel
	}
}

func (XY XY) WithOptions(key string, value interface{}) func(request *XYRequest) {
	return func(request *XYRequest) {
		if request.Options == nil {
			request.WithOption.Options = make(map[string]interface{})
		}
		request.WithOption.Set(key, value)
	}
}

func (XY *XY) WithDataLabels(value []interface{}) func(request *XYRequest) {
	return func(request *XYRequest) {
		request.WithDataCollection.Data.Labels = value
	}
}
func (XY *XY) WithDataDataSets(label string, data []interface{}) func(request *XYRequest) {
	return func(request *XYRequest) {
		values := WithDataDataSets{Label: label, Data: data}
		request.WithDataCollection.Data.DataSets = append(request.WithDataCollection.Data.DataSets, values)
	}
}
