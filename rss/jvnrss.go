package rss

import (
	"encoding/json"
	"encoding/xml"

	"github.com/spiegel-im-spiegel/go-myjvn/values"
)

const (
	//MaxItems is max of items for "getVulnOverviewList" API
	MaxItems = 50
)

//JVNRSS is structure of JVNRSS
type JVNRSS struct {
	Channel struct {
		Title       string      `xml:"title" json:"title"`
		Description string      `xml:"description" json:"description"`
		Link        string      `xml:"link" json:"link"`
		Date        values.Time `xml:"date" json:"date"`
		Modified    values.Time `xml:"modified" json:"modified"`
	} `xml:"channel" json:"channel"`
	Items []struct {
		Title       string      `xml:"title" json:"title"`
		Description string      `xml:"description" json:"description"`
		Link        string      `xml:"link" json:"link"`
		Date        values.Time `xml:"date" json:"date"`
		Issued      values.Time `xml:"issued" json:"issued"`
		Modified    values.Time `xml:"modified" json:"modified"`
		Creator     string      `xml:"creator" json:"creator"`
		Identifier  string      `xml:"identifier" json:"identifier"`
		References  []struct {
			ID     string `xml:"id,attr" json:"id"`
			Title  string `xml:"title,attr,omitempty" json:"title,omitempty"`
			Source string `xml:"source,attr,omitempty" json:"source,omitempty"`
			Value  string `xml:",chardata" json:"value"`
		} `xml:"references" json:"references"`
		Cpe []struct {
			Version string `xml:"version,attr" json:"version"`
			Vendor  string `xml:"vendor,attr" json:"vendor"`
			Product string `xml:"product,attr" json:"product"`
			Value   string `xml:",chardata" json:"value"`
		} `xml:"cpe" json:"cpe"`
		CVSS []struct {
			Version  string `xml:"version,attr" json:"version"`
			Type     string `xml:"type,attr" json:"type"`
			Vector   string `xml:"vector,attr" json:"vector"`
			Score    string `xml:"score,attr" json:"score"`
			Severity string `xml:"severity,attr" json:"severity"`
		} `xml:"cvss" json:"cvss"`
	} `xml:"item" json:"item"`
}

//Unmarshal returns JVNRSS value, unmarshalling XML data
func Unmarshal(data []byte) (*JVNRSS, error) {
	rss := &JVNRSS{}
	if err := xml.Unmarshal(data, rss); err != nil {
		return nil, err
	}
	return rss, nil
}

//JSON returns JSON string
func (rss *JVNRSS) JSON(indent string) ([]byte, error) {
	if len(indent) > 0 {
		return json.MarshalIndent(rss, "", indent)
	}
	return json.Marshal(rss)
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
