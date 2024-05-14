package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client is gosplash client
type Client struct {
	clientId string
	client   *http.Client
	proxyURL string
}

type ClientOption func(c *Client)

// SetClientProxy can set proxy url when NewClient
func SetClientProxy(proxyURL string) ClientOption {
	return func(c *Client) {
		c.proxyURL = proxyURL
	}
}

// OmitResponse define response headers and  other info
type OmitResponse struct {
	RateLimitLimit     int `json:"X-Ratelimit-Limit"`
	RateLimitRemaining int `json:"X-Ratelimit-Remaining"`
	// original response body
	OriginResponseBody string
	// request URL
	requestURL string
}

// Req request params interface
type Req interface {
	// API request url
	API() string
	// Params request query
	Params() map[string]string
}

// NewClient when you register unsplash developer, you can get client id (access key)
// register address: https://unsplash.com/developers
func NewClient(clientId string, opt ...ClientOption) *Client {
	c := &Client{}
	for _, o := range opt {
		o(c)
	}
	c.clientId = clientId
	c.client = http.DefaultClient
	var proxyUrl *url.URL
	if c.proxyURL != "" {
		proxyUrl, _ = url.Parse(c.proxyURL)
	}
	c.client.Transport = &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   50,
		Proxy:                 http.ProxyURL(proxyUrl),
	}
	return c
}

func (c *Client) request(method string, req Req, reply interface{}, headers map[string]string, urlPath ...string) (response *OmitResponse, err error) {
	var buf io.Reader
	var values url.Values
	header := make(map[string]string)

	// add headers by content-type
	header["Content-Type"] = "application/json"
	header["Accept-Version"] = "v1"
	if headers != nil {
		for k, v := range headers {
			header[k] = v
		}
	}

	// add default Authorization Client id
	if _, ok := header["Authorization"]; !ok {
		header["Authorization"] = fmt.Sprintf("Client-ID %s", c.clientId)
	}

	// create request URL
	reqURL := req.API()
	if urlPath != nil {
		reqURL, _ = url.JoinPath(reqURL, urlPath...)
	}

	// create body or query
	if method == "GET" {
		values, _ = query.Values(req)
	} else {
		m, _ := json.Marshal(req)
		buf = bytes.NewReader(m)
	}

	// add url params
	if req.Params() != nil {
		for k, v := range req.Params() {
			values.Add(k, v)
		}
	}

	reqURL = reqURL + "?" + values.Encode()
	request, err := http.NewRequest(method, reqURL, buf)
	if err != nil {
		return nil, err
	}

	// set headers
	for k, v := range header {
		request.Header.Set(k, v)
	}

	// do request
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// create response headers
	response = c.parseOmitResponseResponse(resp.Header)

	// add origin response body
	response.OriginResponseBody = string(data)

	// add request URL
	response.requestURL = reqURL

	// check response's status
	if resp.StatusCode != 200 {
		return response, errors.New(string(data))
	}

	// create response
	if reply != nil {
		err = json.Unmarshal(data, reply)
		if err != nil {
			return response, err
		}
	}
	return
}

func (c *Client) parseOmitResponseResponse(header http.Header) *OmitResponse {
	response := &OmitResponse{}
	if header.Get("X-Ratelimit-Limit") != "" {
		rateLimit, _ := strconv.Atoi(header.Get("X-Ratelimit-Limit"))
		response.RateLimitLimit = rateLimit
	}

	if header.Get("X-Ratelimit-Remaining") != "" {
		rateLimit, _ := strconv.Atoi(header.Get("X-Ratelimit-Remaining"))
		response.RateLimitRemaining = rateLimit
	}
	return response
}
