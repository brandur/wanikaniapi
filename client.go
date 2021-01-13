package wanikaniapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const WaniKaniAPIURL = "https://api.wanikani.com"
const WaniKaniRevision = "20170710"

type Client struct {
	APIToken string
	Logger   LeveledLoggerInterface

	baseURL    string
	httpClient *http.Client
}

type ClientConfig struct {
	APIToken string
	Logger   LeveledLoggerInterface
}

func NewClient(config *ClientConfig) *Client {
	var logger LeveledLoggerInterface

	if config.Logger == nil {
		logger = &LeveledLogger{Level: LevelError}
	} else {
		logger = config.Logger
	}

	return &Client{
		APIToken: config.APIToken,
		Logger:   logger,

		baseURL:    WaniKaniAPIURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) request(method, path, query string, data interface{}) error {
	url := c.baseURL + path
	if query != "" {
		url += "?" + query
	}

	c.Logger.Debugf("Requesting URL: %v (revision: %v)", url, WaniKaniRevision)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Wanikani-Revision", WaniKaniRevision)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status from API: %v", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
}
