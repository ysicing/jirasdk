// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/jirasdk"
	"log"
)

var jiraapi *jirasdk.Client

func init()  {
	c, err := jirasdk.NewClient("http://172.16.74.113:8080", "Jarvisbot", "12345678")
	if  err != nil {
		panic(err)
	}
	jiraapi = c
}

func get(name ...string) {
	it := jirasdk.IssueTransitionsGetOption{
		IssueIdOrKey: name[0],
	}
	if len(name) > 1 {
		it.TransitionId = name[1]
	}
	v, resp, err := jiraapi.Issue.TransitionsGet(&it)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
	for _,j := range v.Transitions {
		log.Printf("id: %v, name: %v",j.To.ID, j.To.Name)
	}
}

func post(name string, next string) {

	tt := jirasdk.TTransition{ID: next}

	it := jirasdk.IssueTransitionsPostOption{
		IssueIdOrKey: name,
		Transition: tt,
	}

	v, resp, err := jiraapi.Issue.TransitionsPost(&it)
	if err != nil {
		fmt.Println(resp.StatusCode, err)
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main()  {
	get("DEV-3")
	post("DEV-3", "111")
	get("DEV-3")
}
