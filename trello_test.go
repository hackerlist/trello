package trello

import (
	"testing"
	"sync"
	"io/ioutil"
	"encoding/json"
)

type Creds struct {
	Key, Secret, Token, Member string
}

var (
	creds Creds
	load sync.Once
)

func loadCreds() {
	b, err := ioutil.ReadFile("trello.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &creds)

	if err != nil {
		panic(err)
	}
}

func setupTest() {
	load.Do(loadCreds)

	if creds.Key == "" {
		panic("no key")
	}
}

func TestMember(t *testing.T) {
	setupTest()

	c := New(creds.Key, creds.Secret, creds.Token)
	m, err := c.Member(creds.Member)

	if err != nil {
		t.Errorf("member request: %s", err)
	} else {
		t.Logf("%s name %s bio %s", m.Username(), m.FullName(), m.Bio())
		t.Logf("%+v", m.json)
	}

	if boards, err := m.Boards(); err != nil {
		t.Errorf("board request: %s", err)
	} else {
		for _, b := range boards {
			t.Logf("board %+v", b.json)
		}
	}
}
