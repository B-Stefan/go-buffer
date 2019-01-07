package api

import (
	"encoding/json"
	"github.com/go-playground/form"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	newRequest(method, path string, body interface{}) (*http.Request, error)
	do(req *http.Request, v interface{}) (*http.Response, error)
}

type apiClientV1 struct {
	httpClient *http.Client
	BaseURL    *url.URL
	UserAgent  string
	Profile    *ProfileService
}

func NewClient(httpClient *http.Client) *apiClientV1 {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseUrl, _ := url.Parse("https://api.bufferapp.com/1")
	c := &apiClientV1{httpClient: httpClient, BaseURL: baseUrl, UserAgent: "golang-wrapper"}
	c.Profile = &ProfileService{client: c}

	return c
}

func (c *apiClientV1) encodeValues(body interface{}) (io.Reader, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(&body)
	if err != nil {
		return nil, err
	}
	buf := strings.NewReader(values.Encode())

	return buf, err
}

func (c *apiClientV1) newRequest(method, path string, body interface{}) (*http.Request, error) {

	rel := &url.URL{Path: c.BaseURL.Path + path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.Reader
	var err error
	if body != nil {
		buf, err = c.encodeValues(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}
func (c *apiClientV1) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
