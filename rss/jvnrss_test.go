package rss

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data := `
<?xml version="1.0" encoding="UTF-8" ?>
<rdf:RDF
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns="http://purl.org/rss/1.0/"
xmlns:rss="http://purl.org/rss/1.0/"
xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
xmlns:dc="http://purl.org/dc/elements/1.1/"
xmlns:dcterms="http://purl.org/dc/terms/"
xmlns:sec="http://jvn.jp/rss/mod_sec/3.0/"
xmlns:marking="http://data-marking.mitre.org/Marking-1"
xmlns:tlpMarking="http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1"
xmlns:status="http://jvndb.jvn.jp/myjvn/Status"
xsi:schemaLocation="http://purl.org/rss/1.0/ https://jvndb.jvn.jp/schema/jvnrss_3.2.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd"
xml:lang="ja">
    <channel rdf:about="https://jvndb.jvn.jp/apis/myjvn">
      <title>JVNDB　脆弱性対策情報</title>
      <link>https://jvndb.jvn.jp/apis/myjvn</link>
      <description>JVNDB　脆弱性対策情報</description>
      <dc:date>2018-03-13T11:57:52+09:00</dc:date>
      <dcterms:issued/>
      <dcterms:modified>2018-03-13T11:57:52+09:00</dcterms:modified>
      <sec:handling>
        <marking:Marking>
          <marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/>
        </marking:Marking>
      </sec:handling>
      <items>
        <rdf:Seq>
          <rdf:li rdf:resource="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html"/>
          <rdf:li rdf:resource="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000023.html"/>
          <rdf:li rdf:resource="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000022.html"/>
        </rdf:Seq>
      </items>
    </channel>
    <item rdf:about="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html">
      <title>CG-WGR1200 における複数の脆弱性</title>
      <link>https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html</link>
      <description>株式会社コレガが提供する CG-WGR1200 は無線 LAN ルータです。CG-WGR1200 には、次の複数の脆弱性が存在します。 ・バッファオーバーフロー (CWE-119) - CVE-2017-10852 ・OS コマンドインジェクション (CWE-78) - CVE-2017-10853 ・認証不備 (CWE-306) - CVE-2017-10854  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: 三井物産セキュアディレクション株式会社 塚本 泰三 氏</description>
      <dc:creator>Information-technology Promotion Agency, Japan</dc:creator>
      <sec:identifier>JVNDB-2018-000024</sec:identifier>
      <sec:references source="CVE" id="CVE-2017-10852">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10852</sec:references>
      <sec:references source="CVE" id="CVE-2017-10853">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10853</sec:references>
      <sec:references source="CVE" id="CVE-2017-10854">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10854</sec:references>
      <sec:references source="JVN" id="JVN#15201064">https://jvn.jp/jp/JVN15201064/index.html</sec:references>
      <sec:references id="CWE-19" title="データ処理(CWE-19)">https://cwe.mitre.org/data/definitions/19.html</sec:references>
      <sec:references id="CWE-264" title="認可・権限・アクセス制御(CWE-264)">https://jvndb.jvn.jp/ja/cwe/CWE-264.html</sec:references>
      <sec:references id="CWE-78" title="OSコマンドインジェクション(CWE-78)">https://jvndb.jvn.jp/ja/cwe/CWE-78.html</sec:references>
      <sec:cpe version="2.2" vendor="株式会社コレガ" product="CG-WGR1200">cpe:/h:corega:cg-wgr1200</sec:cpe>
      <sec:cvss score="8.8" severity="High" vector="CVSS:3.0/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H" version="3.0" type="Base"/>
      <sec:cvss score="5.8" severity="Medium" vector="AV:A/AC:L/Au:N/C:P/I:P/A:P" version="2.0" type="Base"/>
      <dc:date>2018-03-09T12:04:48+09:00</dc:date>
      <dcterms:issued>2018-03-09T12:04:48+09:00</dcterms:issued>
      <dcterms:modified>2018-03-09T12:04:48+09:00</dcterms:modified>
    </item>
    <item rdf:about="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000023.html">
      <title>WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性</title>
      <link>https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000023.html</link>
      <description>Soflyy が提供する WordPress 用プラグイン WP All Import には、反射型のクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#33527174 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: NTTコミュニケーションズ株式会社 東内裕二 氏</description>
      <dc:creator>Information-technology Promotion Agency, Japan</dc:creator>
      <sec:identifier>JVNDB-2018-000023</sec:identifier>
      <sec:references source="CVE" id="CVE-2018-0547">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0547</sec:references>
      <sec:references source="JVN" id="JVN#60032768">https://jvn.jp/jp/JVN60032768/index.html</sec:references>
      <sec:references id="CWE-79" title="クロスサイトスクリプティング(CWE-79)">https://jvndb.jvn.jp/ja/cwe/CWE-79.html</sec:references>
      <sec:cpe version="2.2" vendor="Soflyy" product="WP All Import">cpe:/a:misc:soflyy_wp_all_import</sec:cpe>
      <sec:cvss score="6.1" severity="Medium" vector="CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N" version="3.0" type="Base"/>
      <sec:cvss score="2.6" severity="Low" vector="AV:N/AC:H/Au:N/C:N/I:P/A:N" version="2.0" type="Base"/>
      <dc:date>2018-03-08T12:04:53+09:00</dc:date>
      <dcterms:issued>2018-03-08T12:04:53+09:00</dcterms:issued>
      <dcterms:modified>2018-03-08T12:04:53+09:00</dcterms:modified>
    </item>
    <item rdf:about="https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000022.html">
      <title>WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性</title>
      <link>https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000022.html</link>
      <description>Soflyy が提供する WordPress 用プラグイン WP All Import には、ファイルアップロード機能に関するクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#60032768 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: ゲヒルン株式会社 マルダン ムイデン 氏</description>
      <dc:creator>Information-technology Promotion Agency, Japan</dc:creator>
      <sec:identifier>JVNDB-2018-000022</sec:identifier>
      <sec:references source="CVE" id="CVE-2018-0546">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0546</sec:references>
      <sec:references source="JVN" id="JVN#33527174">https://jvn.jp/jp/JVN33527174/index.html</sec:references>
      <sec:references id="CWE-79" title="クロスサイトスクリプティング(CWE-79)">https://jvndb.jvn.jp/ja/cwe/CWE-79.html</sec:references>
      <sec:cpe version="2.2" vendor="Soflyy" product="WP All Import">cpe:/a:misc:soflyy_wp_all_import</sec:cpe>
      <sec:cvss score="6.1" severity="Medium" vector="CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N" version="3.0" type="Base"/>
      <sec:cvss score="4.3" severity="Medium" vector="AV:N/AC:M/Au:N/C:N/I:P/A:N" version="2.0" type="Base"/>
      <dc:date>2018-03-08T12:02:28+09:00</dc:date>
      <dcterms:issued>2018-03-08T12:02:28+09:00</dcterms:issued>
      <dcterms:modified>2018-03-08T12:02:28+09:00</dcterms:modified>
    </item>
    <status:Status version="3.3" method="getVulnOverviewList" lang="ja" retCd="0" retMax="50" errCd="" errMsg="" totalRes="3" totalResRet="3" firstRes="1" feed="hnd"/>
  </rdf:RDF>
`
	res := `{
  "channel": {
    "title": "JVNDB　脆弱性対策情報",
    "description": "JVNDB　脆弱性対策情報",
    "link": "https://jvndb.jvn.jp/apis/myjvn",
    "date": "2018-03-13T11:57:52+09:00",
    "modified": "2018-03-13T11:57:52+09:00"
  },
  "item": [
    {
      "title": "CG-WGR1200 における複数の脆弱性",
      "description": "株式会社コレガが提供する CG-WGR1200 は無線 LAN ルータです。CG-WGR1200 には、次の複数の脆弱性が存在します。 ・バッファオーバーフロー (CWE-119) - CVE-2017-10852 ・OS コマンドインジェクション (CWE-78) - CVE-2017-10853 ・認証不備 (CWE-306) - CVE-2017-10854  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: 三井物産セキュアディレクション株式会社 塚本 泰三 氏",
      "link": "https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000024.html",
      "date": "2018-03-09T12:04:48+09:00",
      "issued": "2018-03-09T12:04:48+09:00",
      "modified": "2018-03-09T12:04:48+09:00",
      "creator": "Information-technology Promotion Agency, Japan",
      "identifier": "JVNDB-2018-000024",
      "references": [
        {
          "id": "CVE-2017-10852",
          "source": "CVE",
          "value": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10852"
        },
        {
          "id": "CVE-2017-10853",
          "source": "CVE",
          "value": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10853"
        },
        {
          "id": "CVE-2017-10854",
          "source": "CVE",
          "value": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10854"
        },
        {
          "id": "JVN#15201064",
          "source": "JVN",
          "value": "https://jvn.jp/jp/JVN15201064/index.html"
        },
        {
          "id": "CWE-19",
          "title": "データ処理(CWE-19)",
          "value": "https://cwe.mitre.org/data/definitions/19.html"
        },
        {
          "id": "CWE-264",
          "title": "認可・権限・アクセス制御(CWE-264)",
          "value": "https://jvndb.jvn.jp/ja/cwe/CWE-264.html"
        },
        {
          "id": "CWE-78",
          "title": "OSコマンドインジェクション(CWE-78)",
          "value": "https://jvndb.jvn.jp/ja/cwe/CWE-78.html"
        }
      ],
      "cpe": {
        "version": "2.2",
        "vendor": "株式会社コレガ",
        "product": "CG-WGR1200",
        "value": "cpe:/h:corega:cg-wgr1200"
      },
      "cvss": [
        {
          "version": "3.0",
          "type": "Base",
          "vector": "CVSS:3.0/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
          "score": "8.8",
          "severity": "High"
        },
        {
          "version": "2.0",
          "type": "Base",
          "vector": "AV:A/AC:L/Au:N/C:P/I:P/A:P",
          "score": "5.8",
          "severity": "Medium"
        }
      ]
    },
    {
      "title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
      "description": "Soflyy が提供する WordPress 用プラグイン WP All Import には、反射型のクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#33527174 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: NTTコミュニケーションズ株式会社 東内裕二 氏",
      "link": "https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000023.html",
      "date": "2018-03-08T12:04:53+09:00",
      "issued": "2018-03-08T12:04:53+09:00",
      "modified": "2018-03-08T12:04:53+09:00",
      "creator": "Information-technology Promotion Agency, Japan",
      "identifier": "JVNDB-2018-000023",
      "references": [
        {
          "id": "CVE-2018-0547",
          "source": "CVE",
          "value": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0547"
        },
        {
          "id": "JVN#60032768",
          "source": "JVN",
          "value": "https://jvn.jp/jp/JVN60032768/index.html"
        },
        {
          "id": "CWE-79",
          "title": "クロスサイトスクリプティング(CWE-79)",
          "value": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
        }
      ],
      "cpe": {
        "version": "2.2",
        "vendor": "Soflyy",
        "product": "WP All Import",
        "value": "cpe:/a:misc:soflyy_wp_all_import"
      },
      "cvss": [
        {
          "version": "3.0",
          "type": "Base",
          "vector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
          "score": "6.1",
          "severity": "Medium"
        },
        {
          "version": "2.0",
          "type": "Base",
          "vector": "AV:N/AC:H/Au:N/C:N/I:P/A:N",
          "score": "2.6",
          "severity": "Low"
        }
      ]
    },
    {
      "title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
      "description": "Soflyy が提供する WordPress 用プラグイン WP All Import には、ファイルアップロード機能に関するクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#60032768 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: ゲヒルン株式会社 マルダン ムイデン 氏",
      "link": "https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000022.html",
      "date": "2018-03-08T12:02:28+09:00",
      "issued": "2018-03-08T12:02:28+09:00",
      "modified": "2018-03-08T12:02:28+09:00",
      "creator": "Information-technology Promotion Agency, Japan",
      "identifier": "JVNDB-2018-000022",
      "references": [
        {
          "id": "CVE-2018-0546",
          "source": "CVE",
          "value": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0546"
        },
        {
          "id": "JVN#33527174",
          "source": "JVN",
          "value": "https://jvn.jp/jp/JVN33527174/index.html"
        },
        {
          "id": "CWE-79",
          "title": "クロスサイトスクリプティング(CWE-79)",
          "value": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
        }
      ],
      "cpe": {
        "version": "2.2",
        "vendor": "Soflyy",
        "product": "WP All Import",
        "value": "cpe:/a:misc:soflyy_wp_all_import"
      },
      "cvss": [
        {
          "version": "3.0",
          "type": "Base",
          "vector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
          "score": "6.1",
          "severity": "Medium"
        },
        {
          "version": "2.0",
          "type": "Base",
          "vector": "AV:N/AC:M/Au:N/C:N/I:P/A:N",
          "score": "4.3",
          "severity": "Medium"
        }
      ]
    }
  ],
  "Status": {
    "version": "3.3",
    "method": "getVulnOverviewList",
    "lang": "ja",
    "feed": "hnd",
    "retCd": 0,
    "retMax": 50,
    "errCd": "",
    "errMsg": "",
    "totalRes": 3,
    "totalResRet": 3,
    "firstRes": 1
  }
}`
	rss, err := Unmarshal([]byte(data))
	if err != nil {
		t.Errorf("VulnOverviewList() = \"%v\", want nil.", err)
	}
	json, err := rss.JSON("  ")
	if err != nil {
		t.Errorf("VulnOverviewList() = \"%v\", want nil.", err)
	}
	if string(json) != res {
		t.Errorf("VulnOverviewList()Channel.Title = \"%v\", want \"%v\".", string(json), res)
	}
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
