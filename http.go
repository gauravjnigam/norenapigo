package norenapigo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// HTTPClient represents an HTTP client.
type HTTPClient interface {
	Do(method, rURL string, params map[string]interface{}, headers http.Header, accessToken string) (HTTPResponse, error)
	DoEnvelope(method, url string, params map[string]interface{}, headers http.Header, accessToken string, obj interface{}) error
	GetClient() *httpClient
}

// httpClient is the default implementation of HTTPClient.
type httpClient struct {
	client *http.Client
	hLog   *log.Logger
	debug  bool
}

// HTTPResponse encompasses byte body  + the response of an HTTP request.
type HTTPResponse struct {
	Body     []byte
	Response *http.Response
}

type envelope struct {
	Status    bool        `json:"status"`
	ErrorCode string      `json:"errorcode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// NewHTTPClient returns a self-contained HTTP request object
// with underlying keep-alive transport.
func NewHTTPClient(h *http.Client, hLog *log.Logger, debug bool) HTTPClient {
	if hLog == nil {
		hLog = log.New(os.Stdout, "base.HTTP: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if h == nil {
		h = &http.Client{
			Timeout: time.Duration(5) * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second * time.Duration(5),
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	return &httpClient{
		hLog:   hLog,
		client: h,
		debug:  debug,
	}
}

// Do executes an HTTP request and returns the response.
func (h *httpClient) Do(method, rURL string, params map[string]interface{}, headers http.Header, accessToken string) (HTTPResponse, error) {
	var (
		resp       = HTTPResponse{}
		postParams io.Reader
		err        error
	)

	if method == http.MethodPost && params != nil {
		jsonParams, err := json.Marshal(params)

		if err != nil {
			return resp, err
		}
		jData := "jData=" + string(jsonParams)
		if len(accessToken) > 0 {
			// fmt.Printf("Adding access token on payload...")
			jData = jData + "&jKey=" + accessToken
		}
		postParams = bytes.NewBuffer([]byte(jData))
		// fmt.Printf("PostParam : %v", postParams)
	}

	req, err := http.NewRequest(method, rURL, postParams)

	if err != nil {
		h.hLog.Printf("Request preparation failed: %v", err)
		return resp, err
	}

	if headers != nil {
		req.Header = headers
	}

	// If a content-type isn't set, set the default one.
	if req.Header.Get("Content-Type") == "" {
		if method == http.MethodPost || method == http.MethodPut {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}
	}

	// If the request method is GET or DELETE, add the params as QueryString.
	//if method == http.MethodGet || method == http.MethodDelete {
	//	req.URL.RawQuery = params.Encode()
	//}

	r, err := h.client.Do(req)
	if err != nil {
		h.hLog.Printf("Request failed: %v", err)
		return resp, err
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.hLog.Printf("Unable to read response: %v", err)
		return resp, err
	}

	resp.Response = r
	resp.Body = body
	if h.debug {
		h.hLog.Printf("%s %s -- %d %v", method, req.URL.RequestURI(), resp.Response.StatusCode, req.Header)
	}

	return resp, nil
}

// DoEnvelope makes an HTTP request and parses the JSON response (fastglue envelop structure)
func (h *httpClient) DoEnvelope(method, url string, params map[string]interface{}, headers http.Header, accessToken string, obj interface{}) error {
	resp, err := h.Do(method, url, params, headers, accessToken)
	if err != nil {
		return err
	}

	// fmt.Printf("\nResponse : %v - %v\n", resp.Response.Status, resp.Response.Request.URL)

	// fmt.Printf("\nResponse : %v \n", string(resp.Body))

	// Successful request, but error envelope.
	if resp.Response.StatusCode >= http.StatusBadRequest {
		var e envelope
		if err := json.Unmarshal(resp.Body, &e); err != nil {
			h.hLog.Printf("Error parsing JSON response: %v", err)
			return err
		}

		return NewError(e.ErrorCode, e.Message, e.Data)
	}

	// We now unmarshal the body.
	envl := envelope{}
	envl.Data = obj

	if err := json.Unmarshal(resp.Body, &obj); err != nil {
		h.hLog.Printf("Error parsing JSON response: %v | %s", err, resp.Body)
		return err
	}

	return nil
}

// GetClient return's the underlying net/http client.
func (h *httpClient) GetClient() *httpClient {
	return h
}
