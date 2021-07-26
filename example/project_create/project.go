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
	pp := jirasdk.ProjectPostOption{
		Key:  "DEV11",
		Name: "DEV11",
	}
	v, resp, err := jiraapi.Project.Post(&pp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	create()
}
