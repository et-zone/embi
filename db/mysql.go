package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type EConfig struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
}

//初始化gorm
func InitOrm(cfg *EConfig) *gorm.DB {

	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.UserName,
		cfg.PassWord,
		cfg.Addr,
		cfg.Port,
		cfg.DB,
	))

	if err != nil {
		fmt.Println(err.Error())
		panic("orm init error")
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(100)
	return db

}
