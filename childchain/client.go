// Copyright 2019 OmiseGO Pte Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package childchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Child chain client
type Client struct {
	Watcher    *url.URL
	HttpClient *http.Client
}

type ClientError struct {
	Version string `json:"version"`
	Success bool   `json:"success"`
	Data    struct {
		Object      string `json:"object"`
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"data"`
}

// Creates new instance of child chain client with a watcher endpoint
func NewClient(watcher string, httpclient *http.Client) (*Client, error) {
	if watcher == "" {
		return nil, fmt.Errorf("Error parsing watcher: %v", watcher)
	}
	if httpclient == nil {
		return nil, fmt.Errorf("Error creating new client, http.Client is: %v", nil)
	}
	w, err := url.ParseRequestURI(watcher)
	if err != nil {
		return nil, fmt.Errorf("Error parsing watcher: %v", err)
	}
	return &Client{w, httpclient}, nil
}

// returns the watcher URL string
func (c *Client) GetWatcher() string {
	return c.Watcher.String()
}

func (c *Client) do(method string, postData interface{}) ([]byte, error) {
	//build request
	var url strings.Builder
	url.WriteString(c.GetWatcher())
	url.WriteString(method)
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	resp, err := c.HttpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rstring, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rstring, err
}
