package route4me

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout time.Duration = time.Minute * 30
	baseURL                      = "https://www.route4me.com"
)

type Client struct {
	APIKey string
	Client *http.Client
}

func NewClientWithTimeout(APIKey string, timeout time.Duration) *Client {
	return &Client{
		APIKey: APIKey,
		Client: &http.Client{Timeout: timeout},
	}
}

// NewClient creates a route4me client
func NewClient(APIKey string) *Client {
	return NewClientWithTimeout(APIKey, defaultTimeout)
}

// func (c *Client) NewRequest(method string, url string) *http.Request {
// 	request := http.NewRequest(method, url, nil)
// 	request.
// }

func (c *Client) DoNoDecode(method string, endpoint string, data interface{}) ([]byte, error) {
	var reader io.Reader
	var byt []byte
	//We might change this to == for better accuracy
	if method != http.MethodGet && method != http.MethodOptions {
		serialized, err := json.Marshal(data)
		if err != nil {
			return byt, err
		}
		reader = bytes.NewReader(serialized)
	}

	request, err := http.NewRequest(http.MethodPost, baseURL+endpoint, reader)
	if err != nil {
		return byt, err
	}
	//Prepare query string
	params := structToMap(data)
	params.Add("api_key", c.APIKey)
	request.URL.RawQuery = params.Encode()
	// params := request.URL.Query()
	// params.Add("api_key", c.APIKey)
	// request.URL.RawQuery = params.Encode()
	resp, err := c.Client.Do(request)
	if err != nil {
		return byt, err
	}
	defer resp.Body.Close()
	byt, err = ioutil.ReadAll(resp.Body)
	return byt, err
}

func (c *Client) Do(method string, endpoint string, data interface{}, out interface{}) error {
	read, err := c.DoNoDecode(method, endpoint, data)
	err = json.Unmarshal(read, out)
	return err
}
