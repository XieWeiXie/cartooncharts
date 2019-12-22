package cartooncharts

import (
	"github.com/wuxiaoxiaoshen/cartoonycharts/cartoontransport"
)

type CartoonCharts struct {
	*cartoontransport.ChartsTransport
}

func NewCartoonCharts() *CartoonCharts {
	return &CartoonCharts{cartoontransport.NewChartsTransport()}
}
