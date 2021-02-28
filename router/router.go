package router

import (
	"github.com/et-zone/embi/control"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.LoadHTMLGlob("template/*")
	// r.Static("/static", "./static")
	r.Use(control.Cors())
	r.GET("/", control.NowMsg)
	r.POST("/ping", control.Ping)
	r.POST("/hmsg", control.HInsert)
	// r.GET("/hmsg/search", control.HGet)
	// r.GET("/hmsg/searcht", control.Htable)
	// r.POST("/tmsg", control.TPost)
	// r.GET("/tmsg/search", control.TGet)
	r.GET("/history", control.HistoryMsg)
	return r
}

// curl -X POST --data '' 'http://127.0.0.1:6661/ping'
