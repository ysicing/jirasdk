// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RoleService struct {
	client *Client
}

type RoleListOption struct{}
type RoleListObject []RoleObject

type RoleScope struct {
	Type    string `json:"type"`
	Project struct {
		ID   string `json:"id"`
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"project"`
}

type RoleObject struct {
	Self        string    `json:"self" structs:"self"`
	Name        string    `json:"name" structs:"name"`
	ID          int       `json:"id" structs:"id"`
	Description string    `json:"description" structs:"description"`
	Actors      []Actor   `json:"actors" structs:"actors"`
	Scope       RoleScope `json:"scope,omitempty"`
}

type Actor struct {
	ID          int        `json:"id"`
	DisplayName string     `json:"displayName"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	AvatarURL   string     `json:"avatarUrl"`
	ActorUser   ActorUser  `json:"actorUser,omitempty"`
	ActorGroup  ActorGroup `json:"actorGroup,omitempty"`
}

type ActorGroup struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type ActorUser struct {
	AccountID string `json:"accountId"`
}

func (u *RoleService) List(opts *RoleListOption) (v *RoleListObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/role"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleListObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RolePostOption struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RolePostObject RoleUserGetObject

func (u *RoleService) Post(opts *RolePostOption) (v *RolePostObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/role"
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RolePostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RoleGetOption struct {
	ProjectIdOrKey string `url:"projectIdOrKey"`
	ID             string `url:"id"`
}

type RoleGetObject RoleUserGetObject

func (u *RoleService) PGet(opts *RoleGetOption) (v *RoleGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/role/%v", u.client.endpoint, opts.ProjectIdOrKey, opts.ID)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RolePGetsObject struct{}

func (u *RoleService) PGets(opts *RoleGetOption) (v *interface{}, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/role", u.client.endpoint, opts.ProjectIdOrKey)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(interface{})
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
