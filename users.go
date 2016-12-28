package twitch

import (
	"fmt"
	"net/http"
	"net/url"
)

//UsersService handles communication with activity related methods of GitHub API.
//https://dev.twitch.tv/docs/v5/reference/users/
type UsersService service

//GetUser gets a user object based on the OAuth token provided.
//If the user’s Twitch-registered email address is not verified, null is returned.
//Get User returns more data than Get User by ID, because Get User is privileged.
//https://dev.twitch.tv/docs/v5/reference/users/#get-user
func (s *UsersService) GetUser() (*UserSelf, *http.Response, error) {
	v := new(UserSelf)
	resp, err := s.getUser("user", v)
	return v, resp, err
}

//GetUserByID gets a specified user object.
//https://dev.twitch.tv/docs/v5/reference/users/#get-user-by-id
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

//GetUserEmotes gets a list of the emojis and emoticons that the specified user can use in chat.
//These are both the globally available ones and the channel-specific ones (which can be accessed by any user subscribed to the channel).
//https://dev.twitch.tv/docs/v5/reference/users/#get-user-emotes
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

//CheckUserSubscriptionByChannel checks if a specified user is subscribed to a specified channel.
//Intended for viewers.
//There is an error response (422 Unprocessable Entity) if the channel does not have a subscription program.
//https://dev.twitch.tv/docs/v5/reference/users/#check-user-subscription-by-channel
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

//GetUserFollows gets a list of all channels followed by a specified user, sorted by the date when they started following each channel.
//https://dev.twitch.tv/docs/v5/reference/users/#get-user-follows
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

//CheckUserFollowsByChannel checks if a specified user follows a specified channel.
//If the user is following the channel, a follow object is returned.
//https://dev.twitch.tv/docs/v5/reference/users/#check-user-follows-by-channel
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

//FollowChannel adds a specified user to the followers of a specified channel.
//There is an error response (422 Unprocessable Entity) if the channel could not be followed.
//https://dev.twitch.tv/docs/v5/reference/users/#follow-channel
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

//UnfollowChannel deletes a specified user from the followers of a specified channel.
//https://dev.twitch.tv/docs/v5/reference/users/#unfollow-channel
func (s *UsersService) UnfollowChannel(userID, channelID int64) (*http.Response, error) {
	url := fmt.Sprintf("users/%d/follows/channels/%d", userID, channelID)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

//GetUserBlockList gets a user’s block list.
//https://dev.twitch.tv/docs/v5/reference/users/#get-user-block-list
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

//BlockUser blocks the target user.
//https://dev.twitch.tv/docs/v5/reference/users/#block-user
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

//UnblockUser unblocks the target user.
//https://dev.twitch.tv/docs/v5/reference/users/#unblock-user
func (s *UsersService) UnblockUser(userID, channelID int64) (*http.Response, error) {
	url := fmt.Sprintf("users/%d/blocks/%d", userID, channelID)
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
