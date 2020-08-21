package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const (
	TokenHeader = "X-Riot-Token"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	apiKey string
	region Region
	doer   Doer
}

func NewClient(apiKey string, region Region, doer Doer) *Client {
	return &Client{
		apiKey: apiKey,
		region: region,
		doer:   doer,
	}
}

func (c *Client) Get(ctx context.Context, endpoint string, target interface{}) error {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.url(endpoint),
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "cannot build http request")
	}
	request.Header.Add(TokenHeader, c.apiKey)
	request.Header.Add("Accept", "application/json")

	response, err := c.doer.Do(request)
	if err != nil {
		return errors.Wrap(err, "client error")
	}

	if err := NewError(response); err != nil {
		return err
	}

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return errors.Wrap(err, "cannot decode json")
	}

	return nil
}

func (c *Client) url(endpoint string) string {
	return "https://" + c.region.GetHost() + endpoint
}
