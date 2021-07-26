// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type PriorityService struct {
	client *Client
}

type PriorityObject struct {
	Self        string `json:"self"`
	StatusColor string `json:"statusColor"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type PriorityGetOption struct {}

type PriorityGetObject []PriorityObject

func (u *PriorityService) Get(opts *PriorityGetOption) (v *PriorityGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/priority"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path + "?" + optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(PriorityGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
