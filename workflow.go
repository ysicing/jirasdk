// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type WorkflowService struct {
	client *Client
}

type WorkflowGetOption struct {
	WorkflowName string `url:"workflowName"`
}

type WorkflowGetObject []WorkflowObject
type WorkflowObject struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Steps            int    `json:"steps"`
	Default          bool   `json:"default"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	LastModifiedUser string `json:"lastModifiedUser,omitempty"`
}

func (u *WorkflowService) Get(opts *WorkflowGetOption) (v *WorkflowGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/workflow"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(WorkflowGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
