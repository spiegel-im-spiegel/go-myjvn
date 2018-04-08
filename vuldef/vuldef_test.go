package vuldef

import (
	"testing"
)

var (
	data1 = `
<?xml version="1.0" encoding="UTF-8" ?>
<VULDEF-Document
version="3.2"
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns="http://jvn.jp/vuldef/"
xmlns:vuldef="http://jvn.jp/vuldef/"
xmlns:status="http://jvndb.jvn.jp/myjvn/Status"
xmlns:sec="http://jvn.jp/rss/mod_sec/3.0/"
xmlns:marking="http://data-marking.mitre.org/Marking-1"
xmlns:tlpMarking="http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1"
xsi:schemaLocation="http://jvn.jp/vuldef/ https://jvndb.jvn.jp/schema/vuldef_3.2.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd"
xml:lang="ja">
    <Vulinfo>
      <VulinfoID>JVNDB-2018-000022</VulinfoID>
      <VulinfoData>
        <Title>WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性</Title>
        <VulinfoDescription>
          <Overview>Soflyy が提供する WordPress 用プラグイン WP All Import には、ファイルアップロード機能に関するクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#60032768 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: ゲヒルン株式会社 マルダン ムイデン 氏</Overview>
        </VulinfoDescription>
        <Affected>
          <AffectedItem>
            <Name>Soflyy</Name>
            <ProductName>WP All Import</ProductName>
            <Cpe version="2.2">cpe:/a:misc:soflyy_wp_all_import</Cpe>
            <VersionNumber>3.4.6 より前のバージョン</VersionNumber>
          </AffectedItem>
        </Affected>
        <Impact>
          <Cvss version="2.0">
            <Severity type="Base">Medium</Severity>
            <Base>4.3</Base>
            <Vector>AV:N/AC:M/Au:N/C:N/I:P/A:N</Vector>
          </Cvss>
          <Cvss version="3.0">
            <Severity type="Base">Medium</Severity>
            <Base>6.1</Base>
            <Vector>CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N</Vector>
          </Cvss>
          <ImpactItem>
            <Description>ユーザのウェブブラウザ上で任意のスクリプトを実行される可能性があります。</Description>
          </ImpactItem>
        </Impact>
        <Solution>
          <SolutionItem>
            <Description>[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。</Description>
          </SolutionItem>
        </Solution>
        <Related>
          <RelatedItem type="vendor">
            <Name>Soflyy</Name>
            <VulinfoID>Changeset 1742744 - WordPress Plugin Repository</VulinfoID>
            <URL>https://plugins.trac.wordpress.org/changeset/1742744/</URL>
          </RelatedItem>
          <RelatedItem type="vendor">
            <Name>Soflyy</Name>
            <VulinfoID>Import any XML or CSV File to WordPress - WordPress Plugins - Changelog</VulinfoID>
            <URL>https://wordpress.org/plugins/wp-all-import/#developers</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>Common Vulnerabilities and Exposures (CVE)</Name>
            <VulinfoID>CVE-2018-0546</VulinfoID>
            <URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0546</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>JVN</Name>
            <VulinfoID>JVN#33527174</VulinfoID>
            <URL>https://jvn.jp/jp/JVN33527174/index.html</URL>
          </RelatedItem>
          <RelatedItem type="cwe">
            <Name>JVNDB</Name>
            <VulinfoID>CWE-79</VulinfoID>
            <Title>クロスサイトスクリプティング</Title>
            <URL>https://jvndb.jvn.jp/ja/cwe/CWE-79.html</URL>
          </RelatedItem>
        </Related>
        <History>
          <HistoryItem>
            <HistoryNo>1</HistoryNo>
            <DateTime>2018-03-02T15:12:18+09:00</DateTime>
            <Description>[2018年03月08日]\n  掲載\n</Description>
          </HistoryItem>
        </History>
        <DateFirstPublished>2018-03-08T12:02:28+09:00</DateFirstPublished>
        <DateLastUpdated>2018-03-08T12:02:28+09:00</DateLastUpdated>
        <DatePublic>2018-03-08T00:00:00+09:00</DatePublic>
      </VulinfoData>
    </Vulinfo>
    <Vulinfo>
      <VulinfoID>JVNDB-2018-000024</VulinfoID>
      <VulinfoData>
        <Title>CG-WGR1200 における複数の脆弱性</Title>
        <VulinfoDescription>
          <Overview>株式会社コレガが提供する CG-WGR1200 は無線 LAN ルータです。CG-WGR1200 には、次の複数の脆弱性が存在します。 ・バッファオーバーフロー (CWE-119) - CVE-2017-10852 ・OS コマンドインジェクション (CWE-78) - CVE-2017-10853 ・認証不備 (CWE-306) - CVE-2017-10854  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: 三井物産セキュアディレクション株式会社 塚本 泰三 氏</Overview>
        </VulinfoDescription>
        <Affected>
          <AffectedItem>
            <Name>株式会社コレガ</Name>
            <ProductName>CG-WGR1200</ProductName>
            <Cpe version="2.2">cpe:/h:corega:cg-wgr1200</Cpe>
            <VersionNumber>ファームウエア 2.20 およびそれ以前</VersionNumber>
          </AffectedItem>
        </Affected>
        <Impact>
          <Cvss version="2.0">
            <Severity type="Base">Medium</Severity>
            <Base>5.8</Base>
            <Vector>AV:A/AC:L/Au:N/C:P/I:P/A:P</Vector>
          </Cvss>
          <Cvss version="3.0">
            <Severity type="Base">High</Severity>
            <Base>8.8</Base>
            <Vector>CVSS:3.0/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H</Vector>
          </Cvss>
          <ImpactItem>
            <Description>想定される影響は各脆弱性により異なりますが、次のような影響を受ける可能性があります。 ・当該製品にアクセス可能な第三者によって、任意のコードを実行される - CVE-2017-10852 ・当該製品にアクセス可能な第三者によって、任意の OS コマンドを実行される - CVE-2017-10853 ・当該製品にアクセス可能な第三者によって、ログインパスワードを変更される。結果として、管理画面にログインされ、当該機器の設定変更といった任意の操作がおこなわれる - CVE-2017-10854</Description>
          </ImpactItem>
        </Impact>
        <Solution>
          <SolutionItem>
            <Description>[CG-WGR1200 を使用しない] CG-WGR1200 を使用しないでください。CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありません。  [ワークアラウンドを実施する] CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありませんが、当該製品を引き続き使用する場合には、次の回避策を実施し、脆弱性による影響を軽減するようにしてください。 ・第三者が外部から当該製品にアクセスできないようリモート接続機能を無効にする ・LAN 内からルータに対する不正なアクセスを防止する 詳しくは、開発者が提供する情報をご確認ください。</Description>
          </SolutionItem>
        </Solution>
        <Related>
          <RelatedItem type="vendor">
            <Name>corega</Name>
            <VulinfoID>CG-WGR1200 における複数の脆弱性について</VulinfoID>
            <URL>http://corega.jp/support/security/20180309_wgr1200.htm</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>Common Vulnerabilities and Exposures (CVE)</Name>
            <VulinfoID>CVE-2017-10852</VulinfoID>
            <URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10852</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>Common Vulnerabilities and Exposures (CVE)</Name>
            <VulinfoID>CVE-2017-10853</VulinfoID>
            <URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10853</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>Common Vulnerabilities and Exposures (CVE)</Name>
            <VulinfoID>CVE-2017-10854</VulinfoID>
            <URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10854</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>JVN</Name>
            <VulinfoID>JVN#15201064</VulinfoID>
            <URL>https://jvn.jp/jp/JVN15201064/index.html</URL>
          </RelatedItem>
          <RelatedItem type="cwe">
            <Name>JVNDB</Name>
            <VulinfoID>CWE-78</VulinfoID>
            <Title>OSコマンドインジェクション</Title>
            <URL>https://jvndb.jvn.jp/ja/cwe/CWE-78.html</URL>
          </RelatedItem>
          <RelatedItem type="cwe">
            <Name>JVNDB</Name>
            <VulinfoID>CWE-19</VulinfoID>
            <Title>データ処理</Title>
            <URL>https://cwe.mitre.org/data/definitions/19.html</URL>
          </RelatedItem>
          <RelatedItem type="cwe">
            <Name>JVNDB</Name>
            <VulinfoID>CWE-264</VulinfoID>
            <Title>認可・権限・アクセス制御</Title>
            <URL>https://jvndb.jvn.jp/ja/cwe/CWE-264.html</URL>
          </RelatedItem>
        </Related>
        <History>
          <HistoryItem>
            <HistoryNo>1</HistoryNo>
            <DateTime>2018-03-09T12:06:58+09:00</DateTime>
            <Description>[2018年03月09日]\n  掲載\n</Description>
          </HistoryItem>
        </History>
        <DateFirstPublished>2018-03-09T12:04:48+09:00</DateFirstPublished>
        <DateLastUpdated>2018-03-09T12:04:48+09:00</DateLastUpdated>
        <DatePublic>2018-03-09T00:00:00+09:00</DatePublic>
      </VulinfoData>
    </Vulinfo>
    <sec:handling>
      <marking:Marking>
        <marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/>
      </marking:Marking>
    </sec:handling>
    <status:Status version="3.3" method="getVulnDetailInfo" lang="ja" retCd="0" retMax="10" errCd="" errMsg="" totalRes="2" totalResRet="2" firstRes="1" feed="hnd" vulnId="JVNDB-2018-000024+JVNDB-2018-000022"/>
  </VULDEF-Document>
`
	data2 = `
<?xml version="1.0" encoding="UTF-8" ?>
<VULDEF-Document
version="3.2"
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns="http://jvn.jp/vuldef/"
xmlns:vuldef="http://jvn.jp/vuldef/"
xmlns:status="http://jvndb.jvn.jp/myjvn/Status"
xmlns:sec="http://jvn.jp/rss/mod_sec/3.0/"
xmlns:marking="http://data-marking.mitre.org/Marking-1"
xmlns:tlpMarking="http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1"
xsi:schemaLocation="http://jvn.jp/vuldef/ https://jvndb.jvn.jp/schema/vuldef_3.2.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd"
xml:lang="ja">
    <Vulinfo>
      <VulinfoID>JVNDB-2018-000023</VulinfoID>
      <VulinfoData>
        <Title>WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性</Title>
        <VulinfoDescription>
          <Overview>Soflyy が提供する WordPress 用プラグイン WP All Import には、反射型のクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#33527174 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: NTTコミュニケーションズ株式会社 東内裕二 氏</Overview>
        </VulinfoDescription>
        <Affected>
          <AffectedItem>
            <Name>Soflyy</Name>
            <ProductName>WP All Import</ProductName>
            <Cpe version="2.2">cpe:/a:misc:soflyy_wp_all_import</Cpe>
            <VersionNumber>3.4.7 より前のバージョン</VersionNumber>
          </AffectedItem>
        </Affected>
        <Impact>
          <Cvss version="2.0">
            <Severity type="Base">Low</Severity>
            <Base>2.6</Base>
            <Vector>AV:N/AC:H/Au:N/C:N/I:P/A:N</Vector>
          </Cvss>
          <Cvss version="3.0">
            <Severity type="Base">Medium</Severity>
            <Base>6.1</Base>
            <Vector>CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N</Vector>
          </Cvss>
          <ImpactItem>
            <Description>当該製品にログインしているユーザのウェブブラウザ上で、任意のスクリプトを実行される可能性があります。</Description>
          </ImpactItem>
        </Impact>
        <Solution>
          <SolutionItem>
            <Description>[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。</Description>
          </SolutionItem>
        </Solution>
        <Related>
          <RelatedItem type="vendor">
            <Name>Soflyy</Name>
            <VulinfoID>Changeset 1827741 - WordPress Plugin Repository</VulinfoID>
            <URL>https://plugins.trac.wordpress.org/changeset/1827741/</URL>
          </RelatedItem>
          <RelatedItem type="vendor">
            <Name>Soflyy</Name>
            <VulinfoID>Import any XML or CSV File to WordPress - WordPress Plugins - Changelog</VulinfoID>
            <URL>https://wordpress.org/plugins/wp-all-import/#developers</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>Common Vulnerabilities and Exposures (CVE)</Name>
            <VulinfoID>CVE-2018-0547</VulinfoID>
            <URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0547</URL>
          </RelatedItem>
          <RelatedItem type="advisory">
            <Name>JVN</Name>
            <VulinfoID>JVN#60032768</VulinfoID>
            <URL>https://jvn.jp/jp/JVN60032768/index.html</URL>
          </RelatedItem>
          <RelatedItem type="cwe">
            <Name>JVNDB</Name>
            <VulinfoID>CWE-79</VulinfoID>
            <Title>クロスサイトスクリプティング</Title>
            <URL>https://jvndb.jvn.jp/ja/cwe/CWE-79.html</URL>
          </RelatedItem>
        </Related>
        <History>
          <HistoryItem>
            <HistoryNo>1</HistoryNo>
            <DateTime>2018-03-07T16:15:05+09:00</DateTime>
            <Description>[2018年03月08日]\n  掲載\n</Description>
          </HistoryItem>
        </History>
        <DateFirstPublished>2018-03-08T12:04:53+09:00</DateFirstPublished>
        <DateLastUpdated>2018-03-08T12:04:53+09:00</DateLastUpdated>
        <DatePublic>2018-03-08T00:00:00+09:00</DatePublic>
      </VulinfoData>
    </Vulinfo>
    <sec:handling>
      <marking:Marking>
        <marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/>
      </marking:Marking>
    </sec:handling>
    <status:Status version="3.3" method="getVulnDetailInfo" lang="ja" retCd="0" retMax="10" errCd="" errMsg="" totalRes="1" totalResRet="1" firstRes="1" feed="hnd" vulnId="JVNDB-2018-000023"/>
  </VULDEF-Document>
`

	res1 = `{
  "Vulinfo": [
    {
      "VulinfoID": "JVNDB-2018-000022",
      "VulinfoData": {
        "Title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
        "Overview": "Soflyy が提供する WordPress 用プラグイン WP All Import には、ファイルアップロード機能に関するクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#60032768 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: ゲヒルン株式会社 マルダン ムイデン 氏",
        "Affected": [
          {
            "Name": "Soflyy",
            "ProductName": "WP All Import",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/a:misc:soflyy_wp_all_import"
            },
            "VersionNumber": [
              "3.4.6 より前のバージョン"
            ]
          }
        ],
        "Impact": {
          "Description": "ユーザのウェブブラウザ上で任意のスクリプトを実行される可能性があります。",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:N/AC:M/Au:N/C:N/I:P/A:N",
              "BaseScore": "4.3",
              "Severity": "Medium"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
              "BaseScore": "6.1",
              "Severity": "Medium"
            }
          ]
        },
        "Solution": {
          "Description": "[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Changeset 1742744 - WordPress Plugin Repository",
            "URL": "https://plugins.trac.wordpress.org/changeset/1742744/"
          },
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Import any XML or CSV File to WordPress - WordPress Plugins - Changelog",
            "URL": "https://wordpress.org/plugins/wp-all-import/#developers"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2018-0546",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0546"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#33527174",
            "URL": "https://jvn.jp/jp/JVN33527174/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-79",
            "Title": "クロスサイトスクリプティング",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-02T15:12:18+09:00",
            "Description": "[2018年03月08日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-08T12:02:28+09:00",
        "DateLastUpdated": "2018-03-08T12:02:28+09:00",
        "DatePublic": "2018-03-08T00:00:00+09:00"
      }
    },
    {
      "VulinfoID": "JVNDB-2018-000024",
      "VulinfoData": {
        "Title": "CG-WGR1200 における複数の脆弱性",
        "Overview": "株式会社コレガが提供する CG-WGR1200 は無線 LAN ルータです。CG-WGR1200 には、次の複数の脆弱性が存在します。 ・バッファオーバーフロー (CWE-119) - CVE-2017-10852 ・OS コマンドインジェクション (CWE-78) - CVE-2017-10853 ・認証不備 (CWE-306) - CVE-2017-10854  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: 三井物産セキュアディレクション株式会社 塚本 泰三 氏",
        "Affected": [
          {
            "Name": "株式会社コレガ",
            "ProductName": "CG-WGR1200",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/h:corega:cg-wgr1200"
            },
            "VersionNumber": [
              "ファームウエア 2.20 およびそれ以前"
            ]
          }
        ],
        "Impact": {
          "Description": "想定される影響は各脆弱性により異なりますが、次のような影響を受ける可能性があります。 ・当該製品にアクセス可能な第三者によって、任意のコードを実行される - CVE-2017-10852 ・当該製品にアクセス可能な第三者によって、任意の OS コマンドを実行される - CVE-2017-10853 ・当該製品にアクセス可能な第三者によって、ログインパスワードを変更される。結果として、管理画面にログインされ、当該機器の設定変更といった任意の操作がおこなわれる - CVE-2017-10854",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:A/AC:L/Au:N/C:P/I:P/A:P",
              "BaseScore": "5.8",
              "Severity": "Medium"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
              "BaseScore": "8.8",
              "Severity": "High"
            }
          ]
        },
        "Solution": {
          "Description": "[CG-WGR1200 を使用しない] CG-WGR1200 を使用しないでください。CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありません。  [ワークアラウンドを実施する] CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありませんが、当該製品を引き続き使用する場合には、次の回避策を実施し、脆弱性による影響を軽減するようにしてください。 ・第三者が外部から当該製品にアクセスできないようリモート接続機能を無効にする ・LAN 内からルータに対する不正なアクセスを防止する 詳しくは、開発者が提供する情報をご確認ください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "corega",
            "VulinfoID": "CG-WGR1200 における複数の脆弱性について",
            "URL": "http://corega.jp/support/security/20180309_wgr1200.htm"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10852",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10852"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10853",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10853"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10854",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10854"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#15201064",
            "URL": "https://jvn.jp/jp/JVN15201064/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-78",
            "Title": "OSコマンドインジェクション",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-78.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-19",
            "Title": "データ処理",
            "URL": "https://cwe.mitre.org/data/definitions/19.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-264",
            "Title": "認可・権限・アクセス制御",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-264.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-09T12:06:58+09:00",
            "Description": "[2018年03月09日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-09T12:04:48+09:00",
        "DateLastUpdated": "2018-03-09T12:04:48+09:00",
        "DatePublic": "2018-03-09T00:00:00+09:00"
      }
    }
  ]
}`
	res2 = `{
  "Vulinfo": [
    {
      "VulinfoID": "JVNDB-2018-000023",
      "VulinfoData": {
        "Title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
        "Overview": "Soflyy が提供する WordPress 用プラグイン WP All Import には、反射型のクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#33527174 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: NTTコミュニケーションズ株式会社 東内裕二 氏",
        "Affected": [
          {
            "Name": "Soflyy",
            "ProductName": "WP All Import",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/a:misc:soflyy_wp_all_import"
            },
            "VersionNumber": [
              "3.4.7 より前のバージョン"
            ]
          }
        ],
        "Impact": {
          "Description": "当該製品にログインしているユーザのウェブブラウザ上で、任意のスクリプトを実行される可能性があります。",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:N/AC:H/Au:N/C:N/I:P/A:N",
              "BaseScore": "2.6",
              "Severity": "Low"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
              "BaseScore": "6.1",
              "Severity": "Medium"
            }
          ]
        },
        "Solution": {
          "Description": "[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Changeset 1827741 - WordPress Plugin Repository",
            "URL": "https://plugins.trac.wordpress.org/changeset/1827741/"
          },
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Import any XML or CSV File to WordPress - WordPress Plugins - Changelog",
            "URL": "https://wordpress.org/plugins/wp-all-import/#developers"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2018-0547",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0547"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#60032768",
            "URL": "https://jvn.jp/jp/JVN60032768/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-79",
            "Title": "クロスサイトスクリプティング",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-07T16:15:05+09:00",
            "Description": "[2018年03月08日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-08T12:04:53+09:00",
        "DateLastUpdated": "2018-03-08T12:04:53+09:00",
        "DatePublic": "2018-03-08T00:00:00+09:00"
      }
    }
  ]
}`
	res3a = `{
  "Vulinfo": [
    {
      "VulinfoID": "JVNDB-2018-000022",
      "VulinfoData": {
        "Title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
        "Overview": "Soflyy が提供する WordPress 用プラグイン WP All Import には、ファイルアップロード機能に関するクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#60032768 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: ゲヒルン株式会社 マルダン ムイデン 氏",
        "Affected": [
          {
            "Name": "Soflyy",
            "ProductName": "WP All Import",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/a:misc:soflyy_wp_all_import"
            },
            "VersionNumber": [
              "3.4.6 より前のバージョン"
            ]
          }
        ],
        "Impact": {
          "Description": "ユーザのウェブブラウザ上で任意のスクリプトを実行される可能性があります。",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:N/AC:M/Au:N/C:N/I:P/A:N",
              "BaseScore": "4.3",
              "Severity": "Medium"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
              "BaseScore": "6.1",
              "Severity": "Medium"
            }
          ]
        },
        "Solution": {
          "Description": "[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Changeset 1742744 - WordPress Plugin Repository",
            "URL": "https://plugins.trac.wordpress.org/changeset/1742744/"
          },
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Import any XML or CSV File to WordPress - WordPress Plugins - Changelog",
            "URL": "https://wordpress.org/plugins/wp-all-import/#developers"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2018-0546",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0546"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#33527174",
            "URL": "https://jvn.jp/jp/JVN33527174/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-79",
            "Title": "クロスサイトスクリプティング",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-02T15:12:18+09:00",
            "Description": "[2018年03月08日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-08T12:02:28+09:00",
        "DateLastUpdated": "2018-03-08T12:02:28+09:00",
        "DatePublic": "2018-03-08T00:00:00+09:00"
      }
    },
    {
      "VulinfoID": "JVNDB-2018-000024",
      "VulinfoData": {
        "Title": "CG-WGR1200 における複数の脆弱性",
        "Overview": "株式会社コレガが提供する CG-WGR1200 は無線 LAN ルータです。CG-WGR1200 には、次の複数の脆弱性が存在します。 ・バッファオーバーフロー (CWE-119) - CVE-2017-10852 ・OS コマンドインジェクション (CWE-78) - CVE-2017-10853 ・認証不備 (CWE-306) - CVE-2017-10854  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: 三井物産セキュアディレクション株式会社 塚本 泰三 氏",
        "Affected": [
          {
            "Name": "株式会社コレガ",
            "ProductName": "CG-WGR1200",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/h:corega:cg-wgr1200"
            },
            "VersionNumber": [
              "ファームウエア 2.20 およびそれ以前"
            ]
          }
        ],
        "Impact": {
          "Description": "想定される影響は各脆弱性により異なりますが、次のような影響を受ける可能性があります。 ・当該製品にアクセス可能な第三者によって、任意のコードを実行される - CVE-2017-10852 ・当該製品にアクセス可能な第三者によって、任意の OS コマンドを実行される - CVE-2017-10853 ・当該製品にアクセス可能な第三者によって、ログインパスワードを変更される。結果として、管理画面にログインされ、当該機器の設定変更といった任意の操作がおこなわれる - CVE-2017-10854",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:A/AC:L/Au:N/C:P/I:P/A:P",
              "BaseScore": "5.8",
              "Severity": "Medium"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
              "BaseScore": "8.8",
              "Severity": "High"
            }
          ]
        },
        "Solution": {
          "Description": "[CG-WGR1200 を使用しない] CG-WGR1200 を使用しないでください。CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありません。  [ワークアラウンドを実施する] CG-WGR1200 のサポートは終了しているため、対策版ファームウェアのリリース予定はありませんが、当該製品を引き続き使用する場合には、次の回避策を実施し、脆弱性による影響を軽減するようにしてください。 ・第三者が外部から当該製品にアクセスできないようリモート接続機能を無効にする ・LAN 内からルータに対する不正なアクセスを防止する 詳しくは、開発者が提供する情報をご確認ください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "corega",
            "VulinfoID": "CG-WGR1200 における複数の脆弱性について",
            "URL": "http://corega.jp/support/security/20180309_wgr1200.htm"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10852",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10852"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10853",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10853"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2017-10854",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-10854"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#15201064",
            "URL": "https://jvn.jp/jp/JVN15201064/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-78",
            "Title": "OSコマンドインジェクション",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-78.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-19",
            "Title": "データ処理",
            "URL": "https://cwe.mitre.org/data/definitions/19.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-264",
            "Title": "認可・権限・アクセス制御",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-264.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-09T12:06:58+09:00",
            "Description": "[2018年03月09日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-09T12:04:48+09:00",
        "DateLastUpdated": "2018-03-09T12:04:48+09:00",
        "DatePublic": "2018-03-09T00:00:00+09:00"
      }
    },
    {
      "VulinfoID": "JVNDB-2018-000023",
      "VulinfoData": {
        "Title": "WordPress 用プラグイン WP All Import におけるクロスサイトスクリプティングの脆弱性",
        "Overview": "Soflyy が提供する WordPress 用プラグイン WP All Import には、反射型のクロスサイトスクリプティング (CWE-79) の脆弱性が存在します。  なお、本脆弱性は JVN#33527174 とは異なる問題です。  この脆弱性情報は、情報セキュリティ早期警戒パートナーシップに基づき下記の方が IPA に報告し、JPCERT/CC が開発者との調整を行いました。 報告者: NTTコミュニケーションズ株式会社 東内裕二 氏",
        "Affected": [
          {
            "Name": "Soflyy",
            "ProductName": "WP All Import",
            "Cpe": {
              "version": "2.2",
              "value": "cpe:/a:misc:soflyy_wp_all_import"
            },
            "VersionNumber": [
              "3.4.7 より前のバージョン"
            ]
          }
        ],
        "Impact": {
          "Description": "当該製品にログインしているユーザのウェブブラウザ上で、任意のスクリプトを実行される可能性があります。",
          "Cvss": [
            {
              "Version": "2.0",
              "BaseVector": "AV:N/AC:H/Au:N/C:N/I:P/A:N",
              "BaseScore": "2.6",
              "Severity": "Low"
            },
            {
              "Version": "3.0",
              "BaseVector": "CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
              "BaseScore": "6.1",
              "Severity": "Medium"
            }
          ]
        },
        "Solution": {
          "Description": "[アップデートする] 開発者が提供する情報をもとに、最新版へアップデートしてください。"
        },
        "Related": [
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Changeset 1827741 - WordPress Plugin Repository",
            "URL": "https://plugins.trac.wordpress.org/changeset/1827741/"
          },
          {
            "Type": "vendor",
            "Name": "Soflyy",
            "VulinfoID": "Import any XML or CSV File to WordPress - WordPress Plugins - Changelog",
            "URL": "https://wordpress.org/plugins/wp-all-import/#developers"
          },
          {
            "Type": "advisory",
            "Name": "Common Vulnerabilities and Exposures (CVE)",
            "VulinfoID": "CVE-2018-0547",
            "URL": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-0547"
          },
          {
            "Type": "advisory",
            "Name": "JVN",
            "VulinfoID": "JVN#60032768",
            "URL": "https://jvn.jp/jp/JVN60032768/index.html"
          },
          {
            "Type": "cwe",
            "Name": "JVNDB",
            "VulinfoID": "CWE-79",
            "Title": "クロスサイトスクリプティング",
            "URL": "https://jvndb.jvn.jp/ja/cwe/CWE-79.html"
          }
        ],
        "History": [
          {
            "HistoryNo": 1,
            "DateTime": "2018-03-07T16:15:05+09:00",
            "Description": "[2018年03月08日]\\n  掲載\\n"
          }
        ],
        "DateFirstPublished": "2018-03-08T12:04:53+09:00",
        "DateLastUpdated": "2018-03-08T12:04:53+09:00",
        "DatePublic": "2018-03-08T00:00:00+09:00"
      }
    }
  ]
}`
)

