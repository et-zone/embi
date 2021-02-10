package chart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/types"
)

//多重折线图
func LineMulti(table *Table) *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(

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
	line.SetXAxis(table.DimValue)
	for i, field := range table.Field {
		if i == 0 {
			continue
		}
		line.AddSeries(field, generateLineItems(table, field))
	}
	//展示悬停值
	line.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "top",
		}))

	return line
}

// generate random data for line chart
func generateLineItems(table *Table, field string) []opts.LineData {

	values := table.Value[field]

	lens := len(table.DimValue)

	items := make([]opts.LineData, 0)
	for i := 0; i < lens; i++ {

		items = append(items, opts.LineData{Value: values[i]})
	}
	return items
}
