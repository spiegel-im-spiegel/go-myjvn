package myjvn

import (
	"net/url"

	"github.com/spiegel-im-spiegel/go-myjvn/request"
	"github.com/spiegel-im-spiegel/go-myjvn/rss"
)

func VulnOverviewListXml() ([]byte, error) {
	values := url.Values{
		"method": {"getVulnOverviewList"},
		"feed":   {"hnd"},
		"lang":   {"ja"},
	}
	return request.Api(values)
}

func VulnOverviewList() (*rss.JVNRSS, error) {
	data, err := VulnOverviewListXml()
	if err != nil {
		return nil, err
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
