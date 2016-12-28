package twitch

import (
	"net/url"
	"strconv"
)

type GetTopGamesOptions url.Values

//Limit set maximum number of objects in array.
//Default is 10.
//Maximum is 100.
func (o GetTopGamesOptions) Limit(limit int) {
	url.Values(o).Set("limit", strconv.Itoa(limit))
}

//Offset set object offset for pagination.
//Default is 0.
func (o GetTopGamesOptions) Offset(limit int) {
	url.Values(o).Set("offset", strconv.Itoa(limit))
}
