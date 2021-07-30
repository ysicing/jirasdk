// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type CustomFieldService struct {
	client *Client
}

type FieldGetOption struct {
	ProjectId string `url:"projectId"`
}

type FieldGetObject struct {
	MaxResults int           `json:"maxResults"`
	StartAt    int           `json:"startAt"`
	Total      int           `json:"total"`
	IsLast     bool          `json:"isLast"`
	Values     []FieldObject `json:"values"`
}

type FieldObject struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Type          string `json:"type"`
	SearcherKey   string `json:"searcherKey"`
	Self          string `json:"self"`
	NumericID     int    `json:"numericId"`
	IsLocked      bool   `json:"isLocked"`
	IsManaged     bool   `json:"isManaged"`
	IsAllProjects bool   `json:"isAllProjects"`
	ProjectsCount int    `json:"projectsCount"`
	ScreensCount  int    `json:"screensCount"`
}

func (u *CustomFieldService) Get(opts *FieldGetOption) (v *FieldGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/customFields"
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(FieldGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
