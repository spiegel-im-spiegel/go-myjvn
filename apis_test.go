package myjvn

import (
	"testing"

	"github.com/spiegel-im-spiegel/go-myjvn/rss/option"
)

func TestVulnOverviewListXml(t *testing.T) {
	_, err := New().VulnOverviewListXML(option.New())
	if err != nil {
		t.Errorf("VulnOverviewListXml() = \"%v\", want nil.", err)
	}
}

func TestVulnOverviewList(t *testing.T) {
	_, err := New().VulnOverviewList(option.New())
	if err != nil {
		t.Errorf("VulnOverviewList() = \"%v\", want nil.", err)
	}
}

func TestVulnDetailInfoXML(t *testing.T) {
	vulnID := []string{"JVNDB-2018-000022", "JVNDB-2018-000024"}
	_, err := New().VulnDetailInfoXML(vulnID)
	if err != nil {
		t.Errorf("VulnOverviewListXml() = \"%v\", want nil.", err)
	}
}

func TestVulnDetailInfo(t *testing.T) {
	vulnID := []string{"JVNDB-2018-000022", "JVNDB-2018-000024"}
	_, err := New().VulnDetailInfo(vulnID)
	if err != nil {
		t.Errorf("VulnOverviewList() = \"%v\", want nil.", err)
	}
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
