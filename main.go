package servicediscovery

import (
	consul "github.com/hashicorp/consul/api"
)

type Client struct {
	*consul.Client
}

func NewClient(address string) (*Client, error) {
	client, err := consul.NewClient(&consul.Config{
		Address:    address,
		Scheme:     "",
		Datacenter: "",
		Transport:  nil,
		HttpClient: nil,
		HttpAuth:   nil,
		WaitTime:   0,
		Token:      "",
		TokenFile:  "",
		Namespace:  "",
		TLSConfig:  consul.TLSConfig{},
	})
	if err != nil {
		return nil, err
	}
	return &Client{client}, err
}

func (c *Client) GetNodes(svc string, tag string) ([]string, error) {
		service, _, err := c.Catalog().Service(svc, tag, _)
		if err != nil {
			return nil, err
		}

		var discovery []string
		for _, value := range service {
			address := value.Address
			discovery = append(discovery, address)
		}
		return discovery, nil
	}

}
