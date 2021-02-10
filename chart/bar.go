package chart

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/types"
)

//多重柱状图
func BarMulti(table *Table) *charts.Bar {
	bar := charts.NewBar()

	bar.SetGlobalOptions(

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
	bar.SetXAxis(table.DimValue)
	for i, field := range table.Field {
		if i == 0 {
			continue
		}
		bar.AddSeries(field, generateBarItems(table, field))
	}
	//展示悬停值
	bar.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "top",
		}))

	return bar
}

// generate random data for line chart
func generateBarItems(table *Table, field string) []opts.BarData {
	fmt.Println(field)
	values := table.Value[field]
	fmt.Println(values)
	lens := len(table.DimValue)
	fmt.Println(lens)
	items := make([]opts.BarData, 0)
	for i := 0; i < lens; i++ {

		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}
