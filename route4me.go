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

func (c *Client) constructBody(data interface{}) (contentType string, reader bytes.Buffer, err error) {
	//Check if the data struct has any postform data to pass to the body
	params := utils.StructToURLValues("form", data)
	// if there are no form parameters, it's likely the request is a json
	if len(params) == 0 {
		if err = json.NewEncoder(&reader).Encode(data); err != nil {
			return
		}
		contentType = "application/json"
		return
	}

	//otherwise we encode the form as a multipart form
	w := multipart.NewWriter(&reader)
	defer w.Close()
	for key, vals := range params {
		for _, v := range vals {
			err = w.WriteField(key, v)
			if err != nil {
				return
			}
		}
	}
	contentType = w.FormDataContentType()
	return
}

func (c *Client) DoNoDecode(method string, endpoint string, data interface{}) (response []byte, err error) {
	var requestBody bytes.Buffer
	var contentType string
	//We might change this to == for better accuracy
	if method != http.MethodGet && method != http.MethodOptions {
		if contentType, requestBody, err = c.constructBody(data); err != nil {
			return response, err
		}
	}

	request, err := http.NewRequest(method, c.BaseURL+endpoint, &requestBody)
	if err != nil {
		return response, err
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
		return response, err
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return response, InvalidStatusCode
	}
	return response, err
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
