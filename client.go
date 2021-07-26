// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	username, password, endpoint string
	httpClient *http.Client
	User *UserService
	Project *ProjectService
	Issue *IssueService
	IssueType *IssueTypeService
	Priority *PriorityService
	Status *StatusService
}

func NewClient(endpoint, username, password  string) (*Client, error)  {
	c := &Client{username: username, password: password, httpClient: http.DefaultClient}
	if endpoint == "" {
		endpoint = defaultBaseURL
	}
	c.endpoint = endpoint
	c.User = &UserService{client: c}
	c.Project = &ProjectService{client: c}
	c.Issue = &IssueService{client: c}
	c.IssueType = &IssueTypeService{client: c}
	c.Priority = &PriorityService{client: c}
	c.Status = &StatusService{client: c}
	return c,  nil
}

func (c *Client) requestExtHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("User-Agent", userAgent)
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	return Do(c.httpClient, req, v)
}

func Do(c *http.Client, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	if v != nil {
		defer resp.Body.Close()
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			decoder := json.NewDecoder(resp.Body)
			decoder.DisallowUnknownFields()
			err = decoder.Decode(v)
		}
	}
	return resp, err
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	ErrorMessages []string `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.ErrorMessages[0])
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	if r == nil {
		return errors.New("no resp returned")
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse)
	}
	if len(errorResponse.ErrorMessages) > 0 {
		return errors.New(errorResponse.ErrorMessages[0])
	}
	if len(errorResponse.Errors) > 0 {
		for key, value := range errorResponse.Errors {
			return fmt.Errorf("%s - %s", key, value)
		}
	}
	return nil
}