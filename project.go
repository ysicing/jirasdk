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

type ProjectService struct {
	client *Client
}

type ProjectGetOption struct {
	IncludeArchived bool   `url:"includeArchived,omitempty"`
	Expand          string `json:"expand,omitempty"`
}

func (p *ProjectGetOption) Check() {
	if len(p.Expand) == 0 {
		p.Expand = "description,lead,url,projectKeys"
	}
}

type ProjectGetObject []ProjectObject

type ProjectObject struct {
	Expand      string `json:"expand,omitempty"`
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	AvatarUrls  struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory   `json:"projectCategory,omitempty"`
	ProjectTypeKey  string            `json:"projectTypeKey,omitempty"`
	Projectkeys     []string          `json:"projectKeys,omitempty"`
	Issuetypes      []IssueTypeObject `json:"issuetypes,omitempty"`
	Lead            Lead              `json:"lead,omitempty"`
	URL             string            `json:"url,omitempty"`
	Components      []ComponentObject `json:"components,omitempty"`
	IssueTypes      []IssueTypeObject `json:"issueTypes,omitempty"`
	AssigneeType    string            `json:"assigneeType"`
	Versions        []interface{}     `json:"versions,omitempty"`
	Archived        bool              `json:"archived,omitempty"`
	Roles           Roles             `json:"roles,omitempty"`
}

type ProjectCategory struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Roles struct {
	Administrators string `json:"Administrators"`
}

type Lead struct {
	Self       string `json:"self"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	Avatarurls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	Displayname string `json:"displayName"`
	Active      bool   `json:"active"`
}

func (u *ProjectService) Get(opts *ProjectGetOption) (v *ProjectGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/project"
	opts.Check()
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectSearchOption struct {
	IncludeArchived bool   `url:"includeArchived,omitempty"`
	Search          string `url:"search,omitempty"`
	MaxResults      int    `url:"maxResults,omitempty"`
	Expand          string `json:"expand,omitempty"`
}

type ProjectSearchObject ProjectObject

func (p *ProjectSearchOption) Check() {
	if len(p.Expand) == 0 {
		p.Expand = "description,lead,url,projectKeys"
	}
}

func (u *ProjectService) Search(opts *ProjectSearchOption) (v *ProjectSearchObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/api/2/project/%v", u.client.endpoint, opts.Search)
	opts.Check()
	optv, _ := query.Values(opts)
	req, err := http.NewRequest("GET", path+"?"+optv.Encode(), nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectSearchObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectTypeGetOption struct{}

type ProjectTypeGetObject []ProjectTypeObject

type ProjectTypeObject struct {
	Key                string `json:"key"`
	Formattedkey       string `json:"formattedKey"`
	Descriptioni18Nkey string `json:"descriptionI18nKey"`
	Icon               string `json:"icon"`
	Color              string `json:"color"`
}

func (u *ProjectService) Types(opts *ProjectTypeGetOption) (v *ProjectTypeGetObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/project/type"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectTypeGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectPostOption struct {
	Key                 string `json:"key"`
	Name                string `json:"name"`
	Projecttypekey      string `json:"projectTypeKey,omitempty"`
	Projecttemplatekey  string `json:"projectTemplateKey,omitempty"`
	Description         string `json:"description,omitempty"`
	Lead                string `json:"lead,omitempty"`
	URL                 string `json:"url,omitempty"`
	Assigneetype        string `json:"assigneeType,omitempty"`
	Avatarid            int    `json:"avatarId,omitempty"`
	Issuesecurityscheme int    `json:"issueSecurityScheme,omitempty"`
	Permissionscheme    int    `json:"permissionScheme,omitempty"`
	Notificationscheme  int    `json:"notificationScheme,omitempty"`
	Categoryid          int    `json:"categoryId,omitempty"`
}

func (p *ProjectPostOption) Check(user string) {
	if len(p.Projecttypekey) == 0 {
		p.Projecttypekey = "business"
		p.Projecttemplatekey = "com.atlassian.jira-core-project-templates:jira-core-project-management"
	}
	if len(p.Lead) == 0 {
		p.Lead = user
	}
}

type ProjectPostObject struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

func (u *ProjectService) Post(opts *ProjectPostOption) (v *ProjectPostObject, resp *http.Response, err error) {
	path := u.client.endpoint + "/rest/api/2/project"
	opts.Check(u.client.username)

	optv, _ := json.Marshal(opts)
	req, err := http.NewRequest("POST", path, bytes.NewReader(optv))
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectPostObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectIssueTypeGetOption struct {
	Project   string `json:"project"`
	Issuetype string `json:"issuetype"`
}

type ProjectIssueTypeFieldsObject struct {
	ViewScreen ViewScreen `json:"viewScreen"`
}

type SharedProjectObject struct {
	Expand      string `json:"expand,omitempty"`
	Self        string `json:"self,omitempty"`
	ID          int64  `json:"id,omitempty"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	AvatarUrls  struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory   `json:"projectCategory,omitempty"`
	ProjectTypeKey  string            `json:"projectTypeKey,omitempty"`
	Projectkeys     []string          `json:"projectKeys,omitempty"`
	Issuetypes      []IssueTypeObject `json:"issuetypes,omitempty"`
	Lead            Lead              `json:"lead,omitempty"`
	URL             string            `json:"url,omitempty"`
	Components      []ComponentObject `json:"components,omitempty"`
	IssueTypes      []IssueTypeObject `json:"issueTypes,omitempty"`
	AssigneeType    string            `json:"assigneeType"`
	Versions        []interface{}     `json:"versions,omitempty"`
	Archived        bool              `json:"archived,omitempty"`
	Roles           Roles             `json:"roles,omitempty"`
}

type ViewScreen struct {
	SharedWithProjects   []SharedProjectObject `json:"sharedWithProjects"`
	SharedWithIssueTypes []string              `json:"sharedWithIssueTypes"`
	TotalProjectsCount   int                   `json:"totalProjectsCount"`
	ScreenID             int                   `json:"screenId"`
	ScreenName           string                `json:"screenName"`
	CanEdit              bool                  `json:"canEdit"`
	State                string                `json:"state"`
	HiddenProjectsCount  int                   `json:"hiddenProjectsCount"`
}

func (u *ProjectService) IssueTypeFields(opts *ProjectIssueTypeGetOption) (v *ProjectIssueTypeFieldsObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/projectconfig/1/issuetype/%v/%v/fields", u.client.endpoint, opts.Project, opts.Issuetype)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectIssueTypeFieldsObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectIssueTypeGetObject struct {
	SharedWithProjects   []SharedProjectObject `json:"sharedWithProjects"`
	SharedWithIssueTypes []string              `json:"sharedWithIssueTypes"`
	TotalProjectsCount   int                   `json:"totalProjectsCount"`
	Name                 string                `json:"name"`
	State                string                `json:"state"`
	DisplayName          string                `json:"displayName"`
	IsDraftWithChanges   bool                  `json:"isDraftWithChanges"`
	DraftWithChanges     bool                  `json:"draftWithChanges"`
	HiddenProjectsCount  int                   `json:"hiddenProjectsCount"`
}

func (u *ProjectService) IssueTypeWorkflow(opts *ProjectIssueTypeGetOption) (v *ProjectIssueTypeGetObject, resp *http.Response, err error) {
	path := fmt.Sprintf("%v/rest/projectconfig/1/issuetype/%v/%v/workflow", u.client.endpoint, opts.Project, opts.Issuetype)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	u.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectIssueTypeGetObject)
	resp, err = u.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
