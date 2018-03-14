package status

import (
	"encoding/json"
	"encoding/xml"
)

//Ststus is structure of Ststus
type Ststus struct {
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
func Unmarshal(data []byte) (*Ststus, error) {
	rss := &Ststus{}
	if err := xml.Unmarshal(data, rss); err != nil {
		return nil, err
	}
	return rss, nil
}

//JSON returns JSON string
func (stat *Ststus) JSON(indent string) ([]byte, error) {
	if len(indent) > 0 {
		return json.MarshalIndent(stat, "", indent)
	}
	return json.Marshal(stat)
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
