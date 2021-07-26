package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/jirasdk"
)

var jiraapi *jirasdk.Client

func init() {
	c, err := jirasdk.NewClient("http://172.16.74.113:8080", "Jarvisbot", "12345678")
	if err != nil {
		panic(err)
	}
	jiraapi = c
}

func list(keyid string) {
	cg := jirasdk.ComponentGetOption{
		ProjectIdOrKey: keyid,
	}
	v, resp, err := jiraapi.Component.Get(&cg)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func post(p, n, d string) {
	cg := jirasdk.ComponentPostOption{
		Name:        n,
		Description: d,
		Project:     p,
	}
	v, resp, err := jiraapi.Component.Post(&cg)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	list("DEV")
	post("DEV", "Applex", "apple demo")
	list("DEV")
}
