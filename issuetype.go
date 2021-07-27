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

type IssueTypeService struct {
	client *Client
}

type IssueTypeObject struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
	AvatarID    int    `json:"avatarId,omitempty"`
}

type IssueTypeListOption struct{}

type IssueTypeListObject []IssueTypeObject

func (u *IssueTypeService) List(opts *IssueTypeListOption) (v *IssueTypeListObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/issuetype"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueTypeListObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueTypeGetOption struct {
	ID string `json:"id"`
}

type IssueTypeGetObject IssueTypeObject

func (u *IssueTypeService) Get(opts *IssueTypeGetOption) (v *IssueTypeGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issuetype/%v", u.client.endpoint, opts.ID)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueTypeGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueTypeType string

const (
	// Standard 标准
	Standard IssueTypeType = "standard"
	// Subtask 子任务
	Subtask IssueTypeType = "subtask"
)

type IssueTypePostOption struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Type        IssueTypeType `json:"type,omitempty"`
}

func (it *IssueTypePostOption) Check() {
	if len(it.Type) == 0 {
		it.Type = Standard
	}
}

type IssueTypePostObject IssueTypeObject

func (u *IssueTypeService) Post(opts *IssueTypePostOption) (v *IssueTypePostObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/issuetype"
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueTypePostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
