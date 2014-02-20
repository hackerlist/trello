package trello

import (
	"encoding/json"
	"net/url"
	"time"
)

const cardurl = "cards"

type Card struct {
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
	IdList         string
	IdMembers      []string
	IdMembersVoted []string
	IdShort        float64
	//	Labels                []string
	ManualCoverAttachment bool
	Name                  string
	Pos                   float64
	ShortLink             string
	ShortUrl              string
	Subscribed            bool
	Url                   string
	c                     *Client `json"-"`
}

// Card retrieves a trello card by ID
func (c *Client) Card(id string) (*Card, error) {
	b, err := c.Request("GET", cardurl+"/"+id, nil, nil)

	if err != nil {
		return nil, err
	}

	card := Card{
		c: c,
	}

	err = json.Unmarshal(b, &card)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *Card) AddComment(comment string) error {
	extra := url.Values{"text": {comment}}
	_, err := c.c.Request("POST", cardurl+"/"+c.Id+"/actions/comments", nil, extra)
	if err != nil {
		return err
	}
	return nil
}

// AddChecklist created a new checklist on the card.
func (c *Card) AddChecklist(name string) (*Checklist, error) {
	qp := url.Values{"name":{name}}
	b, err := c.c.Request("POST", cardurl+"/"+c.Id+"/checklists", nil, qp)
	if err != nil {
		return nil, err
	}

	var cl *Checklist
	err = json.Unmarshal(b, &cl)
	return cl, err
}

// Checklists retrieves all checklists from a trello card
func (c *Card) Checklists() ([]*Checklist, error) {
	b, err := c.c.Request("GET", cardurl+"/"+c.Id+"/checklists", nil, nil)
	if err != nil {
		return nil, err
	}

	var checklists []*Checklist

	err = json.Unmarshal(b, &checklists)
	if err != nil {
		return nil, err
	}

	for _, checklist := range checklists {
		checklist.c = c.c
		for _, ci := range checklist.CheckItems {
			ci.c = c.c
		}
	}

	return checklists, nil
}

// Actions retrieves a list of all actions (e.g. events, activity)
// performed on a card
func (c *Card) Actions() ([]*Action, error) {
	b, err := c.c.Request("GET", cardurl+"/"+c.Id+"/actions", nil, nil)
	if err != nil {
		return nil, err
	}

	var act []*Action

	err = json.Unmarshal(b, &act)

	if err != nil {
		return nil, err
	}

	for _, a := range act {
		a.c = c.c
	}

	return act, nil
}
