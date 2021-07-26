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

func list() {
	pg := jirasdk.PriorityGetOption{}
	v, resp, err := jiraapi.Priority.Get(&pg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	list()
}
