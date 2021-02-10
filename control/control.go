package control

import (
	"net/http"

	"fmt"
	"strconv"
	"time"

	"github.com/et-zone/embi/chart"
	"github.com/et-zone/embi/dao"
	"github.com/gin-gonic/gin"
)

// "github.com/et-zone/embi/model"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

const ()

//table 标题
func baseTableChangeName(table *chart.Table) {
	for i, field := range table.Field {

		if dao.FieldTableNameMap[field] != "" {
			tmp := table.Value[field]
			delete(table.Value, field)
			table.Value[dao.FieldTableNameMap[field]] = tmp
			table.Field[i] = dao.FieldTableNameMap[field]
		}
	}
}

//BI 图标题
func baseBIChangeName(table *chart.Table) {
	for i, field := range table.Field {

		if dao.FieldBINameMap[field] != "" {
			tmp := table.Value[field]
			delete(table.Value, field)
			table.Value[dao.FieldBINameMap[field]] = tmp
			table.Field[i] = dao.FieldBINameMap[field]
		}
	}
}

//table 表格类型数据处理
func baseTableAddmsg(table *chart.Table, val map[string]dao.BaseStruct, fields []string) {
	dimvalue := []string{}
	// fmt.Println(fields[0])
	dimType := val[fields[0]].Type
	for _, v := range *val[fields[0]].Values {
		if dimType == "int" {

			dimvalue = append(dimvalue, strconv.Itoa(v.(int)))
		} else if dimType == "float" {
			dimvalue = append(dimvalue, strconv.FormatFloat(v.(float64), 'f', 4, 64))
		} else {
			dimvalue = append(dimvalue, v.(string))
		}

	}
	table.DimValue = dimvalue
	for k, field := range val {

		if dao.NeedChange(k) {
			t := []interface{}{}
			for _, v := range *field.Values {
				if field.Type == "float" {
					data, err := strconv.Atoi(strconv.FormatFloat(v.(float64), 'f', 0, 64))
					if err != nil {
						fmt.Println(err.Error())
					}
					t = append(t, fmt.Sprintf("%v", time.Duration(data)))
				} else {
					t = append(t, fmt.Sprintf("%v", time.Duration(v.(int))))
				}

			}
			table.Value[k] = t
		} else {
			table.Value[k] = *field.Values
		}

	}

	table.Len = len(table.DimValue)

}

//BI    类型数据处理
func baseBIAddmsg(table *chart.Table, val map[string]dao.BaseStruct, fields []string) {
	dimvalue := []string{}
	// fmt.Println(fields[0])
	dimType := val[fields[0]].Type
	for _, v := range *val[fields[0]].Values {
		if dimType == "int" {
			dimvalue = append(dimvalue, strconv.Itoa(v.(int)))
		} else if dimType == "float" {
			dimvalue = append(dimvalue, strconv.FormatFloat(v.(float64), 'f', 4, 64))
		} else {
			dimvalue = append(dimvalue, v.(string))
		}

	}
	table.DimValue = dimvalue
	for k, field := range val {

		if dao.NeedChange(k) {
			t := []interface{}{}
			for _, v := range *field.Values {
				if field.Type == "float" {
					data := strconv.FormatFloat(v.(float64)/1000, 'f', 2, 64)
					f, _ := strconv.ParseFloat(data, 64)
					t = append(t, f)

				} else {
					f, _ := strconv.ParseFloat(strconv.Itoa(v.(int)/1000), 64)
					t = append(t, f)
				}
				t = append(t, float64(v.(int))/1000)
			}
			table.Value[k] = t
		} else {
			table.Value[k] = *field.Values
		}

	}
	table.Len = len(table.DimValue)

}

//维度字段需要转换为名称，如是否成功0和1 图表使用
func dimValueChange(table *chart.Table) bool {
	dimField := table.Field[0]

	if dao.FieldToValue[dimField] == nil {
		return false
	}
	mv := dao.FieldToValue[dimField]
	for i, v := range table.DimValue {
		table.DimValue[i] = mv[v]
	}
	return true
}

//表格的所有字段需要转换为名称，如是否成功0和1
func tableValueChange(table *chart.Table) bool {

	for i, field := range table.Field {
		if i == 0 {
			dimValueChange(table)
		} else {
			if dao.FieldToValue[field] == nil {
				continue
			} else {
				dlist := []interface{}{}
				mv := dao.FieldToValue[field]
				for _, v := range table.Value[field] {
					dlist = append(dlist, mv[fmt.Sprintf("%v", v)])
				}
				table.Value[field] = dlist
			}
		}

	}

	return true
}
