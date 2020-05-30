package galley_client

import (
	"bytes"
	"encoding/json"
	"github.com/skvoch/galley/internal/galley/model"
	"net/http"
)

func New() *Client {
	return &Client{
		URL: "https://galley-jr6l7s7e6a-uc.a.run.app",
		//URL: "http://127.0.0.1:8080",

		Client: http.Client{},
	}
}

type Client struct {
	URL    string
	Client http.Client
}

func (c *Client) Handshake(user *model.User) error {

	data, err := json.Marshal(user)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.URL+HandshakeEndpoint, bytes.NewReader(data))

	if err != nil {
		return err
	}

	_, err = c.Client.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendStats(stats *model.ClickStats) error {
	data, err := json.Marshal(stats)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.URL+SendStatsEndpoint, bytes.NewReader(data))

	if err != nil {
		return err
	}

	_, err = c.Client.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetCurrentPush() (*model.Push, error) {
	req, err := http.NewRequest(http.MethodGet, c.URL+GetPushEndpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	push := &model.Push{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(push); err != nil {
		return nil, err
	}

	return push, nil
}
