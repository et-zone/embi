package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"

	"log"
	"net"
	"net/http"
	"time"

	"github.com/et-zone/embi/model"
)

const (
	MaxIdleConnsPerHost = 100
	IdleConnTimeout     = 300
	MaxIdleConns        = 1000
	MaxConnsPerHost     = 1000
	path                = "/hmsg"
)

type Client struct {
	http.Client
}

func InitDefaultClient() *Client {
	client := &Client{
		Client: http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: false,
				Proxy:             http.ProxyFromEnvironment,
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        MaxIdleConns,
				MaxConnsPerHost:     MaxConnsPerHost,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Second * 10,
		}}

	return client
}

func (this *Client) Dao(ctx *eContext, addr string) error {
	if addr == "" {
		return errors.New("addr can not nil")
	}
	if ctx.flag != true {
		return errors.New("eContext not set value")
	}

	ehttp := &model.EHttp{
		AppName:  ctx.appName,
		Method:   ctx.method,
		Ip:       ctx.ip,
		Path:     ctx.path,
		Nowtime:  ctx.nowtime.Format("2006-01-02 15:04:05"),
		Duration: int(ctx.duration),
		Code:     ctx.code,
	}

	b, err := json.Marshal(ehttp)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", addr+path, bytes.NewReader([]byte(b)))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := this.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer res.Body.Close()
	return nil
}
