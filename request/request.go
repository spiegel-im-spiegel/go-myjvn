package request

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Client is structural subtype for MyJVN API Client
type Client interface {
	Request(url.Values) ([]byte, error)
}

//MyJVNClient is http.Client for MyJVN RESTful API
type MyJVNClient struct {
	client *http.Client
	path   string
}

//New creates MyJVNClient instance
func New() Client {
	return &MyJVNClient{client: &http.Client{}, path: "https://jvndb.jvn.jp/myjvn"}
}

//Request calls a MyJVN RESTful API.
func (c *MyJVNClient) Request(v url.Values) ([]byte, error) {
	resp, err := c.client.Get(c.path + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = checkStatus(resp); err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func checkStatus(resp *http.Response) error {
	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return errors.New(resp.Status)
	}
	return nil
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
