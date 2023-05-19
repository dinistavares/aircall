package aircall

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
  apiBaseURL        = "https://api.aircall.io/"
  apiDefaultVersion = "v1"
)

// A Client struct is holding developer app information
// to perform API requests
type Client struct {
  AppID    string
  AppToken string
  AccessToken string
  Version  string
}

// NewClient creates a new Client struct using the default API version
//
// client := aircall.NewClient(appId, appSecret)
// res, err := client.Ping()
// fmt.Println(res.Ping) // "pong"
func NewClient(id, token string) *Client {
  return &Client{
    AppID:    id,
    AppToken: token,
    Version:  apiDefaultVersion,
  }
}

// NewClient creates a new Client struct using the default API version
//
// client := aircall.NewClient(accessToken)
// res, err := client.Ping()
// fmt.Println(res.Ping) // "pong"
func NewClientWithAccessToken(token string) *Client {
  return &Client{
    AccessToken: token,
    Version:  apiDefaultVersion,
  }
}

// Get is a shorthand of Request("GET", path, params, Request{})
func (client *Client) Get(path string, params map[string]string) ([]byte, error) {
  return client.Request("GET", path, params, nil)
}

// Post is a shorthand of Request("POST", path, params, Request{})
func (client *Client) Post(path string, request interface{}) ([]byte, error) {
  return client.Request("POST", path, map[string]string{}, request)
}

// Delete is a shorthand of Request("DELETE", path, params, Request{})
func (client *Client) Delete(path string, params map[string]string) ([]byte, error) {
  return client.Request("DELETE", path, params, nil)
}

// Request sends a HTTP request to the API
func (client *Client) Request(method string, path string, params map[string]string, body interface{}) ([]byte, error) {
  url := buildURL(client, path, params)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)

    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, url, buf)

  if err != nil {
    return nil, err
  }

  req.Header.Set("Content-Type", "application/json")

  if (client.AccessToken != "") {
    req.Header.Add("Authorization", "Bearer "+ client.AccessToken)
  } else {
    req.Header.Add("Authorization", "Basic "+basicAuthHeader(client.AppID, client.AppToken))
  }

  c := &http.Client{}
  res, err := c.Do(req)

  if err != nil {
    return nil, err
  }

  if res.StatusCode >= 400 {
    return nil, errors.New("Client error: " + res.Status)
  } else if res.StatusCode >= 500 {
    return nil, errors.New("Server error: " + res.Status)
  }

  defer res.Body.Close()

  return ioutil.ReadAll(res.Body)
}

func basicAuthHeader(username, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}

func buildURL(client *Client, path string, params map[string]string) string {
  var urlBuf bytes.Buffer
  urlBuf.WriteString(apiBaseURL)
  urlBuf.WriteString(client.Version)
  urlBuf.WriteString(path)

  values := url.Values{}
  for k, v := range params {
    values.Add(k, v)
  }
  urlBuf.WriteString("?" + values.Encode())

  return urlBuf.String()
}
