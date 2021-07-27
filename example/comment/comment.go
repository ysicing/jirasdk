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

func commentlist(keyid string) {
	ig := jirasdk.IssueCommentGetOption{
		IssueIdOrKey: keyid,
	}
	v, resp, err := jiraapi.Comment.Get(&ig)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func commentadd(keyid, msg string, readyonly bool) {
	ig := jirasdk.IssueCommentPostOption{
		IssueIdOrKey: keyid,
		Body:         msg,
	}
	if readyonly {
		ig.Visibility = jirasdk.IssueVisibility{
			Type:  "role",
			Value: "Administrators",
		}
	}
	v, resp, err := jiraapi.Comment.Post(&ig)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	commentlist("DEV-2")
	commentadd("DEV-2", "只读", true)
	commentadd("DEV-2", "读写", false)
}
