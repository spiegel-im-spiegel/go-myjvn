package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	pathMyJVN = "https://jvndb.jvn.jp/myjvn" //path for MyJVN RESTful API
)

//API calls a MyJVN RESTful API.
func API(v url.Values) ([]byte, error) {
	resp, err := new(http.Client).Get(pathMyJVN + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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
