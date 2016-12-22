package twitch

import (
	"fmt"
	"net/http"
	"net/url"
)

type UsersService service

func (s *UsersService) GetUser() (*UserSelf, *http.Response, error) {
	v := new(UserSelf)
	resp, err := s.getUser("user", v)
	return v, resp, err
}

func (s *UsersService) GetUserByID(userID int64) (*User, *http.Response, error) {
	v := new(User)
	resp, err := s.getUser(fmt.Sprintf("users/%d", userID), v)
	return v, resp, err
}

func (s *UsersService) getUser(url string, v interface{}) (*http.Response, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, v)
	return resp, err
}

func (s *UsersService) GetUserEmotes(userID int64) (*EmoticonSetsResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%d/emotes", userID)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(EmoticonSetsResponse)
	resp, err := s.client.Do(req, v)
	return v, resp, err
}

func (s *UsersService) CheckUserSubscriptionByChannel(userID, channelID int64) (*EmoticonSetsResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%d/subscriptions/%d", userID, channelID)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(EmoticonSetsResponse)
	resp, err := s.client.Do(req, v)
	if resp.StatusCode == http.StatusNotFound {
		return nil, resp, nil
	}
	return v, resp, err
}

func (s *UsersService) GetUserFollows(userID int64, opt GetUserFollowsOptions) (*UserFollows, *http.Response, error) {
	urls := fmt.Sprintf("users/%d/follows/channels", userID)
	if opt != nil {
		urls += "?" + ((url.Values)(opt)).Encode()
	}
	req, err := s.client.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(UserFollows)
	resp, err := s.client.Do(req, v)
	return v, resp, err
}

func (s *UsersService) CheckUserFollowsByChannel(userID, channelID int64) (*UserFollow, *http.Response, error) {
	url := fmt.Sprintf("users/%d/follows/channels/%d", userID, channelID)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(UserFollow)
	resp, err := s.client.Do(req, v)
	if resp.StatusCode == http.StatusNotFound {
		return nil, resp, nil
	}
	return v, resp, err
}

func (s *UsersService) FollowChannel(userID, channelID int64, opt *FollowChannelOptions) (*UserFollow, *http.Response, error) {
	url := fmt.Sprintf("users/%d/follows/channels/%d", userID, channelID)
	req, err := s.client.NewRequest("PUT", url, opt)
	if err != nil {
		return nil, nil, err
	}
	v := new(UserFollow)
	resp, err := s.client.Do(req, v)
	return v, resp, err
}

func (s *UsersService) UnfollowChannel(userID, channelID int64) (*http.Response, error) {
	url := fmt.Sprintf("users/%d/follows/channels/%d", userID, channelID)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

func (s *UsersService) GetUserBlockList(userID int64, opt GetUserBlocksOptions) (*UserBlocks, *http.Response, error) {
	urls := fmt.Sprintf("users/%d/blocks", userID)
	if opt != nil {
		urls += "?" + ((url.Values)(opt)).Encode()
	}
	req, err := s.client.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(UserBlocks)
	resp, err := s.client.Do(req, v)
	return v, resp, err
}

func (s *UsersService) BlockUser(userID, channelID int64) (*UserBlock, *http.Response, error) {
	url := fmt.Sprintf("users/%d/blocks/%d", userID, channelID)
	req, err := s.client.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, nil, err
	}
	v := new(UserBlock)
	resp, err := s.client.Do(req, v)
	return v, resp, err
}

func (s *UsersService) UnblockUser(userID, channelID int64) (*http.Response, error) {
	url := fmt.Sprintf("users/%d/blocks/%d", userID, channelID)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
