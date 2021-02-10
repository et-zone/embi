package main

import (
	"net/http"
	"time"

	"github.com/et-zone/embi/db"

	"github.com/et-zone/embi/dao"
	"github.com/et-zone/embi/router"
)

//日志监控bi
func main() {
	Init()
	httpServer := &http.Server{
		Addr:         ":6661",
		Handler:      router.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	httpServer.ListenAndServe()
}

func Init() {
	cfg := &db.EConfig{
		UserName: "root",
		PassWord: "mysql",
		Addr:     "127.0.0.1",
		Port:     3366,
		DB:       "rest",
	}
	dao.Init(cfg)
}
