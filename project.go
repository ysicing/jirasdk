// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"github.com/google/go-querystring/query"
	"net/http"
	"strings"
)

type ProjectService struct {
	client *Client
}

type ProjectGetOption struct {
	IncludeArchived bool `url:"includeArchived,omitempty"`
}

type ProjectGetObject []ProjectObject

type ProjectObject struct {
	Expand     string `json:"expand,omitempty"`
	Self       string `json:"self"`
	ID         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectTypeKey string `json:"projectTypeKey,omitempty"`
	Issuetypes []IssueTypeObject `json:"issuetypes,omitempty"`
}

func (u *ProjectService) Get(opts *ProjectGetOption) (v *ProjectGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/project"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path + "?" + optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectSearchOption struct {
	IncludeArchived bool `url:"includeArchived,omitempty"`
	Search string `url:"search,omitempty"`
	MaxResults int `url:"maxResults,omitempty"`
}

func (u *ProjectService) Search(opts *ProjectSearchOption) (v *ProjectGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/project"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path + "?" + optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	var vv ProjectGetObject
	for i, j := range *v {
		if strings.Contains(j.Name, opts.Search) || strings.Contains(j.Key, opts.Search) {
			vv = append(vv,  j)
		}
		if i == opts.MaxResults {
			break
		}
	}
	return &vv, resp, err
}