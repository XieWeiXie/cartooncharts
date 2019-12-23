package cartoon

type WithTitle struct {
	Title string `json:"title"`
}

type WithXLabel struct {
	XLabel string `json:"xLabel,omitempty"`
}

type WithYLabel struct {
	YLabel string `json:"yLabel,omitempty"`
}

type WithDataCollection struct {
	Data WithData `json:"data"`
}

type WithData struct {
	Labels   []interface{}      `json:"labels"`
	DataSets []WithDataDataSets `json:"datasets"`
}

type WithDataDataSets struct {
	Label string        `json:"label"`
	Data  []interface{} `json:"data"`
}

type WithOption struct {
	Options map[string]interface{} `json:"options,omitempty"`
}

func (W WithOption) Set(key string, value interface{}) {
	W.Options[key] = value
}

func (W WithOption) Get(key string) interface{} {
	if W.Options == nil {
		return ""
	}
	vs := W.Options[key]
	return vs
}

func (W WithOption) Del(key string) {
	delete(W.Options, key)
}
