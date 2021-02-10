package chart

import (
	"fmt"
	// "fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/types"
)

//饼图
func PieMap(table *Table) *charts.Pie {
	pie := charts.NewPie()

	pie.SetGlobalOptions(

		charts.WithTitleOpts(opts.Title{
			Title:    table.Title,
			Subtitle: table.Subtitle,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:  "shine",
			Width:  "1400px",
			Height: "550px",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}), //滑动
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}), //展示具体值
		//可以放大====
		// charts.WithXAxisOpts(opts.XAxis{
		// 	SplitNumber: 20,
		// }),
		// charts.WithYAxisOpts(opts.YAxis{
		// 	Scale: true,
		// }),
		// charts.WithDataZoomOpts(opts.DataZoom{
		// 	Type:       "inside",
		// 	Start:      50,
		// 	End:        100,
		// 	XAxisIndex: []int{0},
		// }),
		//可以放大====

	)

	pie.AddSeries(table.Title, generatePieItems(table))

	//展示悬停值
	pie.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "top",
		}))

	return pie
}

// generate random data for line chart
func generatePieItems(table *Table) []opts.PieData {

	lens := len(table.DimValue)

	items := make([]opts.PieData, 0)
	for i := 0; i < lens; i++ {
		fmt.Println(table.Value[table.Field[1]][i])
		items = append(items, opts.PieData{Name: table.DimValue[i], Value: table.Value[table.Field[1]][i]})

	}
	return items
}
