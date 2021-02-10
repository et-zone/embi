package chart

import (
	_ "github.com/go-echarts/go-echarts/v2/charts"
	_ "github.com/go-echarts/go-echarts/v2/opts"
)

type Table struct {
	Title    string
	Subtitle string
	Field    []string                 //字段
	Value    map[string][]interface{} //value
	DimValue []string                 //维度的值
	Len      int
	TbID     string
}
