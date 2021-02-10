package control

import (
	"fmt"

	"github.com/et-zone/embi/chart"
	"github.com/et-zone/embi/dao"
	"github.com/et-zone/embi/model"
	"github.com/gin-gonic/gin"

	// "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
)

func HInsert(c *gin.Context) {

	h := &model.EHttp{}
	err := c.Bind(h)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"message": "ok",
		})
		return
	}
	err = dao.InsertEhttp(h)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(501, gin.H{
			"message": "false",
		})
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
	return

}

func HGet(c *gin.Context) {

	// method := c.GetQuery("m")
	// h := dao.GetEhttp()
	// table := &chart.Table{
	// 	Title:    "",
	// 	Subtitle: "",
	// 	Field:    []string{"日期", "指标1", "指标2"}, //
	// 	Value: map[string][]interface{}{
	// 		"指标1": []interface{}{1, 4, 7, 2, 6, 8, 11, 10},
	// 		"指标2": []interface{}{3, 5, 7, 9, 11, 13, 15, 17},
	// 	},
	// 	DimValue: []string{"2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08"},
	// }
	// table.Len = len(table.DimValue)
	// line := chart.LineMulti(table)
	// line.Render(c.Writer)
	val, fields := dao.Query("select id,duration,code from e_http")
	table := &chart.Table{
		Title:    "",
		Subtitle: "",
		Field:    fields, //
		Value:    map[string][]interface{}{},
		// DimValue: ,
	}
	baseBIAddmsg(table, val, fields)
	baseBIChangeName(table)
	line := chart.LineMulti(table)
	line.Render(c.Writer)
}

func Htable(c *gin.Context) {
	// method := c.GetQuery("m")
	// h := dao.GetEhttp()

	// table := &chart.Table{
	// 	Title:    "",
	// 	Subtitle: "",
	// 	Field:    []string{"日期", "指标1", "指标2"}, //
	// 	Value: map[string][]interface{}{
	// 		"指标1": []interface{}{1, 4, 7, 2, 6, 8, 11, 10},
	// 		"指标2": []interface{}{3, 5, 7, 9, 11, 13, 15, 17},
	// 	},
	// 	DimValue: []string{"2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08"},
	// }

	// val, fields := dao.Query("select * from e_http")
	val, fields := dao.Query("select id,succ,code from e_http")
	table := &chart.Table{
		Title:    "",
		Subtitle: "",
		Field:    fields, //
		Value:    map[string][]interface{}{},
		// DimValue: ,
	}
	baseTableAddmsg(table, val, fields)
	baseTableChangeName(table)
	line := chart.FTable(table)
	line.Render(c.Writer)

}

func HGetHtml(c *gin.Context) {
	page := components.NewPage()
	page.AddCharts()
	page.Render(c.Writer)
	// c.HTML(200, "index.tmpl", gin.H{"title": "posts/index"})

}

func NowMsg(c *gin.Context) {

	page := components.NewPage()
	page.AddCharts(
		getTableNow(),
		getPidNow(),
		getLineLastTime(),
		getLineNow(),
		getTableNowDetail(),
	)
	page.Render(c.Writer)

}

func HistoryMsg(c *gin.Context) {

	//page操作---近30天
	page := components.NewPage()
	page.AddCharts(
		getTableMonth(),
		getLineMonth(),
		getLineMonthCmp(),
	)
	page.Render(c.Writer)

}
