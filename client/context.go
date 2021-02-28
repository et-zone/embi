package client

import (
	"context"
	"strings"
	"time"
)

type eContext struct {
	context.Context
	appName  string
	method   string
	ip       string
	path     string        //不需要ip地址
	nowtime  time.Time     //
	duration time.Duration //请求时长
	status   int           //状态吗
	flag     bool          //
}

func NewContext() *eContext {
	return &eContext{context.TODO(), "", "", "", "", time.Now(), 0, 0, false}
}

func (ctx *eContext) UpdateContext(appName string, method string, path string, code int) {
	if ctx == nil {
		return
	}
	ctx.flag = true
	plist := strings.Split(path, "?")

	if strings.Contains(plist[0], "http") || strings.Contains(plist[0], "https") {
		list := strings.SplitN(plist[0], "/", 4)
		ctx.ip = list[2]
		ctx.path = "/" + list[3]
	} else {
		list := strings.SplitN(plist[0], "/", 2)
		ctx.ip = list[0]
		if len(list) == 1 {
			ctx.path = "/"
		} else {
			ctx.path = "/" + list[1]
		}
	}
	ctx.appName = appName
	ctx.method = method
	ctx.duration = time.Since(ctx.nowtime)
	ctx.status = code

}

func isSucc(code int) int {
	if code == 200 {
		return 1
	}
	return 0
}

func (ctx *eContext) GeteContextInfo() (nowtime time.Time, appName string, ip string, mothod string, path string, duration time.Duration, code int) {
	if ctx == nil {
		t := time.Now()
		return t, "", "", "", "", 0, 0
	}
	return ctx.nowtime, ctx.appName, ctx.ip, ctx.method, ctx.path, ctx.duration, ctx.status
}
