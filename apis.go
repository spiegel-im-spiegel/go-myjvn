package myjvn

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/go-myjvn/request"
	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"github.com/spiegel-im-spiegel/go-myjvn/status"
)

//VulnOverviewListXML calls a MyJVN RESTful API: "getVulnOverviewList", and returns XML data
func VulnOverviewListXML() ([]byte, error) {
	values := url.Values{
		"method": {"getVulnOverviewList"},
		"feed":   {"hnd"},
		"lang":   {"ja"},
	}
	return request.API(values)
}

//VulnOverviewList calls a MyJVN RESTful API: "getVulnOverviewList", and returns JVNRSS value
func VulnOverviewList() (*rss.JVNRSS, error) {
	data, err := VulnOverviewListXML()
	if err != nil {
		return nil, err
	}
	stat, err := status.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if stat.Status.ReturnCode != 0 {
		return nil, errors.New(stat.Status.ErrorMsg)
	}
	return rss.Unmarshal(data)
}

//VulnDetailInfoXML calls a MyJVN RESTful API: "getVulnDetailInfo", and returns XML data
func VulnDetailInfoXML(vulnID []string) ([]byte, error) {
	if len(vulnID) == 0 {
		return nil, os.ErrInvalid
	}

	values := url.Values{
		"method": {"getVulnDetailInfo"},
		"feed":   {"hnd"},
		"lang":   {"ja"},
		"vulnId": {strings.Join(vulnID, "+")},
	}
	return request.API(values)
}

//VulnDetailInfo calls a MyJVN RESTful API: "getVulnDetailInfo", and returns VULDEF value
func VulnDetailInfo(vulnID []string) (*rss.JVNRSS, error) {
	data, err := VulnDetailInfoXML(vulnID)
	if err != nil {
		return nil, err
	}
	stat, err := status.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if stat.Status.ReturnCode != 0 {
		return nil, errors.New(stat.Status.ErrorMsg)
	}
	return rss.Unmarshal(data)
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
