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

func create() {
	f := jirasdk.Fields{
		Project: jirasdk.FID{
			ID: "10000",
		},
		Summary: "测试",
		Issuetype: jirasdk.FID{
			ID: "10002",
		},
		Assignee: jirasdk.FName{
			Name: "jarvisbot",
		},
		Reporter: jirasdk.FName{
			Name: "jarvisbot",
		},
		Priority: jirasdk.FID{
			ID: "3",
		},
		Labels: []string{"demo"},
	}
	ic := jirasdk.IssuePostOption{
		Fields: f,
	}
	v, resp, err := jiraapi.Issue.Post(&ic)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
	get(v.Key)
}

func get(keyid string) {
	ig := jirasdk.IssueGetOption{
		IssueIdOrKey: keyid,
	}
	v, resp, err := jiraapi.Issue.Get(&ig)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func showmeta(name string) {
	im := jirasdk.IssueMetaOption{
		ProjectKeys: name,
	}
	v, resp, err := jiraapi.Issue.Meta(&im)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	create()
	showmeta("DEV")
}
