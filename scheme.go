// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type SchemeService struct {
	client *Client
}

type SchemeListOptions struct{}

type SchemeListObject struct {
	Expand  string         `json:"expand"`
	Schemes []SchemeObject `json:"schemes"`
}

type SchemeObject struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (u *SchemeService) List(opts *SchemeListOptions) (v *SchemeListObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/issuetypescheme"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(SchemeListObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type SchemeGetOptions struct {
	SchemeId string `url:"schemeId"`
	Expand   string `url:"expand"`
}

func (s *SchemeGetOptions) Check() {
	if len(s.Expand) == 0 {
		s.Expand = "defaultIssueType,issueTypes"
	}
}

type SchemeGetObject struct {
	Expand           string            `json:"expand"`
	Self             string            `json:"self"`
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	DefaultIssueType IssueTypeObject   `json:"defaultIssueType,omitempty"`
	IssueTypes       []IssueTypeObject `json:"issueTypes,omitempty"`
}

func (u *SchemeService) Get(opts *SchemeGetOptions) (v *SchemeGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issuetypescheme/%v", u.client.endpoint, opts.SchemeId)
	opts.Check()
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(SchemeGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
