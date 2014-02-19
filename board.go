package trello

import (
	"encoding/json"
	"time"
)

const boardurl = "boards"

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

func (b *Board) Desc() string {
	return b.json.Desc
}

func (b *Board) Name() string {
	return b.json.Name
}

func (b *Board) Id() string {
	return b.json.Id
}

func (b *Board) ShortUrl() string {
	return b.json.ShortUrl
}

func (b *Board) Cards() ([]Card, error) {
	js, err := b.c.Request("GET", boardurl+"/"+b.id+"/cards", nil, nil)
	if err != nil {
		return nil, err
	}

	var cards []cardJson

	err = json.Unmarshal(js, &cards)

	if err != nil {
		return nil, err
	}

	var out []Card
	for _, cd := range cards {
		cjson := cd
		card := Card{
			id:   cd.Id,
			c:    b.c,
			json: &cjson,
		}
		out = append(out, card)
	}
	return out, nil
}

func (b *Board) Lists() ([]List, error) {
	js, err := b.c.Request("GET", boardurl+"/"+b.id+"/lists", nil, nil)
	if err != nil {
		return nil, err
	}

	var lists []struct{ Id string }

	err = json.Unmarshal(js, &lists)

	if err != nil {
		return nil, err
	}

	var out []List
	for _, ld := range lists {
		list, err := b.c.List(ld.Id)
		if err != nil {
			return nil, err
		}
		out = append(out, *list)
	}
	return out, nil
}
