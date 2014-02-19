package trello

import (
	"encoding/json"
	"time"
	"fmt"
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

func (b *Board) Id() string {
	return b.json.Id
}

func (b *Board) Cards() ([]Card, error) {
	js, err := b.c.Request(boardurl+b.id+"/cards", nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("cards %s\n", js)

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
