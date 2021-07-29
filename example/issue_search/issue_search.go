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

func search(jql string) {
	iso := jirasdk.IssueSearchOption{
		Jql: jql,
	}
	v, resp, err := jiraapi.IssueSearch.Get(&iso)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	search("project = DEV AND status = Open")
	search("status = Open")
}
