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

func userget(user string) {
	ug := jirasdk.UserGetOption{
		UserName: user,
	}
	v, resp, err := jiraapi.User.Get(&ug)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func usersearch(user string) {
	us := jirasdk.UserSearchOption{
		UserName: user,
		// MaxResults:      100,
		IncludeActive:   true,
		IncludeInactive: false,
		// StartAt:         0,
	}
	v, resp, err := jiraapi.User.Search(&us)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func userassignable(user, project string) {
	ua := jirasdk.UserAssignableOption{
		UserName:   user,
		Project:    project,
		MaxResults: 1,
	}
	v, resp, err := jiraapi.User.Assignable(&ua)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	userget("Jarvisbot")
	usersearch("chix")
	userassignable("chix", "DEV")
}
