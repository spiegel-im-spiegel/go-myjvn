package myjvn

import (
	"net/url"

	"github.com/spiegel-im-spiegel/go-myjvn/request"
)

func VulnOverviewListXml() ([]byte, error) {
	values := url.Values{
		"method": {"getVulnOverviewList"},
		"feed":   {"hnd"},
		"lang":   {"ja"},
	}
	return request.Api(values)
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
