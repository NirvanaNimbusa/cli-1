//go:generate go run generate/protocol.go

package ga

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var trackingIDMatcher = regexp.MustCompile(`^UA-\d+-\d+$`)

func NewClient(trackingID string) (*Client, error) {
	if !trackingIDMatcher.MatchString(trackingID) {
		return nil, fmt.Errorf("Invalid Tracking ID: %s", trackingID)
	}
	return &Client{
		UseTLS:             true,
		HttpClient:         http.DefaultClient,
		protocolVersion:    "1",
		protocolVersionSet: true,
		trackingID:         trackingID,
		clientID:           "go-ga",
		clientIDSet:        true,
	}, nil
}

type hitType interface {
	addFields(url.Values) error
}

func (c *Client) SendWithContext(ctx context.Context, h hitType) error {

	cpy := c.Copy()

	v := url.Values{}

	cpy.setType(h)

	err := cpy.addFields(v)
	if err != nil {
		return err
	}

	err = h.addFields(v)
	if err != nil {
		return err
	}

	url := ""
	if cpy.UseTLS {
		url = "https://www.google-analytics.com/collect"
	} else {
		url = "http://ssl.google-analytics.com/collect"
	}

	str := v.Encode()
	buf := bytes.NewBufferString(str)

	req, err := http.NewRequestWithContext(ctx, "POST", url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("Rejected by Google with code %d", resp.StatusCode)
	}

	// fmt.Printf("POST %s => %d\n", str, resp.StatusCode)

	return nil
}
