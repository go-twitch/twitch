package twitch

import (
	"net/url"
	"strconv"
)

type GetMultipleFeedPostsOptions url.Values

//Limit sets maximum number of most-recent objects to return.
//Default: 10.
//Maximum: 100.
func (o GetMultipleFeedPostsOptions) Limit(limit int) {
	url.Values(o).Set("limit", strconv.Itoa(limit))
}

//Cursor sets tells the server where to start fetching the next set of results in a multi-page response.
func (o GetMultipleFeedPostsOptions) Cursor(cursor uint64) {
	url.Values(o).Set("cursor", strconv.FormatUint(cursor, 10))
}

//Comments sets specifies the number of most-recent comments on posts that are included in the response.
//Default: 5.
//Maximum: 5.
func (o GetMultipleFeedPostsOptions) Comments(comments int) {
	url.Values(o).Set("comments", strconv.Itoa(comments))
}

type GetFeedPostOptions url.Values

//Comments sets specifies the number of most-recent comments on posts that are included in the response. Default: 5. Maximum: 5.
func (o GetFeedPostOptions) Comments(comments int) {
	url.Values(o).Set("comments", strconv.Itoa(comments))
}

type CreateFeedPostOptions struct {
	Content string `json:"content"`
	Share   *bool  `json:"share,omitempty"`
}
