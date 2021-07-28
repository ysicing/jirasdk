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

func get(key string) {
	wo := jirasdk.WatcherOption{
		IssueIdOrKey: key,
	}
	v, resp, err := jiraapi.Watcher.Get(&wo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func post(key, user string) {
	wo := jirasdk.WatcherOption{
		IssueIdOrKey: key,
		Username:     user,
	}
	v, resp, err := jiraapi.Watcher.Post(&wo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func del(key, user string) {
	wo := jirasdk.WatcherOption{
		IssueIdOrKey: key,
		Username:     user,
	}
	v, resp, err := jiraapi.Watcher.Del(&wo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	post("DEV-9", "wangshuo01")
	get("DEV-9")
	del("DEV-9", "wangshuo01")
	get("DEV-9")
}
