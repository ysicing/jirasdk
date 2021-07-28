## JiraSDK

> Go client library for [Jira](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/)

> `go get -u github.com/ysicing/jirasdk`


##  Requirements

- Go >= 1.15
- Jira v8.2.2 (self-host)

## Support 

### 用户User

- Get/Search
- Assignable

### 状态Status

- List
- Get 
- ProjectGet 项目级Get

### 项目Project

- Get/Search
- Types
- Post

### 优先级Priority

- Get

### 问题类型IssueType

- List
- Get
- Post

### 问题issue

- Meta
- Get
- Post
- Assignee

## issue流转Transitions

- Get
- Post

### issue模块Component

- Get
- Post

### issue评论Comment

- Get
- Post

### issue流程管理方案Scheme

- List
- Get

## issue搜索IssueSearch

- Get

## issue关注watcher

- Get
- Post
- Del

## 角色role

- List
- Post
- PGet # 查看项目角色成员
- PGets # 查看项目角色列表

## 角色成员roleuser

- Get
- Post
- Del
- DefaultGet # 全局
- DefaultPost # 全局
- DefaultDel # 全局