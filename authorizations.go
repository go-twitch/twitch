package twitch

import (
	"net/http"
	"time"
)

//Scope models a GitHub authorization scope.
type Scope string

//This is the set of scopes for Twitch API v5
const (
	ScopeChannelCheckSubscription = Scope("channel_check_subscription")
	ScopeChannelCommercial        = Scope("channel_commercial")
	ScopeChannelEditor            = Scope("channel_editor")
	ScopeChannelFeedEdit          = Scope("channel_feed_edit")
	ScopeChannelFeedRead          = Scope("channel_feed_read")
	ScopeChannelRead              = Scope("channel_read")
	ScopeChannelStream            = Scope("channel_stream")
	ScopeChannelSubscriptions     = Scope("channel_subscriptions")
	ScopeChatLogin                = Scope("chat_login")
	ScopeUserBlocksEdit           = Scope("user_blocks_edit")
	ScopeUserBlocksRead           = Scope("user_blocks_read")
	ScopeUserFollowsEdit          = Scope("user_follows_edit")
	ScopeUserRead                 = Scope("user_read")
	ScopeUserSubscriptions        = Scope("user_subscriptions")
)

func (c *Client) State() (*StateResponse, *http.Response, error) {
	req, err := c.NewRequest("GET", "", nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StateResponse)
	resp, err := c.Do(req, r)
	if err != nil {
		return nil, resp, err
	}
	return r, resp, nil
}

type StateResponse struct {
	Token struct {
		Authorization struct {
			CreatedAt time.Time `json:"created_at"`
			Scopes    []Scope   `json:"scopes"`
		} `json:"authorization"`
		ClientID string `json:"client_id"`
		UserID   int64  `json:"user_id,string"`
		UserName string `json:"user_name"`
		Valid    bool   `json:"valid"`
	}
}
