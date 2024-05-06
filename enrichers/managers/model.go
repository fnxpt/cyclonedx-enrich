package managers

import "time"

// COCOAPODS
type CocoaPodsPackage struct {
	Name           string     `json:"name"`
	Version        string     `json:"version"`
	License        *string    `json:"license"`
	Summary        string     `json:"summary"`
	Homepage       *string    `json:"homepage"`
	Authors        *Authors   `json:"authors"`
	SocialMediaURL *string    `json:"social_media_url"`
	Source         *Source    `json:"source"`
	Platforms      *Platforms `json:"platforms"`
	SourceFiles    *string    `json:"source_files"`
	SwiftVersions  *string    `json:"swift_versions"`
	SwiftVersion   *string    `json:"swift_version"`
}

type Authors struct {
	RobertPayne string `json:"Robert Payne"`
}

type Platforms struct {
	Ios  string `json:"ios"`
	Osx  string `json:"osx"`
	Tvos string `json:"tvos"`
}

type Source struct {
	Git string `json:"git"`
	Tag string `json:"tag"`
}

// NPM
type NpmPackage struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	// Author       NpmUser    `json:"author"`  #NpmUser can also be a string
	// Contributors []NpmUser   `json:"contributors"` #NpmUser can also be a string
	License     *string     `json:"license"`
	Repository  *Repository `json:"repository"`
	Keywords    []string    `json:"keywords"`
	GitHead     string      `json:"gitHead"`
	Homepage    *string     `json:"homepage"`
	Dist        Dist        `json:"dist"`
	Maintainers *[]NpmUser  `json:"maintainers"`
}

type NpmUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Dist struct {
	Integrity    string      `json:"integrity"`
	Shasum       string      `json:"shasum"`
	Tarball      string      `json:"tarball"`
	FileCount    int64       `json:"fileCount"`
	UnpackedSize int64       `json:"unpackedSize"`
	Signatures   []Signature `json:"signatures"`
}

type Signature struct {
	Keyid string `json:"keyid"`
	Sig   string `json:"sig"`
}

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// PyPi
type PyPiPackage struct {
	Info            Info          `json:"info"`
	LastSerial      int64         `json:"last_serial"`
	Urls            []URL         `json:"urls"`
	Vulnerabilities []interface{} `json:"vulnerabilities"`
}

type Info struct {
	Author                 string        `json:"author"`
	AuthorEmail            string        `json:"author_email"`
	BugtrackURL            interface{}   `json:"bugtrack_url"`
	Classifiers            []interface{} `json:"classifiers"`
	Description            string        `json:"description"`
	DescriptionContentType interface{}   `json:"description_content_type"`
	DocsURL                interface{}   `json:"docs_url"`
	DownloadURL            string        `json:"download_url"`
	Downloads              Downloads     `json:"downloads"`
	Dynamic                interface{}   `json:"dynamic"`
	HomePage               string        `json:"home_page"`
	Keywords               string        `json:"keywords"`
	License                *string       `json:"license"`
	Maintainer             string        `json:"maintainer"`
	MaintainerEmail        string        `json:"maintainer_email"`
	Name                   string        `json:"name"`
	PackageURL             string        `json:"package_url"`
	Platform               string        `json:"platform"`
	ProjectURL             string        `json:"project_url"`
	ProjectUrls            ProjectUrls   `json:"project_urls"`
	ProvidesExtra          interface{}   `json:"provides_extra"`
	ReleaseURL             string        `json:"release_url"`
	RequiresDist           []string      `json:"requires_dist"`
	RequiresPython         string        `json:"requires_python"`
	Summary                string        `json:"summary"`
	Version                string        `json:"version"`
	Yanked                 bool          `json:"yanked"`
	YankedReason           interface{}   `json:"yanked_reason"`
}

type Downloads struct {
	LastDay   int64 `json:"last_day"`
	LastMonth int64 `json:"last_month"`
	LastWeek  int64 `json:"last_week"`
}

type ProjectUrls struct {
	Homepage string `json:"Homepage"`
}

type URL struct {
	CommentText    string      `json:"comment_text"`
	Digests        Digests     `json:"digests"`
	Downloads      int64       `json:"downloads"`
	Filename       string      `json:"filename"`
	HasSig         bool        `json:"has_sig"`
	Md5Digest      string      `json:"md5_digest"`
	Packagetype    string      `json:"packagetype"`
	PythonVersion  string      `json:"python_version"`
	RequiresPython interface{} `json:"requires_python"`
	Size           int64       `json:"size"`
	// UploadTime        time.Time   `json:"upload_time"`
	// UploadTimeISO8601 time.Time   `json:"upload_time_iso_8601"`
	URL          string      `json:"url"`
	Yanked       bool        `json:"yanked"`
	YankedReason interface{} `json:"yanked_reason"`
}

type Digests struct {
	Blake2B256 string `json:"blake2b_256"`
	Md5        string `json:"md5"`
	Sha256     string `json:"sha256"`
}

