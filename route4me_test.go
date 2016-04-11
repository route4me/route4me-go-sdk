package route4me

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const apiKey = "11111111111111111111111111111111"

func TestNewClient(t *testing.T) {
	cli := NewClient(apiKey)
	if cli.Client == nil {
		t.Error("Client has not been initialized.")
	}
	if cli.APIKey != apiKey {
		t.Error("APIKey has not been assigned.")
	}
}

func testClient(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		read, err := ioutil.ReadAll(r.Body)
		if err != nil || len(read) == 0 {
			fmt.Fprint(w, body)
		} else {
			w.Write(read)
		}
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	client := &Client{Client: httpClient, APIKey: apiKey, BaseURL: "http://example.com"}
	return server, client
}

type response struct {
	Number int `json:"json"`
}

func TestDecodingGet(t *testing.T) {
	server, cli := testClient(200, `{"json":42}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodGet, "/whatever/", &struct{}{}, resp)
	if err != nil {
		t.Error(err)
	}
	if resp.Number != 42 {
		t.Error("Unmarshalling went wrong")
	}
}

func TestDecodingPost(t *testing.T) {
	server, cli := testClient(200, `{"json":42}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodPost, "/whatever/", &struct {
		Number int `json:"json"`
	}{Number: 152}, resp)
	if err != nil {
		t.Error(err)
	}
	if resp.Number != 152 {
		t.Error("Error occured during unmarshalling")
	}
}

func TestDecodingErrors(t *testing.T) {
	server, cli := testClient(500, `{"errors":["error#1","error#2"]}`)
	defer server.Close()
	resp := &response{}
	err := cli.Do(http.MethodGet, "/", &struct{}{}, resp)
	if err == nil || err.Error() != "error#1,error#2" {
		t.Error("Expecting error 'error#1,error#2', got: ", err.Error())
	}
}
