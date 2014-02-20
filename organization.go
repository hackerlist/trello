package trello

import (
	"encoding/json"
)

var orgurl = "organizations"

type organizationJson struct {
	Desc string
	//	DescData
	DisplayName string
	Id          string
	//	LogoHash
	Name string
	//	PowerUps
	//	Products
	Url     string
	Website string
}

type Organization struct {
	name string
	c    *Client
	json *organizationJson
}

// Organization retrieves a trello organization
func (c *Client) Organization(name string) (*Organization, error) {
	b, err := c.Request("GET", orgurl+"/"+name, nil, nil)
	if err != nil {
		return nil, err
	}

	o := Organization{
		name: name,
		c:    c,
	}
	err = json.Unmarshal(b, &o.json)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (o *Organization) Desc() string {
	return o.json.Desc
}

func (o *Organization) Members() ([]Member, error) {
	b, err := o.c.Request("GET", orgurl+"/"+o.name+"/members", nil, nil)
	if err != nil {
		return nil, err
	}
	var members []struct{ FullName, Id, Username string }

	err = json.Unmarshal(b, &members)

	if err != nil {
		return nil, err
	}

	var out []Member
	for _, m := range members {
		mem, err := o.c.Member(m.Username)
		if err != nil {
			return nil, err
		}
		out = append(out, *mem)
	}
	return out, nil
}

// Get a Organization's boards
func (o *Organization) Boards() ([]Board, error) {
	b, err := o.c.Request("GET", orgurl+"/"+o.name+"/boards", nil, nil)
	if err != nil {
		return nil, err
	}
	var boards []boardJson

	err = json.Unmarshal(b, &boards)

	if err != nil {
		return nil, err
	}

	var out []Board
	for _, bd := range boards {
		bjson := bd
		board := Board{
			id:   bd.Id,
			c:    o.c,
			json: &bjson,
		}
		out = append(out, board)
	}
	return out, nil
}
