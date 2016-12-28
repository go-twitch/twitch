package twitch

import "testing"

func TestGames_GetTopGames(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Games.GetTopGames(nil)
	if err != nil {
		t.Fatal(err)
	}
}
