package kannel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/cgcorea/ksidekick/internal/debug"
)

type Client struct {
	url      string
	username string
	password string
	client   *http.Client
}

type Options struct {
	Charset  string
	UDH      string
	SMSC     string
	MClass   string
	MWI      string
	Compress string
	Coding   string
	Validity string
	Deferred string
	DLRMask  string
	DLRURL   string
	Account  string
	PID      string
	AltDCS   string
	BInfo    string
	RPI      string
	Priority string
	Metadata string
}

type Response struct {
	Code    int
	Message string
}

func setHeader(request *http.Request, header, value string) {
	if value != "" {
		request.Header.Set(header, value)
	}
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

//
// SMSPush equivalents to HTTP headers from Kannel documentation.
//
// SMSPush eq	X-Kannel Header
// -----------------------------
// username		X-Kannel-Username
// password		X-Kannel-Password
// from			X-Kannel-From
// to			X-Kannel-To
// text			request body
// charset		charset as in Content-Type: text/html; charset=ISO-8859-1
// udh			X-Kannel-UDH
// smsc			X-Kannel-SMSC
// flash		X-Kannel-Flash (deprecated, see X-Kannel-MClass
// mclass		X-Kannel-MClass
// mwi			X-Kannel-MWI
// compress		X-Kannel-Compress
// coding		X-Kannel-Coding. If unset, defaults to 0 (7 bits) if Content-Type is text/plain,
//								 text/html or text/vnd.wap.wml. On application/octet-stream, defaults to 8 bits (1).
//								 All other Content-Type values are rejected.
// validity		X-Kannel-Validity
// deferred		X-Kannel-Deferred
// dlr-mask		X-Kannel-DLR-Mask
// dlr-url		X-Kannel-DLR-Url
// account		X-Kannel-Account
// pid			X-Kannel-PID
// alt-dcs		X-Kannel-Alt-DCS
// binfo		X-Kannel-BInfo
// rpi			X-Kannel-RPI
// priority		X-Kannel-Priority
// meta-data	X-Kannel-Meta-Data
//
// TODO(cgcorea): Set charset in Content-Type
func (k *Client) buildHTTPRequest(from, to, text string, options *Options) (*http.Request, error) {
	req, err := http.NewRequest("POST", k.url, strings.NewReader(text))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Kannel-Username", k.username)
	req.Header.Set("X-Kannel-Password", k.password)
	req.Header.Set("X-Kannel-From", from)
	req.Header.Set("X-Kannel-To", to)

	if options != nil {
		setHeader(req, "X-Kannel-UDH", options.UDH)
		setHeader(req, "X-Kannel-SMSC", options.SMSC)
		setHeader(req, "X-Kannel-MClass", options.MClass)
		setHeader(req, "X-Kannel-MWI", options.MWI)
		setHeader(req, "X-Kannel-Compress", options.Compress)
		setHeader(req, "X-Kannel-Coding", options.Coding)
		setHeader(req, "X-Kannel-Validity", options.Validity)
		setHeader(req, "X-Kannel-Deferred", options.Deferred)
		setHeader(req, "X-Kannel-DLR-Mask", options.DLRMask)
		setHeader(req, "X-Kannel-DLR-Url", options.DLRURL)
		setHeader(req, "X-Kannel-Account", options.Account)
		setHeader(req, "X-Kannel-PID", options.PID)
		setHeader(req, "X-Kannel-Alt-DCS", options.AltDCS)
		setHeader(req, "X-Kannel-BInfo", options.BInfo)
		setHeader(req, "X-Kannel-RPI", options.RPI)
		setHeader(req, "X-Kannel-Priority", options.Priority)
		setHeader(req, "X-Kannel-Meta-Data", options.Metadata)
	}

	debug.Inspect(req.Header, os.Stdout)

	return req, nil
}

// Send ...
func (k *Client) Send(from, to, text string, options *Options) (*Response, error) {
	req, err := k.buildHTTPRequest(from, to, text, options)
	if err != nil {
		return nil, err
	}
	resp, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Code:    resp.StatusCode,
		Message: string(body),
	}

	return response, nil
}
