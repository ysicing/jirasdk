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

func defaultget(id string) {
	rugo := jirasdk.RoleUserGetOption{
		ID: id,
	}
	v, resp, err := jiraapi.RoleUser.DefaultGet(&rugo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func defaultpost(id, user string) {
	rupo := jirasdk.RoleUserPostOption{
		ID:   id,
		User: []string{user},
	}
	v, resp, err := jiraapi.RoleUser.DefaultPost(&rupo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func defaultdel(id, user string) {
	rudo := jirasdk.RoleUserDelOption{
		ID:   id,
		User: user,
	}
	v, resp, err := jiraapi.RoleUser.DefaultDel(&rudo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func get(key, id string) {
	rugo := jirasdk.RoleUserGetOption{
		ProjectIdOrKey: key,
		ID:             id,
	}
	v, resp, err := jiraapi.RoleUser.Get(&rugo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func post(key, id, user string) {
	rupo := jirasdk.RoleUserPostOption{
		ID:             id,
		User:           []string{user},
		ProjectIdOrKey: key,
	}
	v, resp, err := jiraapi.RoleUser.Post(&rupo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func del(key, id, user string) {
	rudo := jirasdk.RoleUserDelOption{
		ID:             id,
		User:           user,
		ProjectIdOrKey: key,
	}
	v, resp, err := jiraapi.RoleUser.DefaultDel(&rudo)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	defaultget("10100")
	defaultpost("10100", "w01")
	defaultpost("10100", "l06")
	defaultdel("10100", "z")
	get("DEV", "10100")
	post("DEV", "10100", "w01")
	post("DEV", "10100", "l06")
	del("DEV", "10100", "z")
}
