package twitch

import (
	"os"
	"testing"
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
	client = NewClient(nil, &OAuth2Config{ClientID: clientID}, nil)
	client.UserAgent = "go-twitch test"
}

func setupWithAccess(t *testing.T) {
	setup(t)
	accessToken := os.Getenv("TWITCH_ACCESS_TOKEN")
	if accessToken == "" {
		t.Log("You need to set TWITCH_ACCESS_TOKEN!!")
		t.FailNow()
	}
	client.Token = &OAuth2Token{
		AccessToken: accessToken,
	}
}
