package trello

import (
	"time"
)

type Action struct {
	Data struct {
		Text string
	}
	Date            *time.Time
	Id              string
	IdMemberCreator string
	Type            string
	c               *Client `json:"-"`
}
