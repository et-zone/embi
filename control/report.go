package control

import (
	"fmt"
	// "fmt"
	"time"

	"github.com/et-zone/embi/chart"
	"github.com/et-zone/embi/dao"
	"github.com/go-echarts/go-echarts/v2/charts"
)

//获取当天图表信息
func getTableNow() *charts.Bar {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := "SELECT DATE( s_time ) createtime, ROUND(AVG(duration)) avgduration, count(id) dcount, max( duration) maxduration, min( duration) minduration FROM e_http WHERE s_time BETWEEN '" + t[:10] + " 00:00:00" + "' AND '" + t[:10] + " 23:59:59" + "' GROUP BY createtime"

	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "今日基本信息",
		Subtitle: "今日基本信息",
		Field:    fields, //
		Value:    map[string][]interface{}{},
		// DimValue: ,
	}
	baseTableAddmsg(table, val, fields)
	tableValueChange(table)
	baseTableChangeName(table)
	tableChart := chart.FTable(table)
	return tableChart

}

//获取当天图表详细信息
func getTableNowDetail() *charts.Bar {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := "SELECT DATE_FORMAT(s_time, '%Y-%m-%d') AS time, method ,path,CASE status WHEN 200 THEN 1 ELSE 0 END succ ,count(id) AS dcount FROM e_http WHERE s_time BETWEEN '" + t[:10] + " 00:00:00" + "' AND '" + t[:10] + " 23:59:59" + "' GROUP BY time,method,path,succ ORDER BY time"
	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "今日详情",
		Subtitle: "今日详情",
		Field:    fields, //
		Value:    map[string][]interface{}{},
		// DimValue: ,
	}
	table.TbID = "detail"
	baseTableAddmsg(table, val, fields)
	tableValueChange(table)
	baseTableChangeName(table)
	tableChart := chart.FTable(table)
	return tableChart

}

//获取当QPS
func getLineNow() *charts.Line {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := "SELECT DATE_FORMAT(s_time, '%Y-%m-%d %H:%i:00') AS time, count(id) AS qpsm FROM e_http WHERE s_time BETWEEN '" + t[:10] + " 00:00:00" + "' AND '" + t[:10] + " 23:59:59" + "' GROUP BY time ORDER BY time"

	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "今日QPS",
		Subtitle: "req/min",
		Field:    fields, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table, val, fields)
	baseBIChangeName(table)
	tableChart := chart.LineMulti(table)
	return tableChart

}

//近12小时的访问量QPS
func getLineLastTime() *charts.Line {
	now := time.Now().Format("2006-01-02 15:04:05")
	bef := time.Now().Add(time.Second * 60 * 60 * -12).Format("2006-01-02 15:04:05")
	sql := "SELECT s_time AS time, count(id) AS qps FROM e_http WHERE s_time BETWEEN '" + bef[:10] + " 00:00:00" + "' AND '" + now[:10] + " 23:59:59" + "' GROUP BY time ORDER BY time"

	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "近12小时QPS",
		Subtitle: "req/s",
		Field:    fields, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table, val, fields)
	baseBIChangeName(table)
	tableChart := chart.LineMulti(table)
	return tableChart

}

//获取今请求成功失败比例
func getPidNow() *charts.Pie {
	t := time.Now().Format("2006-01-02 15:04:05")
	sql := "SELECT CASE status WHEN 200 THEN 1 ELSE 0 END succ ,count(id) succnum FROM e_http where s_time BETWEEN '" + t[:10] + " 00:00:00" + "' AND '" + t[:10] + " 23:59:59" + "' GROUP BY succ "

	val, fields := dao.Query(sql)
	if len(fields) != 2 {
		fmt.Println("查询字段太多，不支持饼图BI ")
		return nil
	}
	table := &chart.Table{
		Title:    "今日请求饼图",
		Subtitle: "",
		Field:    fields, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table, val, fields)
	dimValueChange(table)
	baseBIChangeName(table)
	pieChart := chart.PieMap(table)
	return pieChart
}

//以下是周期的======================
//近30天的表格信息
func getTableMonth() *charts.Bar {
	now := time.Now().Format("2006-01-02 15:04:05")
	bef := time.Now().Add(time.Second * 60 * 60 * 24 * -30).Format("2006-01-02 15:04:05")
	sql := "SELECT DATE( s_time ) createtime, ROUND(AVG(duration)) avgduration, count(id) dcount, max( duration) maxduration, min( duration) minduration FROM e_http WHERE s_time BETWEEN '" + bef[:10] + " 00:00:00" + "' AND '" + now[:10] + " 23:59:59" + "' GROUP BY createtime"

	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "近30天概览",
		Subtitle: "近30天概览",
		Field:    fields, //
		Value:    map[string][]interface{}{},
		// DimValue: ,
	}
	baseTableAddmsg(table, val, fields)
	tableValueChange(table)
	baseTableChangeName(table)
	tableChart := chart.FTable(table)
	return tableChart

}

