package ipfs

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ipfsCli struct {
	rpcURL  string
	httpCli *http.Client
}

type IPFSCliOption func(*ipfsCli)

func WithHTTPClient(httpCli *http.Client) IPFSCliOption {
	return func(opts *ipfsCli) {
		opts.httpCli = httpCli
	}
}

func NewIPFSClient(rpcURL string, opts ...IPFSCliOption) *ipfsCli {
	c := &ipfsCli{
		rpcURL:  rpcURL,
		httpCli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *ipfsCli) Cat(path string) (io.ReadCloser, error) {
	rpcURL := fmt.Sprintf("%s/ipfs/%s",
		strings.TrimSuffix(c.rpcURL, "/"), url.QueryEscape(path),
	)
	resp, err := c.httpCli.Get(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send IPFS get request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code from IPFS getway: %d",
			resp.StatusCode)
	}
	return resp.Body, nil
}
