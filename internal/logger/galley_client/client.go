package galley_client

import (
	"bytes"
	"encoding/json"
	"github.com/skvoch/galley/internal/galley/model"
	"net/http"
)

func New() *Client {
	return &Client{
		URL:    "https://galley-jr6l7s7e6a-uc.a.run.app",
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
