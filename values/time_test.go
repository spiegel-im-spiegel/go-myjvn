package values

import (
	"encoding/xml"
	"testing"
	"time"
)

type ForTestStruct struct {
	Channel struct {
		Date Time `xml:"date" json:"date"`
	} `xml:"channel" json:"channel"`
}

func TestUnmarshal(t *testing.T) {
	data := `
<?xml version="1.0" encoding="UTF-8" ?>
<rdf:RDF
xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
xmlns:status="http://jvndb.jvn.jp/myjvn/Status"
xml:lang="ja">
  <channel rdf:about="https://jvndb.jvn.jp/apis/myjvn">
    <dc:date>2018-03-13T11:57:52+09:00</dc:date>
  </channel>
</rdf:RDF>
`
	tst := &ForTestStruct{}
	if err := xml.Unmarshal([]byte(data), tst); err != nil {
		t.Errorf("Unmarshal() = \"%v\", want nil.", err)
		return
	}
	res := "2018-03-13T11:57:52+09:00"
	str := tst.Channel.Date.Format(time.RFC3339)
	if str != res {
		t.Errorf("ForTestStruct.Channel.Date = \"%v\", want \"%v\".", str, res)
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
