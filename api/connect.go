package api

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Client represents interface for Kite Connect client.
type Client struct {
	password    string
	twoFA       string
	debug       bool
	vendor_code string
	api_secret  string
	imei        string
	httpClient  HTTPClient
}

const (
	name           string        = "shoonyaapi-go"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://api.shoonya.com/NorenWClientTP/"
	// NorenApi.__init__(self, host='https://api.shoonya.com/NorenWClientTP/',
	// websocket='wss://api.shoonya.com/NorenWSTP/',
	// eodhost="https://shoonya.finvasia.com/chartApi/getdata/")
)

// Login using shoonya API client.
func Login(userId string, password string, twoFA string, vendor_code string, api_secret string, imei string) {
	client := &Client{}

	// Create a default http handler with default timeout.
	client.SetHTTPClient(&http.Client{
		Timeout:   requestTimeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	})

	return client
}

// SetHTTPClient overrides default http handler with a custom one.
// This can be used to set custom timeouts and transport.
func (c *Client) SetHTTPClient(h *http.Client) {
	c.httpClient = NewHTTPClient(h, nil, c.debug)
}

// SetDebug sets debug mode to enable HTTP logs.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
	c.httpClient.GetClient().debug = debug
}

// SetBaseURI overrides the base Shoonya api endpoint with custom url.
func (c *Client) SetBaseURI(baseURI string) {
	c.baseURI = baseURI
}

// SetTimeout sets request timeout for default http client.
func (c *Client) SetTimeout(timeout time.Duration) {
	hClient := c.httpClient.GetClient().client
	hClient.Timeout = timeout
}

// SetAccessToken sets the access token to the Kite Connect instance.
func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *Client) doEnvelope(method, uri string, params map[string]interface{}, headers http.Header, v interface{}, authorization ...bool) error {
	if params == nil {
		params = map[string]interface{}{}
	}

	// Send custom headers set
	if headers == nil {
		headers = map[string][]string{}
	}

	localIp, publicIp, mac, err := getIpAndMac()

	if err != nil {
		return err
	}

	// Add Kite Connect version to header
	headers.Add("Content-Type", "application/json")
	headers.Add("X-ClientLocalIP", localIp)
	headers.Add("X-ClientPublicIP", publicIp)
	headers.Add("X-MACAddress", mac)
	headers.Add("Accept", "application/json")
	headers.Add("X-UserType", "USER")
	headers.Add("X-SourceID", "WEB")
	headers.Add("X-PrivateKey", c.apiKey)
	if authorization != nil && authorization[0] {
		headers.Add("Authorization", "Bearer "+c.accessToken)
	}

	return c.httpClient.DoEnvelope(method, c.baseURI+uri, params, headers, v)
}
