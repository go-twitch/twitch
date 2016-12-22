package twitch

import (
	"fmt"
	"net/http"
	"time"
)

type ChannelFeedService service

func (s *ChannelFeedService) CreateFeedPost(channelID int64, content string, opt *CreateFeedPostOptions) (*ChannelFeedPostResult, *http.Response, error) {
	url := fmt.Sprintf("feed/%d/posts", channelID)
	if opt == nil {
		opt = new(CreateFeedPostOptions)
	}
	if opt.Content == "" {
		opt.Content = content
	}
	req, err := s.client.NewRequest("POST", url, &opt)
	if err != nil {
		return nil, nil, err
	}
	res := new(ChannelFeedPostResult)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}

func (s *ChannelFeedService) DeleteFeedPost(channelID, feedID int64) (*ChannelFeedPost, *http.Response, error) {
	url := fmt.Sprintf("feed/%d/posts/%d", channelID, feedID)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, nil, err
	}
	res := new(ChannelFeedPost)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}

type CreateFeedPostOptions struct {
	Content string `json:"content"`
	Share   *bool  `json:"share,omitempty"`
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
