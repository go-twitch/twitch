package twitch

import (
	"fmt"
	"net/http"
	"time"
)

type ChannelFeedService service

func (s *ChannelFeedService) CreateFeedPost(id, content string, options *CreateFeedPostOptions) (*ChannelFeedPostResult, *http.Response, error) {
	url := fmt.Sprintf("feed/%v/posts", id)
	v := map[string]interface{}{"content": content}
	if options != nil {
		if options.Share != nil {
			v["share"] = *options.Share
		}
	}
	req, err := s.client.NewRequest("POST", url, &v)
	if err != nil {
		return nil, nil, err
	}
	res := new(ChannelFeedPostResult)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}

type CreateFeedPostOptions struct {
	Share *bool
}

type ChannelFeedPostResult struct {
	Post ChannelFeedPost `json:"post"`
}

type ChannelFeedPost struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Body      string    `json:"body"`
	User      User      `json:"user"`
}

type ChannelFeedPostComments struct {
	Comments []ChannelFeedPostComment `json:"comments"`
}

type ChannelFeedPostComment struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	User      User      `json:"user"`
}
