// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package jirasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type IssueService struct {
	client *Client
}

type WorklogObject struct {
	Add struct {
		TimeSpent string `json:"timeSpent"`
		Started   string `json:"started"`
	} `json:"add"`
}

type FID struct {
	ID string `json:"id"`
}

type FName struct {
	Name string `json:"name"`
}

type FTime struct {
	OriginalEstimate  string `json:"originalEstimate"`
	RemainingEstimate string `json:"remainingEstimate"`
}

type Fields struct {
	Project          FID      `json:"project"`
	Summary          string   `json:"summary"`
	Issuetype        FID      `json:"issuetype"`
	Assignee         FName    `json:"assignee,omitempty"`
	Reporter         FName    `json:"reporter"`
	Priority         FID      `json:"priority,omitempty"`
	Labels           []string `json:"labels,omitempty"`
	Timetracking     FTime    `json:"timetracking,omitempty"`
	Versions         []FID    `json:"versions,omitempty"`
	Environment      string   `json:"environment,omitempty"`
	Description      string   `json:"description,omitempty"`
	Duedate          string   `json:"duedate,omitempty"`
	FixVersions      []FID    `json:"fixVersions,omitempty"`
	Components       []FID    `json:"component,omitempty"`
	Customfield30000 []string `json:"customfield_30000,omitempty"`
	Customfield20000 string   `json:"customfield_20000,omitempty"`
	Customfield40000 string   `json:"customfield_40000,omitempty"`
	Customfield70000 []string `json:"customfield_70000,omitempty"`
	Customfield60000 string   `json:"customfield_60000,omitempty"`
	Customfield50000 string   `json:"customfield_50000,omitempty"`
	Customfield10000 string   `json:"customfield_10000,omitempty"`
}

type IssuePostOption struct {
	Fields Fields `json:"fields"`
}

type IssuePostObject struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

func (u *IssueService) Post(opts *IssuePostOption) (v *IssuePostObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/issue?updateHistory=true"
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssuePostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueGetOption struct {
	IssueIdOrKey  string `url:"issueIdOrKey"`
	UpdateHistory bool   `url:"updateHistory,omitempty"`
}

type IssueGetObject struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Issuetype struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarID    int    `json:"avatarId"`
		} `json:"issuetype"`
		Components           []interface{} `json:"component"`
		Timespent            interface{}   `json:"timespent"`
		Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
		Description          interface{}   `json:"description"`
		Project              struct {
			Self           string `json:"self"`
			ID             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
			AvatarUrls     struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
		} `json:"project"`
		FixVersions        []interface{} `json:"fixVersions"`
		Aggregatetimespent interface{}   `json:"aggregatetimespent"`
		Resolution         interface{}   `json:"resolution"`
		Timetracking       struct {
		} `json:"timetracking"`
		Customfield10105      string        `json:"customfield_10105"`
		Attachment            []interface{} `json:"attachment"`
		Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
		Resolutiondate        interface{}   `json:"resolutiondate"`
		Workratio             int           `json:"workratio"`
		Summary               string        `json:"summary"`
		LastViewed            string        `json:"lastViewed"`
		Watches               struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
			IsWatching bool   `json:"isWatching"`
		} `json:"watches"`
		Creator struct {
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
		} `json:"creator"`
		Subtasks []interface{} `json:"subtasks"`
		Created  string        `json:"created"`
		Reporter struct {
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
		} `json:"reporter"`
		Customfield10000  string `json:"customfield_10000"`
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Priority struct {
			Self    string `json:"self"`
			IconURL string `json:"iconUrl"`
			Name    string `json:"name"`
			ID      string `json:"id"`
		} `json:"priority"`
		Customfield10100              interface{}   `json:"customfield_10100"`
		Customfield10101              interface{}   `json:"customfield_10101"`
		Labels                        []string      `json:"labels"`
		Environment                   interface{}   `json:"environment"`
		Timeestimate                  interface{}   `json:"timeestimate"`
		Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
		Versions                      []interface{} `json:"versions"`
		Duedate                       interface{}   `json:"duedate"`
		Progress                      struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"progress"`
		Comment struct {
			Comments   []interface{} `json:"comments"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			StartAt    int           `json:"startAt"`
		} `json:"comment"`
		Issuelinks []interface{} `json:"issuelinks"`
		Votes      struct {
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
			HasVoted bool   `json:"hasVoted"`
		} `json:"votes"`
		Worklog struct {
			StartAt    int           `json:"startAt"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Assignee struct {
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
		} `json:"assignee"`
		Updated string `json:"updated"`
		Status  struct {
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
		} `json:"status"`
	} `json:"fields"`
}

func (u *IssueService) Get(opts *IssueGetOption) (v *IssueGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v?updateHistory=true", u.client.endpoint, opts.IssueIdOrKey)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueAssigneePutOption struct {
	Name         string `json:"name"`
	IssueIdOrKey string `url:"issueIdOrKey" json:"-"`
}

type IssueAssigneePutObject struct{}

func (u *IssueService) Assignee(opts *IssueAssigneePutOption) (v *IssueAssigneePutObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/%v/assignee", u.client.endpoint, opts.IssueIdOrKey)
	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("PUT", path, bytes.NewBuffer(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueAssigneePutObject)
	resp, err = u.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssueMetaOption struct {
	ProjectIds     string `url:"projectIds"`
	ProjectKeys    string `url:"projectKeys"`
	IssuetypeIds   string `url:"issuetypeIds"`
	IssuetypeNames string `url:"issuetypeNames"`
}

type IssueMetaObject struct {
	Expand   string          `json:"expand"`
	Projects []ProjectObject `json:"projects"`
}

func (u *IssueService) Meta(opts *IssueMetaOption) (v *IssueMetaObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/issue/createmeta", u.client.endpoint)
	optv, _ := query.Values(opts)

	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(IssueMetaObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
