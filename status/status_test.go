package status

import "testing"

func TestUnmarshal(t *testing.T) {
	data := `
<?xml version="1.0" encoding="UTF-8" ?>
<rdf:RDF
xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
xmlns:status="http://jvndb.jvn.jp/myjvn/Status"
xml:lang="ja">
  <status:Status version="3.3" method="getVulnDetailInfo" lang="ja" retCd="1" retMax="10" errCd="XXX" errMsg="Test Error" totalRes="2" totalResRet="2" firstRes="1" feed="hnd" vulnId="JVNDB-2018-000024+JVNDB-2018-000022"/>
</rdf:RDF>
`
	res := `{
  "Status": {
    "version": "3.3",
    "method": "getVulnDetailInfo",
    "lang": "ja",
    "feed": "hnd",
    "retCd": 1,
    "retMax": 10,
    "errCd": "XXX",
    "errMsg": "Test Error",
    "totalRes": 2,
    "totalResRet": 2,
    "firstRes": 1
  }
}`
	stat, err := Unmarshal([]byte(data))
	if err != nil {
		t.Errorf("Unmarshal() = \"%v\", want nil.", err)
	}
	if err = stat.GetError(); err == nil {
		t.Error("GetError() = nil, not want nil.")
	}
	json, err := stat.JSON("  ")
	if err != nil {
		t.Errorf("JSON() = \"%v\", want nil.", err)
	}
	if string(json) != res {
		t.Errorf("Unmarshal().JSON() = \"%v\", want \"%v\".", string(json), res)
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
