// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type WatcherService struct {
	client *Client
}

type WatcherOption struct {
	IssueIdOrKey string `url:"issueIdOrKey"`
	Username     string `url:"username"`
}

type WatcherGetObject struct {
	Self       string         `json:"self"`
	IsWatching bool           `json:"isWatching"`
	WatchCount int            `json:"watchCount"`
	Watchers   []AuthorObject `json:"watchers"`
}

func (u *WatcherService) Get(opts *WatcherOption) (v *WatcherGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/watchers", u.client.endpoint, opts.IssueIdOrKey)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(WatcherGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type WatcherObject struct{}

func (u *WatcherService) Del(opts *WatcherOption) (v *WatcherObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/watchers", u.client.endpoint, opts.IssueIdOrKey)
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("DELETE", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(WatcherObject)
	resp, err = u.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}

func (u *WatcherService) Post(opts *WatcherOption) (v *WatcherObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/watchers", u.client.endpoint, opts.IssueIdOrKey)
	req, err := http.NewRequest("POST", path, strings.NewReader("\""+opts.Username+"\""))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(WatcherObject)
	resp, err = u.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}
