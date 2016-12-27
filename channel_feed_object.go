package twitch

import "time"

type MultipleFeedPostResult struct {
	Cursor uint64            `json:"_cursor,string"`
	Topic  string            `json:"_topic"`
	Total  int               `json:"_total"`
	Posts  []ChannelFeedPost `json:"posts`
}

type ChannelFeedPostResult struct {
	Post ChannelFeedPost `json:"post"`
}

type ChannelFeedPost struct {
	ID        int64     `json:"id,string"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Body      string    `json:"body"`
	User      *User     `json:"user"`
}

type ChannelFeedPostComments struct {
	Comments []ChannelFeedPostComment `json:"comments"`
}

type ChannelFeedPostComment struct {
	ID        int64     `json:"id,string"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	User      User      `json:"user"`
}
