package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
)

func New(ctx context.Context, options ...Option) (*Client, error) {
	c := &Client{
		timeout:  DefaultTimeout,
		ctx:      ctx,
		unsecure: false,
	}

	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, fmt.Errorf("When Creating Dreamland HTTP Client, parsing options failed with: %s", err.Error())
		}
	}

	c.client = &http.Client{
		Timeout: c.timeout,
	}

	if c.unsecure == false {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		}
	} else {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	c.auth_header = fmt.Sprintf("%s %s", c.provider, c.token)

	return c, nil
}
