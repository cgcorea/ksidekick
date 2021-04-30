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

type Request struct {
	*http.Request
}

type Response struct {
	Code    int
	Message string
}

func NewClient(host string, port int, username, password string) *Client {
	url := fmt.Sprintf("http://%v:%v/cgi-bin/sendsms", host, port)

	client := &Client{
		url:      url,
		username: username,
		password: password,
		client:   &http.Client{},
	}

	return client
}

// NewRequest ...
func (k *Client) NewRequest(from, to, text string, options ...func(*Request)) (request *Request, err error) {
	req, err := http.NewRequest("POST", k.url, strings.NewReader(text))
	if err != nil {
		return nil, err
	}

	request = &Request{Request: req}
	request.Set(Username(k.username))
	request.Set(Password(k.password))
	request.Set(From(from))
	request.Set(To(to))

	for _, option := range options {
		request.Set(option)
	}

	return
}

// Send ...
func (k *Client) Send(request *Request) (*Response, error) {
	resp, err := k.client.Do(request.Request)
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
