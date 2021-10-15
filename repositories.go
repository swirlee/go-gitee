package gitee

import (
	"fmt"
	"net/http"
	"time"
)

type RepositoriesService struct {
	client *Client
}

type Compare struct {
	BaseCommit      *Commit
	MergeBaseCommit *Commit
	Commits         []*Commit
	Files           []*struct {
		SHA        string `json:"sha,omitempty"`
		FileName   string `json:"filename,omitempty"`
		Status     string `json:"status,omitempty"`
		Additions  int    `json:"additions,omitempty"`
		Deletions  int    `json:"deletions,omitempty"`
		Changes    int    `json:"changes,omitempty"`
		BlobURL    string `json:"blob_url,omitempty"`
		RawURL     string `json:"raw_url,omitempty"`
		ContentURL string `json:"content_url,omitempty"`
		Patch      string `json:"patch,omitempty"`
	}
}

func (s *RepositoriesService) Compare(owner, repo, base, head string) (*Compare, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/compare/%s...%s", pathEscape(owner), pathEscape(repo), pathEscape(base), pathEscape(head))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := new(Compare)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

type Namespace struct {
	ID      *int    `json:"id"`
	Type    *string `json:"type"`
	Name    *string `json:"name"`
	Path    *string `json:"path"`
	HTMLURL *string `json:"html_url"`
}

type Owner struct {
	ID                *int    `json:"id"`
	Login             *string `json:"login"`
	Name              *string `json:"name"`
	AvatarURL         *string `json:"avatar_url"`
	URL               *string `json:"url"`
	HTMLURL           *string `json:"html_url"`
	Remark            *string `json:"remark"`
	FollowersURL      *string `json:"followers_url"`
	FollowingURL      *string `json:"following_url"`
	GistsURL          *string `json:"gists_url"`
	StarredURL        *string `json:"starred_url"`
	SubscriptionsURL  *string `json:"subscriptions_url"`
	OrganizationsURL  *string `json:"organizations_url"`
	ReposURL          *string `json:"repos_url"`
	EventsURL         *string `json:"events_url"`
	ReceivedEventsURL *string `json:"received_events_url"`
	Type              *string `json:"type"`
}

type Assigner struct {
	ID                *int    `json:"id"`
	Login             *string `json:"login"`
	Name              *string `json:"name"`
	AvatarURL         *string `json:"avatar_url"`
	URL               *string `json:"url"`
	HTMLURL           *string `json:"html_url"`
	Remark            *string `json:"remark"`
	FollowersURL      *string `json:"followers_url"`
	FollowingURL      *string `json:"following_url"`
	GistsURL          *string `json:"gists_url"`
	StarredURL        *string `json:"starred_url"`
	SubscriptionsURL  *string `json:"subscriptions_url"`
	OrganizationsURL  *string `json:"organizations_url"`
	ReposURL          *string `json:"repos_url"`
	EventsURL         *string `json:"events_url"`
	ReceivedEventsURL *string `json:"received_events_url"`
	Type              *string `json:"type"`
}

type Permission struct {
	Pull  *bool `json:"pull"`
	Push  *bool `json:"push"`
	Admin *bool `json:"admin"`
}

type Enterprise struct {
	ID      *int    `json:"id"`
	Type    *string `json:"type"`
	Name    *string `json:"name"`
	Path    *string `json:"path"`
	HTMLURL *string `json:"html_url"`
}

type Repository struct {
	ID                  *int        `json:"id"`
	FullName            *string     `json:"full_name"`
	HumanName           *string     `json:"human_name"`
	URL                 *string     `json:"url"`
	Namespace           *Namespace  `json:"namespace"`
	Path                *string     `json:"path"`
	Name                *string     `json:"name"`
	Owner               *Owner      `json:"owner"`
	Assigner            *Assigner   `json:"assigner"`
	Description         *string     `json:"description"`
	Private             *bool       `json:"private"`
	Public              *bool       `json:"public"`
	Internal            *bool       `json:"internal"`
	Fork                *bool       `json:"fork"`
	HTMLURL             *string     `json:"html_url"`
	SSHURL              *string     `json:"ssh_url"`
	ForksURL            *string     `json:"forks_url"`
	KeysURL             *string     `json:"keys_url"`
	CollaboratorsURL    *string     `json:"collaborators_url"`
	HooksURL            *string     `json:"hooks_url"`
	BranchesURL         *string     `json:"branches_url"`
	TagsURL             *string     `json:"tags_url"`
	BlobsURL            *string     `json:"blobs_url"`
	StargazersURL       *string     `json:"stargazers_url"`
	ContributorsURL     *string     `json:"contributors_url"`
	CommitsURL          *string     `json:"commits_url"`
	CommentsURL         *string     `json:"comments_url"`
	IssueCommentURL     *string     `json:"issue_comment_url"`
	IssuesURL           *string     `json:"issues_url"`
	PullsURL            *string     `json:"pulls_url"`
	MilestonesURL       *string     `json:"milestones_url"`
	NotificationsURL    *string     `json:"notifications_url"`
	LabelsURL           *string     `json:"labels_url"`
	ReleasesURL         *string     `json:"releases_url"`
	Recommend           *bool       `json:"recommend"`
	Gvp                 *bool       `json:"gvp"`
	Homepage            *string     `json:"homepage"`
	Language            *string     `json:"language"`
	ForksCount          *int        `json:"forks_count"`
	StargazersCount     *int        `json:"stargazers_count"`
	WatchersCount       *int        `json:"watchers_count"`
	DefaultBranch       *string     `json:"default_branch"`
	OpenIssuesCount     *int        `json:"open_issues_count"`
	HasIssues           *bool       `json:"has_issues"`
	HasWiki             *bool       `json:"has_wiki"`
	IssueComment        *bool       `json:"issue_comment"`
	CanComment          *bool       `json:"can_comment"`
	PullRequestsEnabled *bool       `json:"pull_requests_enabled"`
	HasPage             *bool       `json:"has_page"`
	License             *string     `json:"license"`
	Outsourced          *bool       `json:"outsourced"`
	ProjectCreator      *string     `json:"project_creator"`
	Members             []*string   `json:"members"`
	PushedAt            *time.Time  `json:"pushed_at"`
	CreatedAt           *time.Time  `json:"created_at"`
	UpdatedAt           *time.Time  `json:"updated_at"`
	Paas                *string     `json:"paas"`
	Stared              *bool       `json:"stared"`
	Watched             *bool       `json:"watched"`
	Permission          *Permission `json:"permission"`
	Relation            *string     `json:"relation"`
	AssigneesNumber     *int        `json:"assignees_number"`
	TestersNumber       *int        `json:"testers_number"`
	Assignee            []*string   `json:"assignee"`
	Testers             []*string   `json:"testers"`
	Status              *string     `json:"status"`
	EmptyRepo           *bool       `json:"empty_repo"`
	Programs            []*string   `json:"programs"`
	Enterprise          *Enterprise `json:"enterprise"`
	Parent              *Repository `json:"parent"`
}

type VisibilityValue string

const (
	PrivateVisibility VisibilityValue = "private"
	PublicVisibility  VisibilityValue = "public"
	AllVisibility     VisibilityValue = "all"
)

type RepoTypeValue string

const (
	OwnerType    RepoTypeValue = "owner"
	PersonalType RepoTypeValue = "personal"
	MemberType   RepoTypeValue = "member"
	PublicType   RepoTypeValue = "public"
	PrivateType  RepoTypeValue = "private"
)

type SortValue string

const (
	CreatedSortValue  SortValue = "created"
	UpdatedSortValue  SortValue = "updated"
	PushedSortValue   SortValue = "pushed"
	FullNameSortValue SortValue = "full_name"
)

type ListRepositoriesOptions struct {
	ListOptions
	Visibility  *VisibilityValue    `url:"visibility,omitempty" json:"visibility,omitempty"`
	Affiliation *string             `url:"affiliation,omitempty" json:"affiliation,omitempty"`
	Type        *RepoTypeValue      `url:"type,omitempty" json:"type,omitempty"`
	Sort        *SortValue          `url:"sort,omitempty" json:"sort,omitempty"`
	Direction   *SortDirectionValue `url:"direction,omitempty" json:"direction,omitempty"`
	Q           *string             `url:"q,omitempty" json:"q,omitempty"`
}

func (s *RepositoriesService) List(opts *ListRepositoriesOptions) ([]*Repository, *Response, error) {
	url := "user/repos"
	req, err := s.client.NewRequest(http.MethodGet, url, opts)
	if err != nil {
		return nil, nil, err
	}

	var p []*Repository
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}