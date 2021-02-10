package dao

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/et-zone/embi/db"
	"github.com/et-zone/embi/model"
	"github.com/jinzhu/gorm"
)

var edb *gorm.DB

func Init(cfg *db.EConfig) {

	edb = db.InitOrm(cfg)
	//, &model.ETarget{}
	edb.AutoMigrate(&model.EHttp{})
}

func InsertEhttp(eh *model.EHttp) error {
	return edb.Create(eh).Error
}

func InsertEtarget(et *model.ETarget) error {
	return edb.Create(et).Error
}

func GetEhttp() model.EHttp {
	eh := model.EHttp{}
	edb.First(&eh)
	return eh
}

func Query(sql string) (map[string]BaseStruct, []string) {

	rows, err := edb.DB().Query(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	rolsType, _ := rows.ColumnTypes()
	cols, _ := rows.Columns() //返回所有列
	// fmt.Println("字段:", cols)
	// fmt.Println(*rolsType[0])

	ret := map[string]BaseStruct{}
	for i, rol := range cols {
		// fmt.Println("字段类型", rolsType[i].DatabaseTypeName())
		if strings.Contains(rolsType[i].DatabaseTypeName(), "INT") {
			ret[rol] = BaseStruct{
				Type:      "int",
				FieldName: rol,
				Values:    &[]interface{}{},
			}
		} else if rolsType[i].DatabaseTypeName() == "FLOAT" || rolsType[i].DatabaseTypeName() == "DOUBLE" || rolsType[i].DatabaseTypeName() == "DECIMAL" {
			ret[rol] = BaseStruct{
				Type:      "float",
				FieldName: rol,
				Values:    &[]interface{}{},
			}
		} else {
			ret[rol] = BaseStruct{
				Type:      "string",
				FieldName: rol,
				Values:    &[]interface{}{},
			}
		}
	}

	// scans := make([]interface{}, len(cols)) //这里表示一行填充数据
	values := make([][]byte, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		rows.Scan(scanArgs...) //填充数据
		for i, v := range values {

			switch ret[cols[i]].Type {
			case "int":
				d, _ := strconv.Atoi(string(v))
				*ret[cols[i]].Values = append(*ret[cols[i]].Values, d)
			case "string":
				*ret[cols[i]].Values = append(*ret[cols[i]].Values, string(v))
			case "float":
				f, _ := strconv.ParseFloat(string(v), 64)
				*ret[cols[i]].Values = append(*ret[cols[i]].Values, f)
			}

		}
	}
	// _, _ := json.Marshal(&ret)
	// fmt.Println(string(b))
	return ret, cols

}

type BaseStruct struct {
	Type      string
	FieldName string
	Values    *[]interface{}
}
