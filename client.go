package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const trellourl = "https://api.trello.com/1/"

type Client struct {
	apikey    string
	apisecret string
	apitoken  string
}

func New(key, secret, token string) *Client {
	return &Client{key, secret, token}
}

func (c *Client) Request(method, function string, postbody io.Reader, extra url.Values) ([]byte, error) {
	postdata := url.Values{"key": {c.apikey}, "token": {c.apitoken}}
	for k, v := range extra {
		postdata[k] = v
	}
	url := trellourl + function + "?" + postdata.Encode()
	req, err := http.NewRequest(method, url, postbody)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode > 300 {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	return body, nil
}

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

// CreateBoard creats a new board with the given name. Extra options can be passed
// through the extra parameter. For details on options, see
// https://trello.com/docs/api/board/index.html#post-1-boards
func (c *Client) CreateBoard(name string, extra url.Values) (*Board, error) {
	qp := url.Values{"name": {name}}
	for k, v := range extra {
		qp[k] = v
	}

	b, err := c.Request("POST", boardurl, nil, qp)
	if err != nil {
		return nil, err
	}

	board := Board{
		c: c,
	}

	err = json.Unmarshal(b, &board.json)
	if err != nil {
		return nil, err
	}

	board.id = board.json.Id

	return &board, nil
}

func (c *Client) Board(id string) (*Board, error) {
	b, err := c.Request("GET", boardurl+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	board := Board{
		id: id,
		c:  c,
	}

	err = json.Unmarshal(b, &board.json)
	if err != nil {
		return nil, err
	}

	return &board, nil
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

func getfield(js []byte, field string) (string, error) {
	var i interface{}
	err := json.Unmarshal(js, &i)
	if err != nil {
		return "", err
	}

	ma, ok := i.(map[string]interface{})

	if !ok {
		return "", fmt.Errorf("json not a dictionary")
	}

	f, ok := ma[field]
	if !ok {
		return "", fmt.Errorf("no field %s", field)
	}

	str, ok := f.(string)

	if !ok {
		return "", fmt.Errorf("field %s not a string", field)
	}

	return str, nil
}
