package trello

import (
	"encoding/json"
	"net/url"
	"time"
)

const cardurl = "cards"

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
}

type Card struct {
	id   string
	c    *Client
	json *cardJson
}

func (c *Card) Id() string {
	return c.json.Id
}

func (c *Card) Name() string {
	return c.json.Name
}

// Card retrieves a trello card by ID
func (c *Client) Card(id string) (*Card, error) {
	b, err := c.Request("GET", cardurl+"/"+id, nil, nil)

	if err != nil {
		return nil, err
	}

	card := Card{
		id: id,
		c:  c,
	}

	err = json.Unmarshal(b, &card.json)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *Card) AddComment(comment string) error {
	extra := url.Values{"text": {comment}}
	_, err := c.c.Request("POST", cardurl+"/"+c.id+"/actions/comments", nil, extra)
	if err != nil {
		return err
	}
	return nil
}

// Checklists retrieves all checklists from a trello card
func (c *Card) Checklists() ([]Checklist, error) {
	b, err := c.c.Request("GET", cardurl+"/"+c.id+"/checklists", nil, nil)
	if err != nil {
		return nil, err
	}

	var cl []checklistJson

	err = json.Unmarshal(b, &cl)
	if err != nil {
		return nil, err
	}

	var out []Checklist
	for _, ci := range cl {
		cljson := ci
		checklist := Checklist{
			id:   ci.Id,
			c:    c.c,
			json: &cljson,
		}
		out = append(out, checklist)
	}

	return out, nil
}

// Actions retrieves a list of all actions (e.g. events, activity)
// performed on a card
func (c *Card) Actions() ([]Action, error) {
	b, err := c.c.Request("GET", cardurl+"/"+c.id+"/actions", nil, nil)
	if err != nil {
		return nil, err
	}

	var act []actionJson

	err = json.Unmarshal(b, &act)

	if err != nil {
		return nil, err
	}

	var out []Action
	for _, ad := range act {
		ajson := ad
		action := Action{
			id:   ad.Id,
			c:    c.c,
			json: &ajson,
		}
		out = append(out, action)
	}

	return out, nil
}
