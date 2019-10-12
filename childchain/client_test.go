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
	"net/http"
	"testing"
)

//should create a new client with watcher endpoint
func TestClient(t *testing.T) {
	var err error
	var res *Client
	watcherAddr := "http://watcher.samrong.omg.network/"
	_, err = NewClient("", &http.Client{})
	if err == nil {
		t.Errorf("expected %v, got %v", "invalid URL error", err)
	}

	_, err = NewClient("watcher.samrong.omg.network", &http.Client{})
	if err == nil {
		t.Errorf("expected %v, got %v", "invalid URL scheme error", nil)
	}

	_, err = NewClient(watcherAddr, nil)
	if err == nil {
		t.Errorf("expected %v, got %v", "no http.client error", nil)
	}

	res, err = NewClient(watcherAddr, &http.Client{})
	if err != nil {
		t.Errorf("expected %v, got %v", "no error", err)
	}

	if res.GetWatcher() != watcherAddr {
		t.Errorf("expected %v, got %v", watcherAddr, res.GetWatcher())
	}
}
