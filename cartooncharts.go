package cartooncharts

import (
	"github.com/XieWeiXie/cartooncharts/cartoontransport"
)

type CartoonCharts struct {
	*cartoontransport.ChartsTransport
}

func NewCartoonCharts() *CartoonCharts {
	return &CartoonCharts{cartoontransport.NewChartsTransport()}
}
