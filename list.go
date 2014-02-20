package trello

import (
	"encoding/json"
)

var listurl = "lists"

type listJson struct {
	Closed  bool
	Id      string
	IdBoard string
	Name    string
	Pos     float64
}

type List struct {
	id   string
	c    *Client
	json *listJson
}

func (c *Client) List(id string) (*List, error) {
	b, err := c.Request("GET", listurl+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	l := List{
		id: id,
		c:  c,
	}

	err = json.Unmarshal(b, &l.json)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (l *List) Id() string {
	return l.json.Id
}

func (l *List) Name() string {
	return l.json.Name
}

func (l *List) Cards() ([]Card, error) {
	js, err := l.c.Request("GET", listurl+"/"+l.id+"/cards", nil, nil)
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
			c:    l.c,
			json: &cjson,
		}
		out = append(out, card)
	}
	return out, nil
}
