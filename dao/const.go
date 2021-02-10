package dao

//需要改名的
var FieldTableNameMap map[string]string = map[string]string{
	"createtime":  "时间",
	"avgduration": "平均响应时长",
	"qps":         "QPS(s)",
	"number":      "请求数",
	"sumduration": "响应总时长",
	"duration":    "响应时长",
	"maxduration": "最大响应时长",
	"minduration": "最小响应时长",
	"time":        "时间",
	"qpsm":        "QPS(min)",
	"code":        "状态码",
	"dcount":      "日访问量",
	"failcount":   "失败请求",
	"succount":    "成功请求",
}
var FieldBINameMap map[string]string = map[string]string{
	"createtime":  "时间",
	"avgduration": "平均响应时长(μs)",
	"qps":         "QPS(s)",
	"number":      "请求数",
	"sumduration": "响应总时长(μs)",
	"duration":    "响应时长(μs)",
	"maxduration": "最大响应时长(μs)",
	"minduration": "最小响应时长(μs)",
	"time":        "时间",
	"qpsm":        "QPS(min)",
	"code":        "状态码",
	"dcount":      "日访问量",
	"failcount":   "失败请求",
	"succount":    "成功请求",
}

//时间相关的字段,table表格要转为时间类型
var FieldToDuration = []string{
	"avgduration",
	"sumduration",
	"duration",
	"maxduration",
	"minduration",
}

//表格中时间类型需要换类型判断
func NeedChange(field string) bool {
	for i, _ := range FieldToDuration {
		if field == FieldToDuration[i] {
			return true
		}
	}
	return false
}

//字段需要转换为名称
var FieldToValue = map[string]map[string]string{
	"succ": map[string]string{"1": "成功", "0": "失败"}, //pie饼图使用
}
