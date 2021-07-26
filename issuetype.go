// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"github.com/google/go-querystring/query"
	"net/http"
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

type IssueTypeGetOption struct{}

type IssueTypeGetObject []IssueTypeObject

func (u *IssueTypeService) Get(opts *IssueTypeGetOption) (v *IssueTypeGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/issuetype"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
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
