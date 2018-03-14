package vuldef

import (
	"encoding/json"
	"encoding/xml"
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/go-myjvn/values"
)

const (
	MaxItems = 10
)

//VULDEF is structure of VULDEF
type VULDEF struct {
	Vulinfo []struct {
		VulinfoID   string `xml:"VulinfoID" json:"VulinfoID"`
		VulinfoData struct {
			Title    string `xml:"Title" json:"Title"`
			Overview string `xml:"VulinfoDescription>Overview" json:"Overview"`
			Affected []struct {
				Name        string `xml:"Name" json:"Name"`
				ProductName string `xml:"ProductName" json:"ProductName"`
				Cpe         struct {
					Version string `xml:"version,attr" json:"version"`
					Value   string `xml:",chardata" json:"value"`
				} `xml:"Cpe" json:"Cpe"`
				VersionNumber []string `xml:"VersionNumber" json:"VersionNumber"`
			} `xml:"Affected>AffectedItem" json:"Affected"`
			Impact struct {
				Description string `xml:"ImpactItem>Description" json:"Description"`
				CVSS        []struct {
					Version    string `xml:"version,attr" json:"Version"`
					BaseVector string `xml:"Vector" json:"BaseVector"`
					BaseScore  string `xml:"Base" json:"BaseScore"`
					Severity   string `xml:"Severity" json:"Severity"`
				} `xml:"Cvss" json:"Cvss"`
			} `xml:"Impact" json:"Impact"`
			Solution struct {
				Description string `xml:"SolutionItem>Description" json:"Description"`
			} `xml:"Solution" json:"Solution"`
			Related []struct {
				Type      string `xml:"type,attr" json:"Type"`
				Name      string `xml:"Name" json:"Name"`
				VulinfoID string `xml:"VulinfoID" json:"VulinfoID"`
				Title     string `xml:"Title,omitempty" json:"Title,omitempty"`
				URL       string `xml:"URL" json:"URL"`
			} `xml:"Related>RelatedItem" json:"Related"`
			History []struct {
				HistoryNo   int         `xml:"HistoryNo" json:"HistoryNo"`
				DateTime    values.Time `xml:"DateTime" json:"DateTime"`
				Description string      `xml:"Description" json:"Description"`
			} `xml:"History>HistoryItem" json:"History"`
			DateFirstPublished values.Time `xml:"DateFirstPublished" json:"DateFirstPublished"`
			DateLastUpdated    values.Time `xml:"DateLastUpdated" json:"DateLastUpdated"`
			DatePublic         values.Time `xml:"DatePublic" json:"DatePublic"`
		} `xml:"VulinfoData" json:"VulinfoData"`
	} `xml:"Vulinfo" json:"Vulinfo"`
}

//Unmarshal returns JVNRSS value, unmarshalling XML data
func Unmarshal(data []byte) (*VULDEF, error) {
	vuln := &VULDEF{}
	if err := xml.Unmarshal(data, vuln); err != nil {
		return nil, err
	}
	return vuln, nil
}

//JSON returns JSON string
func (vuln *VULDEF) JSON(indent string) ([]byte, error) {
	if len(indent) > 0 {
		return json.MarshalIndent(vuln, "", indent)
	}
	return json.Marshal(vuln)
}

//Append appends Vulinfo data
func (vuln *VULDEF) Append(appnd *VULDEF) {
	if vuln == nil || appnd == nil {
		return
	}
	if len(appnd.Vulinfo) == 0 {
		return
	}
	vuln.Vulinfo = append(vuln.Vulinfo, appnd.Vulinfo...)
}

//Sort sorts Vulinfo data
func (vuln *VULDEF) SortByID(reverseFlag bool) {
	sort.Slice(vuln.Vulinfo, func(i int, j int) bool {
		if reverseFlag {
			return strings.Compare(vuln.Vulinfo[i].VulinfoID, vuln.Vulinfo[j].VulinfoID) > 0
		} else {
			return strings.Compare(vuln.Vulinfo[i].VulinfoID, vuln.Vulinfo[j].VulinfoID) < 0
		}
	})
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
