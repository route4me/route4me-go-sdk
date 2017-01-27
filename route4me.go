package route4me

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	defaultTimeout time.Duration = time.Minute * 30
	BaseURL                      = "https://www.route4me.com"
)

var InvalidStatusCode = errors.New("Invalid status code")

type Client struct {
	APIKey  string
	Client  *http.Client
	BaseURL string
}

func NewClientWithOptions(APIKey string, timeout time.Duration, baseURL string) *Client {
	return &Client{
		APIKey:  APIKey,
		Client:  &http.Client{Timeout: timeout},
		BaseURL: baseURL,
	}
}

// NewClient creates a route4me client
func NewClient(APIKey string) *Client {
	return NewClientWithOptions(APIKey, defaultTimeout, BaseURL)
}

func (c *Client) DoNoDecode(method string, endpoint string, data interface{}) ([]byte, error) {
	var reader bytes.Buffer
	var byt []byte
	contentType := "application/json"
	//We might change this to == for better accuracy
	if method != http.MethodGet && method != http.MethodOptions {
		//Check if the data struct has any postform data to pass to the body
		params := utils.StructToURLValues("form", data)
		if len(params) > 0 {
			w := multipart.NewWriter(&reader)
			for key, vals := range params {
				for _, v := range vals {
					err := w.WriteField(key, v)
					if err != nil {
						return byt, err
					}
				}
			}
			w.Close()
			contentType = w.FormDataContentType()
		} else {
			serialized, err := json.Marshal(data)
			if err != nil {
				return byt, err
			}
			reader.Write(serialized)
		}
	}

	request, err := http.NewRequest(method, c.BaseURL+endpoint, &reader)
	if err != nil {
		return byt, err
	}

	request.Header.Set("Content-Type", contentType)
	params := url.Values{}

	if data != nil {
		//Prepare query string
		params = utils.StructToURLValues("http", data)
	}
	params.Add("api_key", c.APIKey)
	request.URL.RawQuery = params.Encode()

	//b, err := httputil.DumpRequestOut(request, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(b))

	resp, err := c.Client.Do(request)

	// b, err = httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(string(b))
	if err != nil {
		return byt, err
	}
	defer resp.Body.Close()
	byt, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return byt, InvalidStatusCode
	}
	return byt, err
}

func (c *Client) Do(method string, endpoint string, data interface{}, out interface{}) error {
	read, err := c.DoNoDecode(method, endpoint, data)
	//Error handling is a bit weird
	if err != nil {
		//Check if invalid status code - errors:[] response is returned only when statuscode is not 200
		if err == InvalidStatusCode {
			errs := &ErrorResponse{}
			//Try to parse to ErrorResponse
			unmerr := json.Unmarshal(read, errs)
			//Sometimes (status code: 500,404) errors:[] might not be returned, we return the err from the request when it happens
			if unmerr != nil {
				return err
			}
			//Join all errors in the ErrorResponse
			return errors.New(strings.Join(errs.Errors, ","))
		}
		return err
	}
	if out == nil {
		return err
	}
	err = json.Unmarshal(read, out)
	return err
}
