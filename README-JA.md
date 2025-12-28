![](https://libsisimai.org/static/images/logo/sisimai-x01.png)
[![License](https://img.shields.io/badge/license-BSD%202--Clause-orange.svg)](https://github.com/sisimai/go-sisimai/blob/5-stable/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sisimai/go-sisimai)
[![Go Reference](https://pkg.go.dev/badge/libsisimai.org/sisimai/v5.svg)](https://pkg.go.dev/libsisimai.org/sisimai/v5)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/sisimai/go-sisimai/make-test.yml)
![Codecov](https://img.shields.io/codecov/c/github/sisimai/go-sisimai)

> [!NOTE]
> Sisimaiは多種多様な形式のバウンスメールを解析して宛先アドレスやバウンス理由など分析に必要な情報を
> 構造化して出力するライブラリでGoパッケージとして提供しています。Go言語以外でもPHPやJava、Pythonや
> RustなどJSONを読める言語であれば、どのような環境においても解析結果を得ることでバウンスの発生状況を
> 捉えるのにとても有用です。

- [🇬🇧 **README**](README.md)
- [シシマイ? | What is Sisimai](#what-is-sisimai)
    - [主な特徴的機能 | The key features of Sisimai](#the-key-features-of-sisimai)
    - [コマンドラインでのデモ | command line demo](#command-line-demo)
- [シシマイを使う準備 | Setting Up Sisimai](#setting-up-sisimai)
    - [動作環境 | System requirements](#system-requirements)
    - [インストールとビルド | Install and Build](#install)
- [使い方 | Usage](#usage)
    - [基本的な使い方 | Basic usage](#basic-usage)
    - [解析結果をJSONで得る | Convert to JSON](#convert-to-json)
    - [コールバック機能 | Callback feature](#callback-feature)
    - [出力例 | Output example](#output-example)
- [Go言語版Sisimaiと他の言語版との違い](#differences-between-go-and-others)
- [Contributing](#contributing)
    - [バグ報告 | Bug report](#bug-report)
    - [解析できないメール | Emails could not be decoded](#emails-could-not-be-decoded)
- [その他の情報 | Other Information](#other-information)
    - [関連サイト | Related sites](#related-sites)
    - [参考情報 | See also](#see-also)
- [作者 | Author](#author)
- [著作権 | Copyright](#copyright)
- [ライセンス | License](#license)

What is sisimai
===================================================================================================
Sisimai(シシマイ)は複雑で多種多様なバウンスメールを解析してバウンスした理由や宛先メールアドレスなど
配信が失敗した結果を構造化データで出力するライブラリでJSONでの出力も可能です

![](https://libsisimai.org/static/images/figure/sisimai-overview-2.png)

The key features of Sisimai
---------------------------------------------------------------------------------------------------
* __バウンスメールを構造化したデータに変換__
  * 以下27項目の情報を含むデータ構造[^1]
    * __基本的情報__: `Timestamp`, `Origin`
    * __発信者情報__: `Addresser`, `SenderDomain`, 
    * __受信者情報__: `Recipient`, `Destination`, `Alias`
    * __配信の情報__: `Action`, `ReplyCode`, `DeliveryStatus`, `Command`
    * __エラー情報__: `Reason`, `DiagnosticCode`, `DiagnosticType`, `FeedbackType`, `FeedbacID`, `HardBounce`
    * __メール情報__: `Subject`, `MessageID`, `ListID`,
    * __その他情報__: `DecodedBy`, `TimezoneOffset`, `Lhost`, `Rhost`, `Token`, `Catch`, `Toxic`
  * __出力可能な形式__
    * struct ([sisimai/siba.Fact](https://github.com/sisimai/go-sisimai/blob/5-stable/siba/fact.go))
    * JSON ([`encoding/json`](https://pkg.go.dev/encoding/json)を使用)
* __インストールも使用も簡単__
  * `$ go get -u libsisimai.org/sisimai/v5@latest`
  * `import "libsisimai.org/sisimai/v5"`
* __高い解析精度__
  * [60種類のMTAs/MDAs/ESPs](https://libsisimai.org/en/engine/)に対応
  * Feedback Loop(ARF)にも対応
  * [34種類のバウンス理由](https://libsisimai.org/en/reason/)を検出

[^1]: コールバック機能を使用すると`Catch`フィールドの下に独自のデータを追加できます

Command line demo
---------------------------------------------------------------------------------------------------
次の画像のように、Go版シシマイ(go-sisimai)の`Dump()`関数を使うとコマンドラインから簡単に
バウンスメールを解析することができます。
![](https://libsisimai.org/static/images/demo/sisimai-5-cli-dump-g01.gif)

Setting Up Sisimai
===================================================================================================
System requirements
---------------------------------------------------------------------------------------------------
シシマイの動作環境についての詳細は[Sisimai | シシマイを使ってみる](https://libsisimai.org/ja/start/)
をご覧ください。

* [Go 1.24.0 or later](http://go.dev/dl/) (v5.4.0からGo 1.24以上が必要になりました)
* v5.2.1で標準モジュールを除く外部モジュール依存は無くなりました

Install
---------------------------------------------------------------------------------------------------
### Install
```shell
$ mkdir ./sisimai
$ cd ./sisimai
$ go mod init example.com/sisimaicli
go: creating new go.mod: module example.com/sisimaicli

$ go get -u libsisimai.org/sisimai/v5@latest
go: added libsisimai.org/sisimai/v5 v5.5.0

$ cat ./go.mod
module example.com/sisimaicli

go 1.25

require (
	libsisimai.org/sisimai/v5 v5.5.0 // indirect
)
```

### Build
例えば、以下のコードはバウンスメールを解析してその結果をそれぞれJSON文字列として出力する最小限の
プログラムです。**sisi**mai **d**ecoderで`sisid.go`としていますが、名前は何でもよいです。

```shell
$ vi ./sisid.go
```

```go:sisid.go
// sisimai decoder program
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai/v5"

func main() {
    path := os.Args[1]
    args := sisimai.Args()

    sisi, nyaan := sisimai.Rise(path, args)
    for _, e := range sisi {
        cv, _ := e.Dump()
        fmt.Printf("%s\n",cv)
    }
    if len(nyaan) > 0 { fmt.Frpintf(os.Stderr, "%v\n", nyaan) }
}
```

`sisid.go`が書けたら次のように`go build`コマンドで実行可能なバイナリを作ります。
```shell
$ CGO_ENABLED=0 go build -o ./sisid ./sisid.go
```

第一引数にバウンスメール(またはMaildir/)へのPATHを指定すると解析結果がJSON文字列で出力されます。
```shell
$ ./sisid ./path/to/bounce-mail.eml | jq
{
  "addresser": "michistuna@example.org",
  "recipient": "kijitora@google.example.com",
  "timestamp": 1650119685,
  "action": "failed",
  "alias": "contact@example.co.jp",
  "catch": null,
  "decodedby": "Postfix",
  "deliverystatus": "5.7.26",
  "destination": "google.example.com",
  "diagnosticcode": "host gmail-smtp-in.l.google.com[64.233.187.27] said: This mail has been blocked because the sender is unauthenticated. Gmail requires all senders to authenticate with either SPF or DKIM. Authentication results: DKIM = did not pass SPF [relay3.example.com] with ip: [192.0.2.22] = did not pass For instructions on setting up authentication, go to https://support.google.com/mail/answer/81126#authentication c2-202200202020202020222222cat.127 - gsmtp (in reply to end of DATA command)",
  "diagnostictype": "SMTP",
  "feedbackid": "",
  "feedbacktype": "",
  "hardbounce": false,
  "lhost": "relay3.example.com",
  "listid": "",
  "messageid": "hwK7pzjzJtz0RF9Y@relay3.example.com",
  "origin": "./path/to/bounce-mail.eml",
  "reason": "authfailure",
  "rhost": "gmail-smtp-in.l.google.com",
  "replycode": "550",
  "command": "DATA",
  "senderdomain": "google.example.com",
  "subject": "Nyaan",
  "timezoneoffset": "+0900",
  "token": "5253e9da9dd67573851b057a89cbcf41293e99bf",
  "toxic": false
}
```

Usage
===================================================================================================
Basic usage
---------------------------------------------------------------------------------------------------
以下のように`libsisimai.org/sisimai.Rise()`関数にバウンスメールへのPATHを渡して呼び出すと解析結果が
[`[]siba.Fact`](https://github.com/sisimai/go-sisimai/blob/5-stable/siba/fact.go)構造体として、発生
したエラーが[`[]siba.NotDecoded`](https://github.com/sisimai/go-sisimai/blob/5-stable/siba/not-decoded.go)
構造体としてそれぞれ得られます。

```go
import "os"
import "fmt"
import "libsisimai.org/sisimai/v5"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // siba.DecodingArgs{}

    // バウンス理由が"delivered"になった結果も必要ならargs.Deliveredにtrueを入れる
    args.Delivered = true

    // バウンス理由が"vacation"になった結果も必要ならargs.Vacationにtrueを入れる
    args.Vacation  = true

    // sisiは[]siba.Fact構造体スライス
    sisi, nyaan := sisimai.Rise(path, args)
    if len(sisi) > 0 {
        for _, e := range sisi {
            // e is a siba.Fact struct
            fmt.Printf("- Sender is %s\n", e.Addresser.Address)
            fmt.Printf("- Recipient is %s\n", e.Recipient.Address)
            fmt.Printf("- Bounced due to %s\n", e.Reason)
            fmt.Printf("- With the error message: %s\n", e.DiagnosticCode)

            cv, _ := e.Dump()     // 解析結果ごとにJSON文字列化する
            fmt.Printf("%s\n",cv) // jqコマンドで読めるJSON文字列を出力する
        }
    }
    // nyaanは[]siba.NotDecoded構造体スライス
    if len(nyaan) > 0 { fmt.Fprintf(os.Stderr, "%v\n", nyaan) }
}
```

Convert to JSON
---------------------------------------------------------------------------------------------------
下記のようにlibsisimai.org/sisimai.Dump()関数を、mboxかMaildir/のPATHを引数にして実行すると解析結果
が文字列(JSON)で返ってきます。

```go
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai/v5"

func main() {
    path := os.Args[1]
    args := sisimai.Args()

    json, nyaan := sisimai.Dump(path, args)
    if json != nil && *json != "" { fmt.Printf("%s\n", *json) }
    if len(nyaan) > 0 { fmt.Fprintf(os.Stderr, "%v\n", nyaan) }
}
```

Callback feature
---------------------------------------------------------------------------------------------------
`sisimai.Rise()`または`sisimai.Dump()`の`args.Callback0`と`args.Callback1`は、解析前および解析後に
実行したいコールバック関数を入れる項目です。
`args.Callback0`は`sisimai/message.sift()`で呼び出される関数でメールヘッダと本文に対して行う処理を、
`args.Callback1`は解析対象のメールファイルに対して行う処理をそれぞれ入れます。

コールバック関数0(`args.Callback0`)で処理した結果は`siba.Fact.Catch`を通して得られます。

### Callback0: メールヘッダと本文に対して
`args.Callback0`は`sisimai/message.sift()`で呼び出される関数でメールヘッダと本文に対して行う処理を
定義します。

```perl
package main
import "os"
import "fmt"
import "strings"
import "libsisimai.org/sisimai/v5"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // siba.DecodingArgs{}

    args.Callback0 = func(arg *sisimai.CallbackArg0) (map[string]interface{}, error) {
        // - この関数は解析処理実行前に呼び出される
        // - 例えば元メールにある"X-Delivery-App-ID:"ヘッダーの値を取り出してdata["x-delivery-app-id"]に入れる
        // - dataに入れた値はsiba.Fact構造体のCatchを通して型アサーションを経て参照可能
        name := "X-Delivery-App-ID"
        data := make(map[string]interface{})
        data[strings.ToLower(name)] = ""

        if arg.Payload != nil && len(*arg.Payload) > 0 {
            mesg := *arg.Payload
            if p0 := strings.Index(mesg, "\n" + name + ":"); p0 > 0 {
                cw := p0 + len(name) + 2
                p1 := strings.Index(mesg[cw:], "\n")
                if p1 > 0 {
                    data[strings.ToLower(name)] = mesg[cw + 1:cw + p1]
                }
            }
        }
        return data, nil
    }

    sisi, _ := sisimai.Rise(path, args)
    if len(sisi) > 0 {
        for _, e := range sisi {
            // eはsiba.Fact構造体
            re, as := e.Catch.(map[string]interface{})
            if as == false { continue }
            if ca, ok := re["x-delivery-app-id"].(string); ok {
                fmt.Printf("- Catch[X-Delivery-App-ID] = %s\n", ca)
            }
        }
    }
}
```

### Callback1: 各メールのファイルに対して
`sisimai.Rise()`または`sisimai.Dump()`関数に渡せる`args.Callback1`に入れたコールバック関数は、解析
が終わった後に対象のメールファイルごとに呼び出されます。

```go
package main
import "os"
import "io/ioutil"
import "libsisimai.org/sisimai/v5"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // siba.DecodingArgs{}

    args.Callback1 = func(arg *sisimai.CallbackArg1) (bool, error) {
        // - 解析対象になったメールのファイルごとに呼び出される
        // - 例えば解析したメールの中身を/tmpにファイルとして保存するなど
        if nyaan := ioutil.WriteFile("/tmp/copy.eml", []byte(*arg.Payload), 0400); nyaan != nil {
            return false, nyaan
        }
        return true, nil
    }

    // sisiは[]siba.Fact構造体スライス
    sisi, nyaan := sisimai.Rise(path, args)
    if len(sisi) > 0 {
        for _, e := range sisi {
            // e is a siba.Fact struct
            ...
        }
    }
}
```

コールバック機能のより詳細な使い方は
[Sisimai | 解析方法 - コールバック機能](https://libsisimai.org/ja/usage/#callback)をご覧ください。

Output example
---------------------------------------------------------------------------------------------------
```json
[
  {
    "addresser": "michistuna@example.org",
    "recipient": "kijitora@google.example.com",
    "timestamp": 1650119685,
    "action": "failed",
    "alias": "contact@example.co.jp",
    "catch": null,
    "decodedby": "Postfix",
    "deliverystatus": "5.7.26",
    "destination": "google.example.com",
    "diagnosticcode": "host gmail-smtp-in.l.google.com[64.233.187.27] said: This mail has been blocked because the sender is unauthenticated. Gmail requires all senders to authenticate with either SPF or DKIM. Authentication results: DKIM = did not pass SPF [relay3.example.com] with ip: [192.0.2.22] = did not pass For instructions on setting up authentication, go to https://support.google.com/mail/answer/81126#authentication c2-202200202020202020222222cat.127 - gsmtp (in reply to end of DATA command)",
    "diagnostictype": "SMTP",
    "feedbackid": "",
    "feedbacktype": "",
    "hardbounce": false,
    "lhost": "relay3.example.com",
    "listid": "",
    "messageid": "hwK7pzjzJtz0RF9Y@relay3.example.com",
    "origin": "./path/to/bounce-mail.eml",
    "reason": "authfailure",
    "rhost": "gmail-smtp-in.l.google.com",
    "replycode": "550",
    "command": "DATA",
    "senderdomain": "google.example.com",
    "subject": "Nyaan",
    "timezoneoffset": "+0900",
    "token": "5253e9da9dd67573851b057a89cbcf41293e99bf",
    "toxic": false
  }
]
```

Differences between Go and Others
===================================================================================================
Go版シシマイと他の言語版([p5-sisimai](https://github.com/sisimai/p5-sisimai/tree/5-stable)と
[rb-sisimai](https://github.com/sisimai/rb-sisimai/tree/5-stable))には以下のような違いがありますが、
解析精度や対応しているMTAやメール形式は同じです。

Features
---------------------------------------------------------------------------------------------------
| 機能                                    | Go              | Perl              | Ruby  / JRuby   |
|-----------------------------------------|-----------------|-------------------|-----------------|
| 動作環境                                | 1.24 -          | 5.26 -            | 2.4 - / 9.2 -   |
| 依存モジュール数(標準パッケージを除く)  | **0**           | 2 モジュール      | 1 gem           |
| 対応している文字コード                  | **UTF-8のみ**   | UTF-8と他[^2]     | UTF-8と他[^3]   |
| ソースコードの行数                      | 9,200 行        | 9,990 行          | 9,970 行        |
| テスト件数                              | 255,000 件      | 340,000 件        | 240,000 件      |
| 1秒間に解析できるバウンスメール数[^4]   | 2900 通         | 750 通            | 620 通          |
| ライセンス                              | 二条項BSD       | 二条項BSD         | 二条項BSD       |
| 開発会社による商用サポート              | 提供中          | 提供中            | 提供中          |

[^2]: `Encode`と`Encode::Guess`に対応している文字コード
[^3]: `String#encode`メソッドが解釈できる文字コード
[^4]: macOS Monterey/1.6GHz Dual-Core Intel Core i5/16GB-RAM/Go 1.22/Perl 5.30/Ruby 2.6.4

Contributing
===================================================================================================
Bug report
---------------------------------------------------------------------------------------------------
もしもSisimaiにバグを発見した場合は[Issues](https://github.com/sisimai/go-sisimai/issues)にて連絡を
いただけると助かります。

Emails could not be decoded
---------------------------------------------------------------------------------------------------
Sisimaiで解析できないバウンスメールは
[set-of-emails/to-be-debugged-because/sisimai-cannot-parse-yet](https://github.com/sisimai/set-of-emails/tree/master/to-be-debugged-because/sisimai-cannot-parse-yet)リポジトリに追加してPull-Requestを送ってください。


Other Information
===================================================================================================
Related sites
---------------------------------------------------------------------------------------------------
* __@libsisimai__ | [Sisimai on Twitter (@libsisimai)](https://twitter.com/libsisimai)
* __LIBSISIMAI.ORG__ | [SISIMAI | MAIL ANALYZING INTERFACE | DECODING BOUNCES, BETTER AND FASTER.](https://libsisimai.org/)
* __Sisimai Blog__ | [blog.libsisimai.org](http://blog.libsisimai.org/)
* __Facebook Page__ | [facebook.com/libsisimai](https://www.facebook.com/libsisimai/)
* __GitHub__ | [github.com/sisimai/go-sisimai](https://github.com/sisimai/go-sisimai)
* __Perl version__ | [Perl version of Sisimai](https://github.com/sisimai/p5-sisimai)
* __Ruby version__ | [Ruby version of Sisimai](https://github.com/sisimai/rb-sisimai)
* __Fixtures__ | [set-of-emails - Sample emails for "make test"](https://github.com/sisimai/set-of-emails)

See also
---------------------------------------------------------------------------------------------------
* [README.md - README.md in English(🇬🇧)](https://github.com/sisimai/go-sisimai/blob/5-stable/README.md)
* [RFC3463 - Enhanced Mail System Status Codes](https://tools.ietf.org/html/rfc3463)
* [RFC3464 - An Extensible Message Format for Delivery Status Notifications](https://tools.ietf.org/html/rfc3464)
* [RFC3834 - Recommendations for Automatic Responses to Electronic Mail](https://tools.ietf.org/html/rfc3834)
* [RFC5321 - Simple Mail Transfer Protocol](https://tools.ietf.org/html/rfc5321)
* [RFC5322 - Internet Message Format](https://tools.ietf.org/html/rfc5322)

Author
===================================================================================================
[@azumakuniyuki](https://twitter.com/azumakuniyuki) and sisimai development team

Copyright
===================================================================================================
Copyright (C) 2014-2025 azumakuniyuki and sisimai development team, All Rights Reserved.

License
===================================================================================================
This software is distributed under The BSD 2-Clause License.

