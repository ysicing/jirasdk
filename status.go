// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"fmt"
	"net/http"
)

type StatusService struct {
	client *Client
}

type StatusGetOptions struct {
	IdOrName string `url:"idOrName"`
}

type StatusGetObject []StatusObject

type StatusObject struct {
	Self           string `json:"self"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	ID             string `json:"id"`
	StatusCategory struct {
		Self      string `json:"self"`
		ID        int    `json:"id"`
		Key       string `json:"key"`
		ColorName string `json:"colorName"`
		Name      string `json:"name"`
	} `json:"statusCategory"`
}

func (u *StatusService) List(opts *StatusGetOptions) (v *StatusGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/status"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(StatusGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

func (u *StatusService) Get(opts *StatusGetOptions) (v *StatusObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/status/%v", u.client.endpoint, opts.IdOrName)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(StatusObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectStatusGetOptions struct {
	ProjectIdOrKey string `url:"projectIdOrKey"`
}

type ProjectStatusObject []PStatusObject

type PStatusObject struct {
	Self     string         `json:"self"`
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Subtask  bool           `json:"subtask"`
	Statuses []StatusObject `json:"statuses"`
}

func (u *StatusService) ProjectStatusGet(opts *ProjectStatusGetOptions) (v *ProjectStatusObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project_get/%v/statuses", u.client.endpoint, opts.ProjectIdOrKey)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectStatusObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
