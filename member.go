package trello

import (
	"encoding/json"
)

const memberurl = "members"

type memberJson struct {
	Id              string
	Username        string
	FullName        string
	Url             string
	Bio             string
	IdOrganizations []string
	IdBoards        []string
}

// Trello Member.
type Member struct {
	username string
	c        *Client

	json *memberJson
}

func (m *Member) Id() string {
	return m.json.Id
}

func (m *Member) Username() string {
	return m.json.Username
}

func (m *Member) FullName() string {
	return m.json.FullName
}

func (m *Member) Url() string {
	return m.json.Url
}

func (m *Member) Bio() string {
	return m.json.Bio
}

// Get a Member's boards
func (m *Member) Boards() ([]Board, error) {
	b, err := m.c.Request("GET", memberurl+"/"+m.username+"/boards", nil, nil)
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
			c:    m.c,
			json: &bjson,
		}
		out = append(out, board)
	}
	return out, nil
}
