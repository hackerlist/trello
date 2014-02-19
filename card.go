package trello

import (
	//	"encoding/json"
	"time"
)

type cardJson struct {
	//	Badges
	//	CheckItemStates
	Closed           bool
	DateLastActivity *time.Time
	Desc             string
	//	DescData
	Due *time.Time
	Id  string
	//	IdAttachmentCover
	IdBoard string
	//	IdChecklists
	IdList                string
	IdMembers             []string
	IdMembersVoted        []string
	IdShort               float64
	Labels                []string
	ManualCoverAttachment bool
	Name                  string
	Pos                   float64
	ShortLink             string
	ShortUrl              string
	Subscribed            bool
	Url                   string
}

type Card struct {
	id   string
	c    *Client
	json *cardJson
}
