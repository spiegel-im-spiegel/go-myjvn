# go-myjvn - Handling MyJVN RESTful API by Golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/go-myjvn.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/go-myjvn)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/go-myjvn/master/LICENSE)

本パッケージは [JVN] が提供する「[脆弱性対策情報共有フレームワーク]」のひとつである [MyJVN API] を [Go 言語]でハンドリングするためのラッパークラスです。
今のところ，以下の API をサポートしています。

- [脆弱性対策概要情報一覧の取得（getVulnOverviewList）](https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html "MyJVN - API: getVulnOverviewList")
    - 期間および深刻度によるフィルタリングをサポート
    - 日本語のみサポート
- [脆弱性対策詳細情報の取得（getVulnDetailInfo）](https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html "MyJVN - API: getVulnDetailInfo")
    - 日本語のみサポート

## サンプルコード

“[`go-myjvn/cli/vulnlist/main.go`](https://github.com/spiegel-im-spiegel/go-myjvn/blob/master/cli/vulnlist/main.go "")” を参照のこと。

## 参考

- [MyJVN API](https://jvndb.jvn.jp/apis/)
- [MyJVN API に関する覚え書き — しっぽのさきっちょ | text.Baldanders.info](http://text.baldanders.info/remark/2018/03/myjvn-api/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[MyJVN API]: https://jvndb.jvn.jp/apis/
