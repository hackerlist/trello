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

type checklistJson struct {
	Id         string
	Pos        float64
	Name       string
	CheckItems []CheckItem
}

type Checklist struct {
	id   string
	c    *Client
	json *checklistJson
}

// Checklist retrieves a checklist by id
func (c *Client) Checklist(id string) (*Checklist, error) {
	b, err := c.Request("GET", checklisturl+"/"+id, nil, nil)

	if err != nil {
		return nil, err
	}

	checklist := Checklist{
		id: id,
		c:  c,
	}

	err = json.Unmarshal(b, &checklist.json)
	if err != nil {
		return nil, err
	}

	return &checklist, nil
}

func (cl *Checklist) Id() string {
	return cl.json.Id
}

func (cl *Checklist) Name() string {
	return cl.json.Name
}
