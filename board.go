package trello

import (
	"time"
)

const boardurl = "/boards/"

type boardJson struct {
	Closed           bool
	DateLastActivity *time.Time
	DateLastView     *time.Time
	Desc             string
	//	DescData string
	Id             string
	IdOrganization string
	//	Invitations []string
	Invited bool
	//	LabelNames []LabelName
	//	Memberships []MembershiP
	Name string
	//	Pinned
	//	PowerUps
	//	Prefs
	ShortLink string
	ShortUrl  string
	//	Starred
	//	Subcribed
	Url string
}

// Trello Board.
type Board struct {
	id   string
	c    *Client
	json *boardJson
}

func (b *Board) Name() string {
	return b.json.Name
}
