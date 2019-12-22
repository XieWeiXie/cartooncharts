<h1 align="center">CartoonCharts</h1>
<p align="center">
  <em>Why CartoonCharts?  Inspired by <a href="https://github.com/timqian/chart.xkcd">chart.xkcd</a>.</em>
</p>
<p align="center">
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/Author-wuxiaoxiaoshen-blue" alt="Author">
    </a>
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/progressing-15%25-red" alt="Author">
    </a>
</p>

> Cartoon Charts ...
---
> Program to an interface, not an implementation

## Install
```
go get -u -v github.com/wuxiaoxiaoshen/cartooncharts
```

## Demo

```go
package _example

import (
	"github.com/wuxiaoxiaoshen/cartoonycharts"
	"github.com/wuxiaoxiaoshen/cartoonycharts/cartoon"
	"github.com/wuxiaoxiaoshen/cartoonycharts/cartoontransport"
	"log"
	"net/http"
	"testing"
)

func TestExampleLine(t *testing.T) {
	req := cartoon.LineRequest{
		WithTitle:  cartoon.WithTitle{Title: "Monthly income of an indie developer"},
		WithXLabel: cartoon.WithXLabel{XLabel: "Month"},
		WithYLabel: cartoon.WithYLabel{YLabel: "$ Dollars"},
		WithDataCollection: cartoon.WithDataCollection{
			Data: cartoon.WithData{
				Labels: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				DataSets: []cartoon.WithDataDataSets{
					cartoon.WithDataDataSets{
						Label: "Plan",
						Data:  []interface{}{30, 70, 200, 300, 500, 800, 1500, 2900, 5000, 8000},
					},
					cartoon.WithDataDataSets{
						Label: "Reality",
						Data:  []interface{}{0, 1, 30, 70, 80, 100, 50, 80, 40, 150},
					},
				},
			},
		},
		WithOption: cartoon.WithOption{
			Options: map[string]interface{}{
				"yTickCount":     3,
				"legendPosition": "chartXkcd.config.positionType.upLeft",
			},
		},
	}
	tt := cartoontransport.NewChartsTransport()
	c := req.Plot(tt)
	http.HandleFunc("/line", c)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func TestExampleLineMethod(t *testing.T) {
	charts := cartooncharts.NewCartoonCharts()
	line := charts.Charts.Line("Monthly income of an indie developer",
		charts.Charts.Line.WithXLabel("Month"),
		charts.Charts.Line.WithYLabel("$ Dollars"),
		charts.Charts.Line.WithDataLabels([]interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}),
		charts.Charts.Line.WithDataDataSets("Plan", []interface{}{30, 70, 200, 300, 500, 800, 1500, 2900, 5000, 8000}),
		charts.Charts.Line.WithDataDataSets("Reality", []interface{}{0, 1, 30, 70, 80, 100, 50, 80, 40, 150}),
		charts.Charts.Line.WithOptions("yTickCount", 3),
		charts.Charts.Line.WithOptions("legendPosition", "chartXkcd.config.positionType.upLeft"),
	)
	http.HandleFunc("/line", line)
	log.Fatal(http.ListenAndServe(":9090", nil))

}


```

Open "http://localhost:9090/line" in your browser. You will see:

![Line](_example/line.png)

> for convenience， You'd better use TestExampleLineMethod




## TODO

- [ ] Bar
- [ ] Pie
- [ ] Radar
- [ ] StackedBar
- [ ] XY

