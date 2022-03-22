package helpers

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --name=HTTPClient --inpackage=false --output=./mocks
// HTTPClient interface helps us to mock in test case
type HTTPClient interface {
	MakeRequest(ctx context.Context, httpMethod string, url *url.URL, requestData interface{}) ([]byte, int, error)
}

// HttPConfig config to make rest call
type HttPConfig struct {
	RequestTimeout int64
	SSLEnabled     bool
	Username       string
	Password 	string
}

type client struct {
	config     HttPConfig
	httpClient *http.Client
}

// NewHTTPClient ...
func NewHTTPClient(config HttPConfig) (HTTPClient, error) {
	return &client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Duration(config.RequestTimeout) * time.Second,
		},
	}, nil
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// MakeRequest to make rest http call to coin market
func (c *client) MakeRequest(ctx context.Context, httpMethod string, url *url.URL, requestData interface{}) ([]byte, int, error) {
	log := logrus.Logger{}
	var statusCode int
	/* #nosec */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.config.SSLEnabled},
	}
	c.httpClient.Transport = tr

	var req *http.Request
	var err error
	req, err = http.NewRequestWithContext(ctx, httpMethod, url.String(), nil)
	if err != nil {
		log.WithError(err).Infof("error occurred while making request")
		return nil, statusCode, fmt.Errorf("error occurred while making token request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Add("authorization", "Bearer " + os.Getenv("COIN_MARKET_TOKEN"))

	resp, err := c.do(req)
	if err != nil {
		return nil, statusCode, fmt.Errorf("failed to do request, %w", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, statusCode, fmt.Errorf("failed to read body: %w", err)
	}
	log.Debugf("Status Code: %v", resp.StatusCode)
	log.Debugf("Response Body: %s", bodyBytes)

	return bodyBytes, resp.StatusCode, nil
}

// HideCredentials masks sensitive information in the Request log
func (c *client) HideCredentials(dump []byte, clientSecret string) []byte {
	resetClientSecret := regexp.MustCompile(clientSecret)
	cleanDataClientSecret := resetClientSecret.ReplaceAllString(string(dump), "********")
	return []byte(cleanDataClientSecret)
}