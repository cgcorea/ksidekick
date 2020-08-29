package kannel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	url      string
	username string
	password string
	client   *http.Client
}

type Response struct {
	Code    int
	Message string
}

func NewClient(host string, port int, username string, password string) *Client {
	url := fmt.Sprintf("http://%v:%v/cgi-bin/sendsms", host, port)

	client := &Client{
		url:      url,
		username: username,
		password: password,
		client:   &http.Client{},
	}

	return client
}

func (k *Client) SendText(from string, to string, text string) (*Response, error) {
	req, err := http.NewRequest("POST", k.url, strings.NewReader(text))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Kannel-Username", k.username)
	req.Header.Add("X-Kannel-Password", k.password)
	req.Header.Add("X-Kannel-From", from)
	req.Header.Add("X-Kannel-To", to)

	resp, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{Code: resp.StatusCode, Message: string(body)}, nil
}
