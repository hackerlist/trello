package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const trellourl = "https://api.trello.com/1"

type Client struct {
	apikey    string
	apisecret string
	apitoken  string
}

func New(key, secret, token string) *Client {
	return &Client{key, secret, token}
}

func (c *Client) Request(function string, extra url.Values) ([]byte, error) {
	postdata := url.Values{"key": {c.apikey}, "token": {c.apitoken}}
	for k, v := range extra {
		postdata[k] = v
	}
	url := trellourl + function + "?" + postdata.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func (c *Client) Member(username string) (*Member, error) {
	extra := url.Values{"fields": {"username,fullName,url,bio,idBoards,idOrganizations"}}
	b, err := c.Request(memberurl+username, extra)
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
