package twitch

import (
	"net/http"
	"net/url"
)

//GamesService handles communication with the activity related methods of the GitHub API.
//https://dev.twitch.tv/docs/v5/reference/games/
type GamesService service

//GetTopGames get games by number of current viewers on Twitch.
//https://dev.twitch.tv/docs/v5/reference/games/#get-top-games
func (s GamesService) GetTopGames(opt GetTopGamesOptions) (*TopGamesResult, *http.Response, error) {
	urls := "games/top"
	if opt != nil {
		urls += "?" + url.Values(opt).Encode()
	}
	req, err := s.client.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, nil, err
	}
	res := new(TopGamesResult)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}
