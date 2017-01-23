package twitch

import (
	"os"
	"testing"

	"golang.org/x/oauth2"
)

var client *Client

const (
	//user id
	testUser = 142484874 //go_twitch
	//channel id
	esl         = 67834893
	dota2ti     = 35630634
	riotgames   = 36029255
	dotamajor   = 106460795
	playparagon = 104243230
)

func setup(t *testing.T) {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	if clientID == "" {
		t.Log("You need to set TWITCH_CLIENT_ID!!")
		t.FailNow()
	}
	client = NewClient(nil)
	client.ClientID = clientID
	client.UserAgent = "go-twitch test"
}

func setupWithAccess(t *testing.T) {
	accessToken := os.Getenv("TWITCH_ACCESS_TOKEN")
	if accessToken == "" {
		t.Log("You need to set TWITCH_ACCESS_TOKEN!!")
		t.FailNow()
	}

	token := &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "OAuth",
	}
	cli := oauth2.NewClient(nil, oauth2.StaticTokenSource(token))
	client = NewClient(cli)
	client.UserAgent = "go-twitch test"
}
