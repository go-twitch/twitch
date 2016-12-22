package twitch

import (
	"net/url"
	"strconv"
)

type GetUserFollowsOptions url.Values

func (o GetUserFollowsOptions) Limit(limit int) {
	url.Values(o).Set("limit", strconv.Itoa(limit))
}

func (o GetUserFollowsOptions) Offset(offset int) {
	url.Values(o).Set("offset", strconv.Itoa(offset))
}

func (o GetUserFollowsOptions) Direction(direction string) {
	url.Values(o).Set("direction", direction)
}

func (o GetUserFollowsOptions) SortBy(sortBy string) {
	url.Values(o).Set("sortby", sortBy)
}

type FollowChannelOptions struct {
	Notifications *bool `json:"notifications,omitempty"`
}

type GetUserBlocksOptions url.Values

func (o GetUserBlocksOptions) Limit(limit int) {
	url.Values(o).Set("limit", strconv.Itoa(limit))
}

func (o GetUserBlocksOptions) Offset(offset int) {
	url.Values(o).Set("offset", strconv.Itoa(offset))
}
