package chart

import (
	"fmt"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	tableCSS     = "<style type=\"text/css\">table.etable {font-size:12px;color:#333333;width:100%;border-width: 1px;border-color: #ebab3a;border-collapse: collapse;border:1px}table.etable th {font-size:12px;background-color:#e6983b;border-width: 1px;padding: 8px;border-style: solid;border-color: #ebab3a;text-align:left;}table.etable tr {background-color:#ffffff;}table.etable td {font-size:12px;border-width: 1px;padding: 8px;border-style: solid;border-color: #ebab3a;} table caption {font-size:34px;} </style>"
	tableTypeID  = "etable"
	DefaultTable = "<table class=\"etable\"><tr>#table_title#</tr> #table_info#</table>"
	tableTh      = "<th>%v</th>"
	tableTr      = "<tr>%s</tr>"
	tableTd      = "<td>%v</td>"
	tableTitle   = "<caption  align=\"center\">%v</caption>"
)

func FTable(table *Table) *charts.Bar {
	tableMsg := DefaultTable
	titleinfo := fmt.Sprintf(tableTitle, table.Title)
	for _, dim := range table.Field {
		titleinfo += fmt.Sprintf(tableTh, dim)
	}
	tableMsg = strings.Replace(tableMsg, "#table_title#", titleinfo, 1)

	tablerows := ""

	for i := 0; i < table.Len; i++ {
		row := ""

		s := len(table.Field)
		for j := 0; j < s; j++ {
			if j == 0 {
				row += fmt.Sprintf(tableTd, table.DimValue[i])
			} else {
				row += fmt.Sprintf(tableTd, table.Value[table.Field[j]][i])
			}
		}
		row = fmt.Sprintf(tableTr, row)
		// fmt.Println(row)
		tablerows += row

	}
	tableMsg = strings.Replace(tableMsg, "#table_info#", tablerows, 1)
	if table.TbID == "" {
		tableMsg += tableCSS
	} else {
		strings.ReplaceAll(tableCSS, "etable", table.TbID)
	}

	hostBar := charts.NewBar()
	hostBar.SetGlobalOptions(charts.WithLegendOpts(opts.Legend{Height: "100px"}),
		charts.WithTitleOpts(opts.Title{
			Title:    table.Title,
			Subtitle: table.Subtitle,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:  "shine",
			Width:  "1100px",
			Height: "550px",
		}),
	)
	if table.TbID == "" {
		hostBar.ChartID = tableTypeID
		hostBar.AddJSFuncs(fmt.Sprintf("document.getElementById('%v').innerHTML='%v';", tableTypeID, tableMsg))

	} else {
		hostBar.ChartID = table.TbID
		hostBar.AddJSFuncs(fmt.Sprintf("document.getElementById('%v').innerHTML='%v';", table.TbID, tableMsg))

	}

	return hostBar
}
