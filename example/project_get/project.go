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
	pg := jirasdk.ProjectGetOption{
		IncludeArchived: true,
	}
	v, resp, err := jiraapi.Project.Get(&pg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func search(name string) {
	pg := jirasdk.ProjectSearchOption{
		IncludeArchived: true,
		Search:          name,
		MaxResults:      1,
	}
	v, resp, err := jiraapi.Project.Search(&pg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func projectissuetype() {
	pitgo := jirasdk.ProjectIssueTypeGetOption{
		Project:   "DEV",
		Issuetype: "10000",
	}
	v, resp, err := jiraapi.Project.IssueTypeFields(&pitgo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
	fmt.Sprintln("------------------------------")
	v1, resp1, err := jiraapi.Project.IssueTypeWorkflow(&pitgo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp1.StatusCode)
	spew.Dump(v1)
}

func main() {
	list()
	search("DEV")
	projectissuetype()
}
