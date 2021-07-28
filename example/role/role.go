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
	rlo := jirasdk.RoleListOption{}
	v, resp, err := jiraapi.Role.List(&rlo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func post(name, desc string) {
	rpo := jirasdk.RolePostOption{
		Name:        name,
		Description: desc,
	}
	v, resp, err := jiraapi.Role.Post(&rpo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func pget(key, id string) {
	rlo := jirasdk.RoleGetOption{
		ProjectIdOrKey: key,
		ID:             id,
	}
	v, resp, err := jiraapi.Role.PGet(&rlo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func pgets(key string) {
	rlo := jirasdk.RoleGetOption{
		ProjectIdOrKey: key,
	}
	v, resp, err := jiraapi.Role.PGets(&rlo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	list()
	post("demorolex", "demorole")
	list()
	pget("DEV", "10100")
	pgets("DEV")
}
