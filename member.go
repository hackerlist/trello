package trello

import (
	"encoding/json"
	"net/url"
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

// Member retrieves a trello member's (user) info
func (c *Client) Member(username string) (*Member, error) {
	extra := url.Values{"fields": {"username,fullName,url,bio,idBoards,idOrganizations"}}
	b, err := c.Request("GET", memberurl+"/"+username, nil, extra)
	if err != nil {
		return nil, err
	}

	m := Member{
		username: username,
		c:        c,
	}

	err = json.Unmarshal(b, &m.json)
	if err != nil {
		return nil, err
	}

	return &m, nil
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
