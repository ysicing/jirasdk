// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ComponentService struct {
	client *Client
}

type ComponentGetOption struct {
	ProjectIdOrKey string `url:"projectIdOrKey"`
}

type ComponentGetObject []ComponentObject

type ComponentObject struct {
	Self                string `json:"self"`
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Lead                Lead   `json:"lead,omitempty"`
	Description         string `json:"description"`
	Assigneetype        string `json:"assigneeType"`
	Assignee            Lead   `json:"assignee,omitempty"`
	Realassigneetype    string `json:"realAssigneeType"`
	Realassignee        Lead   `json:"realassignee,omitempty"`
	Isassigneetypevalid bool   `json:"isAssigneeTypeValid"`
	Project             string `json:"project"`
	Projectid           int    `json:"projectId"`
	Archived            bool   `json:"archived"`
}

func (u *ComponentService) Get(opts *ComponentGetOption) (v *ComponentGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/components", u.client.endpoint, opts.ProjectIdOrKey)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ComponentGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentPostOption struct {
	Project             string `json:"project"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Leadusername        string `json:"leadUserName,omitempty"`
	Assigneetype        string `json:"assigneeType,omitempty"`
	Isassigneetypevalid bool   `json:"isAssigneeTypeValid,omitempty"`
}

func (c *ComponentPostOption) Check(user string) {
	if len(c.Assigneetype) == 0 {
		c.Assigneetype = "COMPONENT_LEAD"
	}
	if len(c.Leadusername) == 0 {
		c.Leadusername = user
	}
}

type ComponentPostObject ComponentObject

func (u *ComponentService) Post(opts *ComponentPostOption) (v *ComponentPostObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/component", u.client.endpoint)
	opts.Check(u.client.username)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ComponentPostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
