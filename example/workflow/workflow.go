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

func get() {
	wgo := jirasdk.WorkflowGetOption{}
	v, resp, err := jiraapi.Workflow.Get(&wgo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// func getv2(name string) {
// 	wgo := jirasdk.WorkflowGetOption{
// 		WorkflowName: name,
// 	}
// 	v, resp, err := jiraapi.Workflow.Get(&wgo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(resp.StatusCode)
// 	spew.Dump(v)
// }

func main() {
	get()
	// getv2("jira")
}
