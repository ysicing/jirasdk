// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

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

func show() {
	slo := jirasdk.SchemeListOptions{}
	v, resp, err := jiraapi.Scheme.List(&slo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func get(id string) {
	sgo := jirasdk.SchemeGetOptions{
		SchemeId: id,
	}
	v, resp, err := jiraapi.Scheme.Get(&sgo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	show()
	get("10311")
	get("10312")
}
