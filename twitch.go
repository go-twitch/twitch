package twitch

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.twitch.tv/kraken/"
	mediaTypeV5    = "application/vnd.twitchtv.v5+json"
)

type Client struct {
	client      *http.Client
	BaseURL     *url.URL
	UserAgent   string
	ClientID    string
	AccessToken string

	common service

	ChannelFeed *ChannelFeedService
	Users       *UsersService
}

type service struct {
	client *Client
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  client,
		BaseURL: baseURL,
	}
	c.common.client = c

	c.ChannelFeed = (*ChannelFeedService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", mediaTypeV5)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.ClientID != "" {
		req.Header.Set("Client-ID", c.ClientID)
	}
	if c.AccessToken != "" {
		req.Header.Set("Authorization", "OAuth "+c.AccessToken)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := CheckResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}
		err = json.Unmarshal(data, v)
		if err == io.EOF {
			err = nil
		}
	}

	return resp, err
}
