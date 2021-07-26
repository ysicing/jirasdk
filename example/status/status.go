// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/jirasdk"
)

var jiraapi *jirasdk.Client

func init()  {
	c, err := jirasdk.NewClient("http://172.16.74.113:8080", "Jarvisbot", "12345678")
	if  err != nil {
		panic(err)
	}
	jiraapi = c
}

func statuslist()  {
	sg := jirasdk.StatusGetOptions{}
	v, resp, err := jiraapi.Status.List(&sg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func statusget(name string)  {
	sg := jirasdk.StatusGetOptions{
		IdOrName: name,
	}
	v, resp, err := jiraapi.Status.Get(&sg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func pstatusget(name string)  {
	sg := jirasdk.ProjectStatusGetOptions{
		ProjectIdOrKey: name,
	}
	v, resp, err := jiraapi.Status.ProjectStatusGet(&sg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main()  {
	statuslist()
	statusget("待办")
	statusget("10005")
	pstatusget("DEV")
}

