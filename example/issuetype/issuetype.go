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
	itl := jirasdk.IssueTypeListOption{}
	v, resp, err := jiraapi.IssueType.List(&itl)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func get(id string) {
	itg := jirasdk.IssueTypeGetOption{
		ID: id,
	}
	v, resp, err := jiraapi.IssueType.Get(&itg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func create(name, desc string, keytype jirasdk.IssueTypeType) {
	itp := jirasdk.IssueTypePostOption{
		Name:        name,
		Description: desc,
		Type:        keytype,
	}
	v, resp, err := jiraapi.IssueType.Post(&itp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
	get(v.ID)
}

func main() {
	create("标准xx", "标准xx", jirasdk.Standard)
	create("子任务xx", "子任务xx", jirasdk.Subtask)
	list()
}
