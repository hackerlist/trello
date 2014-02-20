package trello

import (
	"encoding/json"
)

const checklisturl = "checklists"

type CheckItem struct {
	Id   string
	Name string
	//nameData
	Pos   float64
	State string
}

type Checklist struct {
	Id         string
	Pos        float64
	Name       string
	CheckItems []CheckItem
	c          *Client `json:"-"`
}

// Checklist retrieves a checklist by id
func (c *Client) Checklist(id string) (*Checklist, error) {
	b, err := c.Request("GET", checklisturl+"/"+id, nil, nil)

	if err != nil {
		return nil, err
	}

	checklist := Checklist{
		c:  c,
	}

	err = json.Unmarshal(b, &checklist)
	if err != nil {
		return nil, err
	}

	return &checklist, nil
}
