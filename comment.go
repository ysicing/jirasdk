// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommentService struct {
	client *Client
}

type IssueCommentGetOption struct {
	IssueIdOrKey string `url:"issueIdOrKey"`
}

type AuthorObject struct {
	Self         string `json:"self"`
	Name         string `json:"name"`
	Key          string `json:"key"`
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
}

type CommentBody struct {
	Self         string       `json:"self"`
	ID           string       `json:"id"`
	Author       AuthorObject `json:"author"`
	Body         string       `json:"body"`
	UpdateAuthor AuthorObject `json:"updateAuthor"`
	Created      string       `json:"created"`
	Updated      string       `json:"updated"`
	Visibility   struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility,omitempty"`
}

type IssueCommentGetObject struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Comments   []CommentBody `json:"comments"`
}

func (u *CommentService) Get(opts *IssueCommentGetOption) (v *IssueCommentGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/comment", u.client.endpoint, opts.IssueIdOrKey)
	// optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueCommentGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueCommentPostOption struct {
	IssueIdOrKey string          `url:"issueIdOrKey" json:"-"`
	Body         string          `json:"body"`
	Visibility   IssueVisibility `json:"visibility,omitempty"`
}

type IssueVisibility struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type IssueCommentPostObject struct {
	Self         string       `json:"self"`
	ID           string       `json:"id"`
	Author       AuthorObject `json:"author"`
	Body         string       `json:"body"`
	UpdateAuthor AuthorObject `json:"updateAuthor"`
	Created      string       `json:"created"`
	Updated      string       `json:"updated"`
	Visibility   struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
}

func (u *CommentService) Post(opts *IssueCommentPostOption) (v *IssueCommentPostObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/comment", u.client.endpoint, opts.IssueIdOrKey)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueCommentPostObject)
	resp, err = u.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}
