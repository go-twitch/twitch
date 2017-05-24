package twitch

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	oauth2TokenURL  = "https://api.twitch.tv/kraken/oauth2/token"
	oauth2AuthURL   = "https://api.twitch.tv/kraken/oauth2/authorize"
	oauth2RevokeURL = "https://api.twitch.tv/kraken/oauth2/revoke"
)

type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []Scope
}

func (c *OAuth2Config) AuthCodeURL(state string) string {
	v := url.Values{
		"client_id":     {c.ClientID},
		"redirect_uri":  {c.RedirectURI},
		"response_type": {"code"},
	}
	if c.Scopes != nil {
		s := make([]string, len(c.Scopes))
		for i, scope := range c.Scopes {
			s[i] = string(scope)
		}
		v.Set("scope", strings.Join(s, " "))
	}
	if state != "" {
		v.Set("state", state)
	}
	return oauth2AuthURL + "?" + v.Encode()
}

func (c *OAuth2Config) Exchange(client *http.Client, code, state string) (*OAuth2Token, error) {
	if client == nil {
		client = http.DefaultClient
	}
	v := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {c.RedirectURI},
	}
	if state != "" {
		v.Add("state", state)
	}
	url := oauth2TokenURL + "?" + v.Encode()
	resp, err := client.Post(url, "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	token := new(OAuth2Token)
	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, err
	}
	return token, nil
}

func (c *OAuth2Config) Revoke(client *http.Client, token *OAuth2Token) error {
	if client == nil {
		client = http.DefaultClient
	}
	v := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"token":         {token.AccessToken},
	}
	url := oauth2RevokeURL + "?" + v.Encode()
	resp, err := client.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("twitch: failed revoke token")
}

type OAuth2Token struct {
	AccessToken string   `json:"access_token"`
	Scope       []string `json:"scope"`
}