func TestUnmarshal(t *testing.T) {
	testCaces := []struct {
		data string
		res  string
	}{
		{data: data1, res: res1},
		{data: data2, res: res2},
	}
	for _, tc := range testCaces {
		vuln, err := Unmarshal([]byte(tc.data))
		if err != nil {
			t.Errorf("Unmarshal() = \"%v\", want nil.", err)
			continue
		}
		json, err := vuln.JSON("  ")
		if err != nil {
			t.Errorf("JSON() = \"%v\", want nil.", err)
			continue
		}
		if string(json) != tc.res {
			t.Errorf("Unmarshal().JSON() = \"%v\", want \"%v\".", string(json), tc.res)
		}
	}

}

func TestAppend(t *testing.T) {
	vuln1, err := Unmarshal([]byte(data1))
	if err != nil {
		t.Errorf("Unmarshal() = \"%v\", want nil.", err)
		return
	}
	vuln2, err := Unmarshal([]byte(data2))
	if err != nil {
		t.Errorf("Unmarshal() = \"%v\", want nil.", err)
	}
	vuln1.Append(vuln2)
	json, err := vuln1.JSON("  ")
	if err != nil {
		t.Errorf("JSON() = \"%v\", want nil.", err)
		return
	}
	if string(json) != res3a {
		t.Errorf("Unmarshal().JSON() = \"%v\", want \"%v\".", string(json), res3a)
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
