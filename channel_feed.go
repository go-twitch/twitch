package twitch

import (
	"fmt"
	"net/http"
	"net/url"
)

type ChannelFeedService service

func (s *ChannelFeedService) GetMultipleFeedPosts(channelID int64, opt GetMultipleFeedPostsOptions) (*MultipleFeedPostResult, *http.Response, error) {
	urls := fmt.Sprintf("feed/%d/posts", channelID)
	if opt != nil {
		urls += "?" + url.Values(opt).Encode()
	}
	req, err := s.client.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, nil, err
	}
	res := new(MultipleFeedPostResult)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}

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
