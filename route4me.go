package route4me

import (
        "net/http"
        "net/url"
        "bytes"
        "io"
        "io/ioutil"
        "encoding/json"
)

type Route4Me struct {
        Key string
        BaseUrl string
}

func NewClient(key string) *Route4Me {
        //baseUrl := "http://www.route4me.com"
        baseUrl := "http://staging.route4me.com:8080"
        return &Route4Me{Key: key, BaseUrl: baseUrl}
}


func processResponse(response *http.Response, err error, container interface{}) (interface{}, *Exception, error) {
        var exception *Exception


        if err != nil {
                return container, exception, err
        }

        defer response.Body.Close()

        responseBody, err := ioutil.ReadAll(response.Body)
        if err != nil {
                return container, exception, err
        }

        if response.StatusCode != http.StatusOK {
                exception = new(Exception)
                err = json.Unmarshal(responseBody, exception)

                return container, exception, err
        }

        err = json.Unmarshal(responseBody, container)

        return container, exception, err
}

func (r4m *Route4Me) Get(endpoint string, requestParams url.Values) (*http.Response, error) {
        return r4m.request("GET", endpoint, requestParams, nil)
}


func (r4m *Route4Me) Post(endpoint string, requestParams url.Values, body []byte) (*http.Response, error) {
        return r4m.request("POST", endpoint, requestParams, bytes.NewBuffer(body))
}

func (r4m *Route4Me) Put(endpoint string, requestParams url.Values, body []byte) (*http.Response, error) {
        return r4m.request("PUT", endpoint, requestParams, bytes.NewBuffer(body))
}


func (r4m *Route4Me) request(requestType string, endpoint string, requestParams url.Values, body io.Reader) (*http.Response, error) {
        url := r4m.BaseUrl + endpoint + "?" + requestParams.Encode()

        req, err := http.NewRequest(requestType, url, body)
        if err != nil {
                return nil, err
        }
        req.Header.Set("X-R4M-API-Key", r4m.Key)

        client := &http.Client{}
        return client.Do(req)
}