// Github
type GithubPackage struct {
	ID                        int64               `json:"id"`
	NodeID                    string              `json:"node_id"`
	Name                      string              `json:"name"`
	FullName                  string              `json:"full_name"`
	Private                   bool                `json:"private"`
	Owner                     Owner               `json:"owner"`
	HTMLURL                   string              `json:"html_url"`
	Description               string              `json:"description"`
	Fork                      bool                `json:"fork"`
	URL                       string              `json:"url"`
	ForksURL                  string              `json:"forks_url"`
	KeysURL                   string              `json:"keys_url"`
	CollaboratorsURL          string              `json:"collaborators_url"`
	TeamsURL                  string              `json:"teams_url"`
	HooksURL                  string              `json:"hooks_url"`
	IssueEventsURL            string              `json:"issue_events_url"`
	EventsURL                 string              `json:"events_url"`
	AssigneesURL              string              `json:"assignees_url"`
	BranchesURL               string              `json:"branches_url"`
	TagsURL                   string              `json:"tags_url"`
	BlobsURL                  string              `json:"blobs_url"`
	GitTagsURL                string              `json:"git_tags_url"`
	GitRefsURL                string              `json:"git_refs_url"`
	TreesURL                  string              `json:"trees_url"`
	StatusesURL               string              `json:"statuses_url"`
	LanguagesURL              string              `json:"languages_url"`
	StargazersURL             string              `json:"stargazers_url"`
	ContributorsURL           string              `json:"contributors_url"`
	SubscribersURL            string              `json:"subscribers_url"`
	SubscriptionURL           string              `json:"subscription_url"`
	CommitsURL                string              `json:"commits_url"`
	GitCommitsURL             string              `json:"git_commits_url"`
	CommentsURL               string              `json:"comments_url"`
	IssueCommentURL           string              `json:"issue_comment_url"`
	ContentsURL               string              `json:"contents_url"`
	CompareURL                string              `json:"compare_url"`
	MergesURL                 string              `json:"merges_url"`
	ArchiveURL                string              `json:"archive_url"`
	DownloadsURL              string              `json:"downloads_url"`
	IssuesURL                 string              `json:"issues_url"`
	PullsURL                  string              `json:"pulls_url"`
	MilestonesURL             string              `json:"milestones_url"`
	NotificationsURL          string              `json:"notifications_url"`
	LabelsURL                 string              `json:"labels_url"`
	ReleasesURL               string              `json:"releases_url"`
	DeploymentsURL            string              `json:"deployments_url"`
	CreatedAt                 time.Time           `json:"created_at"`
	UpdatedAt                 time.Time           `json:"updated_at"`
	PushedAt                  time.Time           `json:"pushed_at"`
	GitURL                    string              `json:"git_url"`
	SSHURL                    string              `json:"ssh_url"`
	CloneURL                  string              `json:"clone_url"`
	SvnURL                    string              `json:"svn_url"`
	Homepage                  string              `json:"homepage"`
	Size                      int64               `json:"size"`
	StargazersCount           int64               `json:"stargazers_count"`
	WatchersCount             int64               `json:"watchers_count"`
	Language                  string              `json:"language"`
	HasIssues                 bool                `json:"has_issues"`
	HasProjects               bool                `json:"has_projects"`
	HasDownloads              bool                `json:"has_downloads"`
	HasWiki                   bool                `json:"has_wiki"`
	HasPages                  bool                `json:"has_pages"`
	HasDiscussions            bool                `json:"has_discussions"`
	ForksCount                int64               `json:"forks_count"`
	MirrorURL                 interface{}         `json:"mirror_url"`
	Archived                  bool                `json:"archived"`
	Disabled                  bool                `json:"disabled"`
	OpenIssuesCount           int64               `json:"open_issues_count"`
	License                   *License            `json:"license"`
	AllowForking              bool                `json:"allow_forking"`
	IsTemplate                bool                `json:"is_template"`
	WebCommitSignoffRequired  bool                `json:"web_commit_signoff_required"`
	Topics                    []interface{}       `json:"topics"`
	Visibility                string              `json:"visibility"`
	Forks                     int64               `json:"forks"`
	OpenIssues                int64               `json:"open_issues"`
	Watchers                  int64               `json:"watchers"`
	DefaultBranch             string              `json:"default_branch"`
	Permissions               Permissions         `json:"permissions"`
	TempCloneToken            string              `json:"temp_clone_token"`
	AllowSquashMerge          bool                `json:"allow_squash_merge"`
	AllowMergeCommit          bool                `json:"allow_merge_commit"`
	AllowRebaseMerge          bool                `json:"allow_rebase_merge"`
	AllowAutoMerge            bool                `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool                `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool                `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool                `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string              `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string              `json:"squash_merge_commit_title"`
	MergeCommitMessage        string              `json:"merge_commit_message"`
	MergeCommitTitle          string              `json:"merge_commit_title"`
	SecurityAndAnalysis       SecurityAndAnalysis `json:"security_and_analysis"`
	NetworkCount              int64               `json:"network_count"`
	SubscribersCount          int64               `json:"subscribers_count"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type Owner struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}

type SecurityAndAnalysis struct {
	SecretScanning               DependabotSecurityUpdates `json:"secret_scanning"`
	SecretScanningPushProtection DependabotSecurityUpdates `json:"secret_scanning_push_protection"`
	DependabotSecurityUpdates    DependabotSecurityUpdates `json:"dependabot_security_updates"`
	SecretScanningValidityChecks DependabotSecurityUpdates `json:"secret_scanning_validity_checks"`
}

type DependabotSecurityUpdates struct {
	Status string `json:"status"`
}
