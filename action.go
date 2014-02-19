package trello

import (
	"time"
)

type actionJson struct {
	Data struct {
		Text string
	}
	Date            *time.Time
	Id              string
	IdMemberCreator string
	Type            string
}

type Action struct {
	id   string
	c    *Client
	json *actionJson
}

func (a *Action) DataText() string {
	return a.json.Data.Text
}

func (a *Action) Id() string {
	return a.json.Id
}

func (a *Action) IdMemberCreator() string {
	return a.json.IdMemberCreator
}

func (a *Action) Type() string {
	return a.json.Type
}
