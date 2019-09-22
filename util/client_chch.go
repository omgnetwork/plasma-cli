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

package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ChChRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

func MakeChChReq(
	chchRequest ChChRequester,
	chchClient string,
	method string,
	postData interface{},
) ([]byte, error) {

	//build request
	var url strings.Builder
	url.WriteString(chchClient)
	url.WriteString(method)
	js, _ := json.Marshal(postData)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	// Make the request
	resp, err := chchRequest.Do(r)
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
