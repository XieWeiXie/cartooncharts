package main

import (
	"github.com/XieWeiXie/cartooncharts"
	"log"
	"net/http"
)

var charts *cartooncharts.CartoonCharts

func init() {
	charts = cartooncharts.NewCartoonCharts()
}

func ExampleBar() {
	bar := charts.Charts.Bar("github stars VS patron number",
		charts.Charts.Bar.WithDataLabels([]interface{}{"github stars", "patrons"}),
		charts.Charts.Bar.WithDataDataSets("", []interface{}{100, 2}),
		charts.Charts.Bar.WithOptions("yTickCount", 2),
	)
	http.HandleFunc("/bar", bar)
}
func ExampleXY() {
	type point struct {
		X interface{} `json:"x"`
		Y interface{} `json:"y"`
	}
	xy := charts.Charts.XY("Pokemon farms",
		charts.Charts.XY.WithXLabel("Coodinate"),
		charts.Charts.XY.WithYLabel("Count"),
		charts.Charts.XY.WithDataDataSets("Pikachu", []interface{}{point{3, 10}, point{4, 122}, point{10, 100}, point{1, 2}, point{2, 4}}),
		charts.Charts.XY.WithDataDataSets("Squirtle", []interface{}{point{3, 122}, point{4, 212}, point{-3, 100}, point{1, 1}, point{1.5, 12}}),
		charts.Charts.XY.WithOptions("xTickCount", 5),
		charts.Charts.XY.WithOptions("yTickCount", 5),
		charts.Charts.XY.WithOptions("legendPosition", "chartXkcd.config.positionType.upRight"),
		charts.Charts.XY.WithOptions("showLine", false),
		charts.Charts.XY.WithOptions("timeFormat", "undefined"),
		charts.Charts.XY.WithOptions("dotSize", 1),
	)
	http.HandleFunc("/xy", xy)
}
func ExampleStackedBar() {
	stackedBar := charts.Charts.StackedBar("Issues and PR Submissions",
		charts.Charts.StackedBar.WithXLabel("Month"),
		charts.Charts.StackedBar.WithYLabel("Count"),
		charts.Charts.StackedBar.WithDataLabels([]interface{}{"Jan", "Feb", "Mar", "April", "May"}),
		charts.Charts.StackedBar.WithDataDataSets("Issues", []interface{}{12, 19, 11, 29, 17}),
		charts.Charts.StackedBar.WithDataDataSets("PRs", []interface{}{3, 5, 2, 4, 1}),
		charts.Charts.StackedBar.WithDataDataSets("Merges", []interface{}{2, 3, 0, 1, 1}),
	)
	http.HandleFunc("/stackedBar", stackedBar)
}
func ExampleLine() {
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

}
func ExamplePie() {
	pie := charts.Charts.Pie("What Tim made of",
		charts.Charts.Pie.WithDataLabels([]interface{}{"a", "b", "e", "f", "g"}),
		charts.Charts.Pie.WithDataDataSets("", []interface{}{500, 200, 80, 90, 100}),
		charts.Charts.Pie.WithOptions("innerRadius", 0.5),
		charts.Charts.Pie.WithOptions("legendPosition", "chartXkcd.config.positionType.upRight"),
	)
	http.HandleFunc("/pie", pie)
}
func ExampleRadar() {
	radar := charts.Charts.Radar("Letters in random words",
		charts.Charts.Radar.WithDataLabels([]interface{}{"c", "h", "a", "r", "t"}),
		charts.Charts.Radar.WithDataDataSets("ccharrrt", []interface{}{2, 1, 1, 3, 1}),
		charts.Charts.Radar.WithDataDataSets("chhaart", []interface{}{1, 2, 2, 1, 1}),
		charts.Charts.Radar.WithOptions("showLegend", true),
		charts.Charts.Radar.WithOptions("dotSize", 0.8),
		charts.Charts.Radar.WithOptions("showLabels", true),
		charts.Charts.Radar.WithOptions("legendPosition", "chartXkcd.config.positionType.upRight"),
	)
	http.HandleFunc("/radar", radar)
}

func main() {

	ExampleBar()
	ExampleXY()
	ExampleStackedBar()
	ExampleLine()
	ExamplePie()
	ExampleRadar()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
