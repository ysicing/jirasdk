// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransitionsService struct {
	client *Client
}

type IssueTransitionsGetOption struct {
	IssueIdOrKey string `url:"issueIdOrKey"`
	TransitionId string `url:"transitionId"`
}

type IssueTransitionsGetObject struct {
	Expand      string              `json:"expand"`
	Transitions []TransitionsObject `json:"transitions"`
}

type TransitionsObject struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	To   StatusObject `json:"to"`
}

// issue流转
func (u *TransitionsService) Get(opts *IssueTransitionsGetOption) (v *IssueTransitionsGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/transitions", u.client.endpoint, opts.IssueIdOrKey)
	if len(opts.TransitionId) > 0 {
		path = fmt.Sprintf("%v?transitionId=%v", path, opts.TransitionId)
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueTransitionsGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type TTransition struct {
	ID string `json:"id"`
}

type TFields struct {
	Assignee struct {
		Name string `json:"name"`
	} `json:"assignee"`
	Resolution struct {
		Name string `json:"name"`
	} `json:"resolution"`
}

type TUpdate struct {
	Comment []struct {
		Add struct {
			Body string `json:"body"`
		} `json:"add"`
	} `json:"comment"`
}

type IssueTransitionsPostOption struct {
	IssueIdOrKey string `url:"issueIdOrKey" json:"-"`
	//Update TUpdate `json:"update"`
	//Fields TFields`json:"fields"`
	Transition TTransition `json:"transition"`
}

type IssueTransitionsPostObject struct {
}

func (u *TransitionsService) Post(opts *IssueTransitionsPostOption) (v *IssueTransitionsPostObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/transitions", u.client.endpoint, opts.IssueIdOrKey)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueTransitionsPostObject)
	resp, err = u.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}
