package api

import (
	"bytes"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoRequest(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		// equals(t, req.URL.String(), "http://example.com/some/path")
		println(req.Host)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"status": "ok"}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	apiClient := NewClient(client)
	req, err := apiClient.newRequest("GET", "/myPath", nil)

	println(req, err)
	resBody := struct {
		Status string `json:"status"`
	}{}
	res, err := apiClient.do(req, &resBody)

	if err != nil {
		log.Fatal(err)
	}

	println(res, err)
	assert.Equal(t, "ok", resBody.Status)
	assert.Equal(t, 200, res.StatusCode)
}

func TestDefaultBaseUrl(t *testing.T) {

	apiClient := NewClient(nil)

	assert.Equal(t, "https://api.bufferapp.com/1", apiClient.BaseURL.String())

}

func TestEncodeValues(t *testing.T) {

	apiClient := NewClient(nil)

	type Nested struct {
		Name string `form:"name"`
	}
	body := struct {
		Test string   `json:"test",form:"active"`
		List []Nested `json:"test"`
	}{
		Test: "Bob",
		List: []Nested{{Name: "Yoda"}, {Name: "Luke"}},
	}

	reader, err := apiClient.encodeValues(body)

	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)

	if err != nil {
		t.Fatal(err)
	}

	newStr := buf.String()
	assert.Equal(t, "List%5B0%5D.name=Yoda&List%5B1%5D.name=Luke&Test=Bob", newStr)

}

func TestContentType(t *testing.T) {

	apiClient := NewClient(nil)
	req, err := apiClient.newRequest("GET", "/profiles.json", nil)

	if err != nil {
		log.Fatal(err)
	}

	contentType := req.Header.Get("Content-Type")
	assert.Equal(t, "application/x-www-form-urlencoded", contentType)
}

func TestMethod(t *testing.T) {

	apiClient := NewClient(nil)
	req, err := apiClient.newRequest("GET", "/profiles.json", nil)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "GET", req.Method)
}

func TestPath(t *testing.T) {

	baseUrl, err := url.Parse("https://mydomain.com")
	apiClient := NewClient(nil)
	apiClient.BaseURL = baseUrl
	req, err := apiClient.newRequest("GET", "/profiles.json", nil)

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "/profiles.json", req.URL.Path)
}

func TestPathWithDefaultUrl(t *testing.T) {

	baseUrl, err := url.Parse("https://mydomain.com/v1")
	apiClient := NewClient(nil)
	apiClient.BaseURL = baseUrl
	req, err := apiClient.newRequest("GET", "/profiles.json", nil)

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "https://mydomain.com/v1/profiles.json", req.URL.String())
}
