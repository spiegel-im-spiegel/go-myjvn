package rss

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

//Time as time.Time type
type Time struct {
	time.Time
}

//UnmarshalXML returns result of Unmarshal for xml.Unmarshal()
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	tm, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*t = Time{tm}
	return nil
}

//JVNRSS is structure of JVNRSS
type JVNRSS struct {
	Channel struct {
		Title       string `xml:"title" json:"title"`
		Description string `xml:"description" json:"description"`
		Link        string `xml:"link" json:"link"`
		Date        Time   `xml:"date" json:"date"`
		Modified    Time   `xml:"modified" json:"modified"`
	} `xml:"channel" json:"channel"`
	Items []struct {
		Title       string `xml:"title" json:"title"`
		Description string `xml:"description" json:"description"`
		Link        string `xml:"link" json:"link"`
		Date        Time   `xml:"date" json:"date"`
		Issued      Time   `xml:"issued" json:"issued"`
		Modified    Time   `xml:"modified" json:"modified"`
		Creator     string `xml:"creator" json:"creator"`
		Identifier  string `xml:"identifier" json:"identifier"`
		References  []struct {
			ID     string `xml:"id,attr" json:"id"`
			Title  string `xml:"title,attr,omitempty" json:"title,omitempty"`
			Source string `xml:"source,attr,omitempty" json:"source,omitempty"`
			Value  string `xml:",chardata" json:"value"`
		} `xml:"references" json:"references"`
		Cpe struct {
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
	Status struct {
		Version     string `xml:"version,attr" json:"version"`         //MyJVN API Schema Version
		Method      string `xml:"method,attr" json:"method"`           //Method
		Language    string `xml:"lang,attr" json:"lang"`               //Language (ja/en)
		Feed        string `xml:"feed,attr" json:"feed"`               //Feed version
		ReturnCode  int    `xml:"retCd,attr" json:"retCd"`             //Return Code/Interger (0:success, 1:failure)
		ReturnMax   int    `xml:"retMax,attr" json:"retMax"`           //Maximum number of Entry/Interger
		ErrorCode   string `xml:"errCd,attr" json:"errCd"`             //Error Code (Null:success)
		ErrorMsg    string `xml:"errMsg,attr" json:"errMsg"`           //Error Message (Null:success)
		TotalRes    int    `xml:"totalRes,attr" json:"totalRes"`       //Total number of Result entries
		TotalResRet int    `xml:"totalResRet,attr" json:"totalResRet"` //Number of Result entries
		FirstRes    int    `xml:"firstRes,attr" json:"firstRes"`       //Start entry number in Result entries
	} `xml:"Status" json:"Status"`
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
