package twitch

import (
	"fmt"
	"net/http"
	"net/url"
)

//ChannelFeedService handles communication with activity related methods of GitHub API.
//https://dev.twitch.tv/docs/v5/reference/channel-feed/
type ChannelFeedService service

//GetMultipleFeedPosts get posts from a specified channel feed.
//https://dev.twitch.tv/docs/v5/reference/channel-feed/#get-multiple-feed-posts
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

//GetFeedPost gets a specified post from a specified channel feed.
//https://dev.twitch.tv/docs/v5/reference/channel-feed/#get-feed-post
func (s *ChannelFeedService) GetFeedPost(channelID, postID int64, opt GetFeedPostOptions) (*ChannelFeedPost, *http.Response, error) {
	urls := fmt.Sprintf("feed/%d/posts/%d", channelID, postID)
	if opt != nil {
		urls += "?" + url.Values(opt).Encode()
	}
	req, err := s.client.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, nil, err
	}
	res := new(ChannelFeedPost)
	resp, err := s.client.Do(req, res)
	return res, resp, err
}

//CreateFeedPost creates a post in a specified channel feed.
//The content of the post is specified in the request body, with a required content parameter.
//https://dev.twitch.tv/docs/v5/reference/channel-feed/#create-feed-post
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

//DeleteFeedPost deletes a specified post in a specified channel feed.
//https://dev.twitch.tv/docs/v5/reference/channel-feed/#delete-feed-post
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
