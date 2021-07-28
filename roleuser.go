// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type RoleUserService struct {
	client *Client
}

type RoleUserGetOption struct {
	ID             string `url:"id"`
	ProjectIdOrKey string `url:"projectIdOrKey"`
}

type RoleUserGetObject struct {
	Self             string    `json:"self"`
	Name             string    `json:"name"`
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Actors           []Actor   `json:"actors"`
	Scope            RoleScope `json:"scope,omitempty"`
	TranslatedName   string    `json:"translatedName,omitempty"`
	CurrentUserRole  bool      `json:"currentUserRole,omitempty"`
	Admin            bool      `json:"admin,omitempty"`
	RoleConfigurable bool      `json:"roleConfigurable,omitempty"`
	Default          bool      `json:"default,omitempty"`
}

// Get 项目级
func (u *RoleUserService) Get(opts *RoleUserGetOption) (v *RoleUserGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/role/%v", u.client.endpoint, opts.ProjectIdOrKey, opts.ID)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// DefaultGet
func (u *RoleUserService) DefaultGet(opts *RoleUserGetOption) (v *RoleUserGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/role/%v/actors", u.client.endpoint, opts.ID)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RoleUserPostOption struct {
	ID             string   `url:"id" json:"-"`
	User           []string `json:"user"`
	ProjectIdOrKey string   `url:"projectIdOrKey" json:"-"`
}

type RoleUserPostObject RoleUserGetObject

// Post
func (u *RoleUserService) Post(opts *RoleUserPostOption) (v *RoleUserPostObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/role/%v", u.client.endpoint, opts.ProjectIdOrKey, opts.ID)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserPostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// DefaultPost
func (u *RoleUserService) DefaultPost(opts *RoleUserPostOption) (v *RoleUserPostObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/role/%v/actors", u.client.endpoint, opts.ID)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserPostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RoleUserDelOption struct {
	ID   string `url:"id" json:"-"`
	User string `url:"user"`
	// Group string `url:"group"`
	ProjectIdOrKey string `url:"projectIdOrKey" json:"-"`
}

type RoleUserDelObject RoleUserGetObject

// Del
func (u *RoleUserService) Del(opts *RoleUserDelOption) (v *RoleUserDelObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v/role/%v", u.client.endpoint, opts.ProjectIdOrKey, opts.ID)
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("DELETE", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserDelObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// DefaultDel
func (u *RoleUserService) DefaultDel(opts *RoleUserDelOption) (v *RoleUserDelObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/role/%v/actors", u.client.endpoint, opts.ID)
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("DELETE", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(RoleUserDelObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
