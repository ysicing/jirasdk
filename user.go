// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type UserService struct {
	client *Client
}

type UserGetOption struct {
	UserName string `url:"username"`
	Key      string `url:"key,omitempty"`
}

type UserGetObject struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls   struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`
	Groups      struct {
		Size  int           `json:"size"`
		Items []interface{} `json:"items"`
	} `json:"groups"`
	ApplicationRoles struct {
		Size  int           `json:"size"`
		Items []interface{} `json:"items"`
	} `json:"applicationRoles"`
	Expand string `json:"expand"`
}

func (u *UserService) Get(opts *UserGetOption) (v *UserGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/user"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(UserGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UserSearchOption struct {
	UserName        string `url:"username"`
	MaxResults      int    `url:"maxResults,omitempty"`
	IncludeActive   bool   `url:"includeActive,omitempty"`
	IncludeInactive bool   `url:"include_inactive,omitempty"`
	StartAt         int    `url:"startAt,omitempty"`
}

type UserObject struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls   struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`
}

type UserSearchObject []UserObject

func (u *UserService) Search(opts *UserSearchOption) (v *UserSearchObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/user/search"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(UserSearchObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UserAssignableOption struct {
	UserName           string `url:"username"`
	Project            string `url:"project"`
	MaxResults         int    `url:"maxResults,omitempty"`
	IssueKey           string `url:"issueKey,omitempty"`
	ActionDescriptorId int    `url:"actionDescriptorId,omitempty"`
	StartAt            int    `url:"startAt,omitempty"`
}

type UserAssignableObject []UserObject

func (u *UserService) Assignable(opts *UserAssignableOption) (v *UserAssignableObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/user/assignable/search"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(UserAssignableObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
