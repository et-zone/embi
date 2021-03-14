package main

import (
	"math/rand"

	"github.com/et-zone/embi/client"
)

// const url = "http://192.168.122.139:6661"
const url = "http://127.0.0.1:6661"

func main() {
	RandDmo()
}

func RandDmo() {
	methods := []string{"GET", "POST", "DELETE", "PUT"}
	urls := []string{"http://192.168.1.220/get", "http://192.168.1.220/post", "http://192.168.1.220/delete", "http://192.168.1.220/put"}
	index := rand.Intn(3)
	codeind := rand.Intn(100)

	cli := client.InitDefaultClient()
	ctx := client.NewContext()
	if codeind <= 98 {
		ctx.UpdateContext("all", methods[index], urls[index], 200)
	} else {
		ctx.UpdateContext("all", methods[index], urls[index], 500)
	}

	cli.Dao(ctx, url)
}
