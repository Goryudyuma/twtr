package twtr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Decode(resp *http.Response, data interface{}) error {
	if err := NewErrors(resp); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Client) ParseURL(path string, params Values) string {
	for k, v := range params {
		if len(k) <= 0 || k[0] != ':' {
			continue
		}
		path = strings.Replace(path, k, v, -1)
		delete(params, k)
	}
	return fmt.Sprintf(c.URLFormat, path)
}

func (c *Client) GET(path string, params Values, data interface{}) error {
	uri := c.ParseURL(path, params)
	resp, err := c.OAuthClient.Get(nil, c.AccessToken, uri, params.ToURLValues())
	if err != nil {
		return err
	}
	return Decode(resp, data)
}

func (c *Client) POST(path string, params Values, data interface{}) error {
	uri := c.ParseURL(path, params)
	resp, err := c.OAuthClient.Post(nil, c.AccessToken, uri, params.ToURLValues())
	if err != nil {
		return err
	}
	return Decode(resp, data)
}