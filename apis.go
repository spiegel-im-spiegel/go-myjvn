package myjvn

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/spiegel-im-spiegel/go-myjvn/option"
	"github.com/spiegel-im-spiegel/go-myjvn/request"
	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"github.com/spiegel-im-spiegel/go-myjvn/status"
	"github.com/spiegel-im-spiegel/go-myjvn/vuldef"
)

//APIs implements MyJVN RESTful API
type APIs struct {
	client request.Client
}

//New creates APIs instance
func New() *APIs {
	return &APIs{client: request.New()}
}

//VulnOverviewListXML calls a MyJVN RESTful API: "getVulnOverviewList", and returns XML data
func (api *APIs) VulnOverviewListXML(opt *option.Option) ([]byte, error) {
	values := url.Values{
		"method":       {"getVulnOverviewList"},
		"feed":         {"hnd"},
		"lang":         {"ja"},
		"maxCountItem": {strconv.Itoa(rss.MaxItems)},
	}
	opt.AddQuery(values)
	return api.client.Request(values)
}

//VulnOverviewList calls a MyJVN RESTful API: "getVulnOverviewList", and returns JVNRSS value
func (api *APIs) VulnOverviewList(opt *option.Option) (*rss.JVNRSS, error) {
	data, err := api.VulnOverviewListXML(opt)
	if err != nil {
		return nil, err
	}
	stat, err := status.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if stat.Status.ReturnCode != 0 {
		return nil, fmt.Errorf("Code %s: %s", stat.Status.ErrorCode, stat.Status.ErrorMsg)
	}
	return rss.Unmarshal(data)
}

//VulnDetailInfoXML calls a MyJVN RESTful API: "getVulnDetailInfo", and returns XML data
func (api *APIs) VulnDetailInfoXML(vulnID []string) ([]byte, error) {
	if len(vulnID) <= 0 || len(vulnID) > vuldef.MaxItems {
		return nil, os.ErrInvalid
	}

	values := url.Values{
		"method":       {"getVulnDetailInfo"},
		"feed":         {"hnd"},
		"lang":         {"ja"},
		"startItem":    {"1"},
		"maxCountItem": {strconv.Itoa(len(vulnID))},
		"vulnId":       {strings.Join(vulnID, "+")},
	}
	return api.client.Request(values)
}

//VulnDetailInfo calls a MyJVN RESTful API: "getVulnDetailInfo", and returns VULDEF value
func (api *APIs) VulnDetailInfo(vulnID []string) (*vuldef.VULDEF, error) {
	data, err := api.VulnDetailInfoXML(vulnID)
	if err != nil {
		return nil, err
	}
	stat, err := status.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if stat.Status.ReturnCode != 0 {
		return nil, fmt.Errorf("Code %s: %s", stat.Status.ErrorCode, stat.Status.ErrorMsg)
	}
	return vuldef.Unmarshal(data)
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
