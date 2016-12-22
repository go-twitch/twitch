package twitch

import (
	"os"
	"testing"
)

func TestClient_State(t *testing.T) {
	setupWithAccess(t)
	s, _, err := client.State()
	if err != nil {
		t.Fatal(err)
	}
	if !s.Token.Valid {
		t.Fatal("invalid token")
	}
	if s.Token.ClientID != os.Getenv("TWITCH_CLIENT_ID") {
		t.Fatal("invalid client id")
	}
}
