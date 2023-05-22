package norenapigo

import (
	"crypto/tls"
	_ "fmt"
	"net/http"
	"time"
)

// Client represents interface for Kite Connect client.
type Client struct {
	clientCode  string
	password    string
	accessToken string
	debug       bool
	baseURI     string
	apiKey      string
	httpClient  HTTPClient
}

const (
	name           string        = "shoonya-go"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://api.shoonya.com/NorenWClientTP/"
	// websockerURI string = "wss://api.shoonya.com/NorenWSTP/"
)

// New creates a new noren API client.
func New(clientCode string, password string, apiKey string) *Client {
	client := &Client{
		clientCode: clientCode,
		password:   password,
		apiKey:     apiKey,
		baseURI:    baseURI,
	}

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

// SetBaseURI overrides the base norenAPI endpoint with custom url.
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

	// localIp, publicIp, mac, err := getIpAndMac()

	// if err != nil {
	// 	return err
	// }

	// Add Kite Connect version to header
	headers.Add("Content-Type", "application/json")

	headers.Add("charset", "utf-8")

	// fmt.Printf("\n--> Method : %s \nURL : %s \nParam : %v\n Header : %v\n, V: %v\n AccessToken: %v\n", method, c.baseURI+uri, params, headers, v, c.accessToken)

	return c.httpClient.DoEnvelope(method, c.baseURI+uri, params, headers, c.accessToken, v)
}
