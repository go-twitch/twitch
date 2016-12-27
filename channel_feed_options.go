package twitch

import (
	"net/url"
	"strconv"
)

type CreateFeedPostOptions struct {
	Content string `json:"content"`
	Share   *bool  `json:"share,omitempty"`
}

type GetMultipleFeedPostsOptions url.Values

func (o GetMultipleFeedPostsOptions) SetLimit(limit int) {
	url.Values(o).Set("limit", strconv.Itoa(limit))
}

func (o GetMultipleFeedPostsOptions) SetCursor(cursor uint64) {
	url.Values(o).Set("cursor", strconv.FormatUint(cursor, 10))
}

func (o GetMultipleFeedPostsOptions) SetComments(comments int) {
	url.Values(o).Set("comments", strconv.Itoa(comments))
}