//近30天的访问量
func getLineMonth() *charts.Line {
	now := time.Now().Format("2006-01-02 15:04:05")
	bef := time.Now().Add(time.Second * 60 * 60 * 24 * -30).Format("2006-01-02 15:04:05")
	sql := "SELECT DATE_FORMAT(s_time, '%Y-%m-%d') AS time, count(id) AS dcount FROM e_http WHERE s_time BETWEEN '" + bef[:10] + " 00:00:00" + "' AND '" + now[:10] + " 23:59:59" + "' GROUP BY time ORDER BY time"

	val, fields := dao.Query(sql)
	table := &chart.Table{
		Title:    "近30天访问量",
		Subtitle: "近30天访问量",
		Field:    fields, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table, val, fields)
	baseBIChangeName(table)
	tableChart := chart.LineMulti(table)
	return tableChart

}

//近30天曲率对比
func getLineMonthCmp() *charts.Line {
	nowtime := time.Now()
	now := nowtime.Format("2006-01-02 15:04:05")
	beftime := time.Now().Add(time.Second * 60 * 60 * 24 * -30)
	bef := beftime.Format("2006-01-02 15:04:05")

	sqlsucc := "SELECT DATE_FORMAT(s_time, '%Y-%m-%d') AS time, count(id) AS succount FROM e_http WHERE status=200 AND s_time BETWEEN'" + bef[:10] + " 00:00:00" + "' AND '" + now[:10] + " 23:59:59" + "' GROUP BY time ORDER BY time"
	val, fields := dao.Query(sqlsucc)
	table := &chart.Table{
		Title:    "近30天曲率对比",
		Subtitle: "近30天曲率对比",
		Field:    fields, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table, val, fields)
	AddValueIfNotExitWithTable(table, beftime, nowtime, "d")

	sqlfail := "SELECT DATE_FORMAT(s_time, '%Y-%m-%d') AS time, count(id) AS failcount FROM e_http WHERE status<>200 AND s_time BETWEEN'" + bef[:10] + " 00:00:00" + "' AND '" + now[:10] + " 23:59:59" + "' GROUP BY time ORDER BY time"

	val2, fields2 := dao.Query(sqlfail)
	table2 := &chart.Table{
		Title:    "近30天曲率对比",
		Subtitle: "近30天曲率对比",
		Field:    fields2, //
		Value:    map[string][]interface{}{},
	}
	baseBIAddmsg(table2, val2, fields2)

	AddValueIfNotExitWithTable(table2, beftime, nowtime, "d")

	table.Field = append(table.Field, table2.Field[1:]...)
	for _, field := range table2.Field[1:] {
		table.Value[field] = table2.Value[field]
	}

	fmt.Println(table.Field)
	baseBIChangeName(table)
	tableChart := chart.LineMulti(table)
	return tableChart

}

//补充所有的date
func dateadd(start time.Time, end time.Time, format string) []string {
	list := []string{}
	switch format {
	case "d":
		for {
			if end.Format("2006-01-02 15:04:05") <= start.Format("2006-01-02 15:04:05") {

				break
			} else {
				tmp := start.Add(time.Second * 24 * 3600).Format("2006-01-02 15:04:05")
				list = append(list, tmp[:10])
				start = start.Add(time.Second * 24 * 3600)

			}
		}
	}
	// fmt.Println(list)
	return list

}

func AddValueIfNotExitWithTable(table *chart.Table, start time.Time, end time.Time, format string) {
	list := dateadd(start, end, format)
	if len(list) == 0 {
		return
	}
	index := 0
	flag := false

	value := map[string][]interface{}{}
	for i, v := range table.Field {
		if i != 0 {
			value[v] = []interface{}{}
		}
	}
	newDimValue := []string{}
	for _, newdim := range list {

		// fmt.Println(table.DimValue)

		for i := index; i < len(table.DimValue); i++ {
			if newdim != table.DimValue[i] {
				newDimValue = append(newDimValue, newdim)
				for vi, _ := range value {
					value[vi] = append(value[vi], 0)
				}
				break
			} else {
				newDimValue = append(newDimValue, table.DimValue[i])
				for vi, _ := range value {

					value[vi] = append(value[vi], table.Value[vi][index])
				}
				index = index + 1
				break

			}
		}

		if index == len(table.DimValue) {
			if flag == false {
				flag = true
				continue
			} else {
				newDimValue = append(newDimValue, newdim)
				for vi, _ := range value {
					value[vi] = append(value[vi], 0)
				}
			}
		}

	}

	// 1 2 3 4 5 6 7
	// 1 4 6 7 9
	table.DimValue = newDimValue
	table.Value = value

}
