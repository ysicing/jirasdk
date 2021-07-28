// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type IssueSearchService struct {
	client *Client
}

type IssueSearchOption struct {
	Jql           string `url:"jql"`
	StartAt       int    `url:"startAt"`
	MaxResults    int    `url:"maxResults"`
	ValidateQuery bool   `url:"validateQuery"`
	Fields        string `url:"fields"`
	Expand        string `url:"expand"`
}

func (i *IssueSearchOption) Check() {
	// if len(i.Expand) == 0 {
	// 	i.Expand = "schema,names"
	// }
	if i.MaxResults < 100 {
		i.MaxResults = 100
	}
}

type IssueSearchObject struct {
	Expand     string        `json:"expand"`
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Issues     []IssueObject `json:"issues"`
}

type IssueObject struct {
	Expand string      `json:"expand"`
	ID     string      `json:"id"`
	Self   string      `json:"self"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Issuetype             IssueTypeObject `json:"issuetype"`
	Components            []interface{}   `json:"components"`
	Timespent             interface{}     `json:"timespent"`
	Timeoriginalestimate  interface{}     `json:"timeoriginalestimate"`
	Description           interface{}     `json:"description"`
	Project               ProjectObject   `json:"project"`
	FixVersions           []interface{}   `json:"fixVersions"`
	Aggregatetimespent    interface{}     `json:"aggregatetimespent"`
	Resolution            interface{}     `json:"resolution"`
	Customfield10105      string          `json:"customfield_10105"`
	Aggregatetimeestimate interface{}     `json:"aggregatetimeestimate"`
	Resolutiondate        interface{}     `json:"resolutiondate"`
	Workratio             int             `json:"workratio"`
	Summary               string          `json:"summary"`
	LastViewed            string          `json:"lastViewed"`
	Watches               Watches         `json:"watches"`
	Creator               AuthorObject    `json:"creator"`
	Subtasks              []interface{}   `json:"subtasks"`
	Created               string          `json:"created"`
	Reporter              AuthorObject    `json:"reporter"`
	Customfield10000      string          `json:"customfield_10000"`
	Aggregateprogress     struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"aggregateprogress"`
	Priority                      PriorityObject `json:"priority"`
	Customfield10100              interface{}    `json:"customfield_10100"`
	Customfield10101              interface{}    `json:"customfield_10101"`
	Labels                        []string       `json:"labels"`
	Environment                   interface{}    `json:"environment"`
	Timeestimate                  interface{}    `json:"timeestimate"`
	Aggregatetimeoriginalestimate interface{}    `json:"aggregatetimeoriginalestimate"`
	Versions                      []interface{}  `json:"versions"`
	Duedate                       interface{}    `json:"duedate"`
	Progress                      struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"progress"`
	Issuelinks []interface{} `json:"issuelinks"`
	Votes      Votes         `json:"votes"`
	Assignee   AuthorObject  `json:"assignee"`
	Updated    string        `json:"updated"`
	Status     StatusObject  `json:"status"`
}

type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}

func (u *IssueSearchService) Get(opts *IssueSearchOption) (v *IssueSearchObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/search"
	opts.Check()
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueSearchObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
