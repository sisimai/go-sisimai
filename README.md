![](https://libsisimai.org/static/images/logo/sisimai-x01.png)
[![License](https://img.shields.io/badge/license-BSD%202--Clause-orange.svg)](https://github.com/sisimai/go-sisimai/blob/5-stable/LICENSE)

> [!CAUTION]
> The Go language version of sisimai is currently in **Public Beta**.
> The API of `libsisimai.org/sisimai` and internal specifications are subject to significant
> changes until the official v5.2.0 release.

> [!NOTE]
> Sisimai is a Go package but it can be used in any environment that JSON can be read, such as PHP,
> Java, Python, and Rust. By obtaining the analysis results, it is very useful for understanding the
> bounce occurrence status. 

- [**README-JA(日本語)**](README-JA.md)
- [What is Sisimai](#what-is-sisimai)
    - [The key features of Sisimai](#the-key-features-of-sisimai)
    - [Command line demo](#command-line-demo)
- [Setting Up Sisimai](#setting-up-sisimai)
    - [System requirements](#system-requirements)
    - [Install and Build](#install)
- [Usage](#usage)
    - [Basic usage](#basic-usage)
    - [Convert to JSON](#convert-to-json)
    - [Callback feature](#callback-feature)
    - [Output example](#output-example)
- [Differences between Go vs. Others](#differences-between-go-and-others)
- [Contributing](#contributing)
    - [Bug report](#bug-report)
    - [Emails could not be decoded](#emails-could-not-be-decoded)
- [Other Information](#other-information)
    - [Related sites](#related-sites)
    - [See also](#see-also)
- [Author](#author)
- [Copyright](#copyright)
- [License](#license)


What is Sisimai
===================================================================================================
Sisimai is a Go package, is a library that decodes complex and diverse bounce emails and outputs the
results of the delivery failure, such as the reason for the bounce and the recipient email address,
in structured data. It is also possible to output in JSON format.

![](https://libsisimai.org/static/images/figure/sisimai-overview-2.png)

The key features of Sisimai
---------------------------------------------------------------------------------------------------
* __Decode email bounces to structured data__
  * Sisimai provides detailed insights into bounce emails by extracting 26 key data points.[^1]
    * __Essential information__: `Timestamp`, `Origin`
    * __Sender information__: `Addresser`, `SenderDomain`, 
    * __Recipient information__: `Recipient`, `Destination`, `Alias`
    * __Delivery information__: `Action`, `ReplyCode`, `DeliveryStatus`, `Command`
    * __Bounce details__: `Reason`, `DiagnosticCode`, `DiagnosticType`, `FeedbackType`, `FeedbackID`, `HardBounce`
    * __Message details__: `Subject`, `MessageID`, `ListID`,
    * __Additional information__: `DecodedBy`, `TimezoneOffset`, `Lhost`, `Rhost`, `Token`, `Catch`
  * Output formats
    * struct ([sisimai/sis.Fact](https://github.com/sisimai/go-sisimai/blob/5-stable/sis/fact.go))
    * JSON (by using [`encoding/json`](https://pkg.go.dev/encoding/json))
* __Easy to Install, Use.__
  * `$ go get -u libsisimai.org/sisimai@latest`
  * `import "libsisimai.org/sisimai"`
* __High Precision of Analysis__
  * Support [58 MTAs/MDAs/ESPs](https://libsisimai.org/en/engine/)
  * Support Feedback Loop Message(ARF)
  * Can detect [36 bounce reasons](https://libsisimai.org/en/reason/)

[^1]: The callback function allows you to add your own data under the `Catch` field.


Command line demo
---------------------------------------------------------------------------------------------------
The following screen shows a demonstration of `Dump` function of libsimai.org/sisimai package at
the command line using Go(go-sisimai) and `jq` command.
![](https://libsisimai.org/static/images/demo/sisimai-5-cli-dump-g01.gif)


Setting Up Sisimai
===================================================================================================
System requirements
---------------------------------------------------------------------------------------------------
More details about system requirements are available at
[Sisimai | Getting Started](https://libsisimai.org/en/start/) page.

* [Go 1.17.0 or later](http://go.dev/dl/)
* [golang.org/x/text/encoding](https://pkg.go.dev/golang.org/x/text/encoding)
* [golang.org/x/net/html/charset](https://pkg.go.dev/golang.org/x/net/html/charset)

Install and Build
---------------------------------------------------------------------------------------------------
### Install
```shell
$ mkdir ./sisimai
$ cd ./sisimai
$ go mod init example.com/sisimaicli
go: creating new go.mod: module example.com/sisimaicli

$ go get -u libsisimai.org/sisimai@latest
go: added golang.org/x/net v0.35.0
go: added golang.org/x/text v0.22.0
go: added libsisimai.org/sisimai v0.0.1

$ cat ./go.mod
module example.com/sisimaicli

go 1.20

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	libsisimai.org/sisimai v0.0.3 // indirect
)
```

### Build
For example, the following `sisid.go`: **sisi**mai **d**ecoder is a minimal program that decodes
bounce emails and outputs them in JSON format.

```shell
$ vi ./sisid.go
```

```go:sisid.go
// sisimai decoder program
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai"

func main() {
    path := os.Args[1]
    args := sisimai.Args()

    sisi, nyaan := sisimai.Rise(path, args)
    for _, e := range *sisi {
        cv, _ := e.Dump()
        fmt.Printf("%s\n",cv)
    }
    if len(*nyaan) > 0 { fmt.Frpintf(os.Stderr, "%v\n", *nyaan) }
}
```

Once you have written `sisid.go`, build an executable binary with `go build` command.
```shell
$ CGO_ENABLED=0 go build -o ./sisid ./sisid.go
```

Specifying the path to a bounce email (or Maildir/) as the first argument will output the decoded 
results as a JSON string.
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
  "token": "5253e9da9dd67573851b057a89cbcf41293e99bf"
}
```

Usage
===================================================================================================
Basic usage
---------------------------------------------------------------------------------------------------
`libsisimai.org/sisimai.Rise()` function provides the feature for getting decoded data as
[`*[]sis.Fact`](https://github.com/sisimai/go-sisimai/blob/5-stable/sis/fact.go) struct, occurred
errors as [`*[]sis.NotDecoded`](https://github.com/sisimai/go-sisimai/blob/5-stable/sis/not-decoded.go)
from bounced email messages as the following.

```go
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // sis.DecodingArgs{}

    // If you also need analysis results that are "delivered" (successfully delivered),
    // set `true` into the "Delivered" option for the Rise() function as shown below.
    args.Delivered = true

    // If you also need analysis results that show a "vacation" reason, set `true` into
    // the "Vacation" option for the Rise() function as shown in the following code.
    args.Vacation  = true

    // sisi is a pointer to []sis.Fact
    sisi, nyaan := sisimai.Rise(path, args)
    if len(*sisi) > 0 {
        for _, e := range *sisi {
            // e is a sis.Fact struct
            fmt.Printf("- Sender is %s\n", e.Addresser.Address)
            fmt.Printf("- Recipient is %s\n", e.Recipient.Address)
            fmt.Printf("- Bounced due to %s\n", e.Reason)
            fmt.Printf("- With the error message: %s\n", e.DiagnosticCode)

            cv, _ := e.Dump()     // Convert the decoded data to JSON string
            fmt.Printf("%s\n",cv) // JSON formatted string the jq command can read
        }
    }
    // nyaan is a pointer to []sis.NotDecoded
    if len(*nyaan) > 0 { fmt.Fprintf(os.Stderr, "%v\n", *nyaan) }
}
```

Convert to JSON
---------------------------------------------------------------------------------------------------
The following code snippet illustrates the use of the `libsisimai.org/sisimai.Dump()` function to
obtain decoded bounce email data in JSON array format.

```go
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai"

func main() {
    path := os.Args[1]
    args := sisimai.Args()

    json, nyaan := sisimai.Dump(path, args)
    if json != nil && *json != "" { fmt.Printf("%s\n", *json)   }
    if len(*nyaan) > 0 { fmt.Fprintf(os.Stderr, "%v\n", *nyaan) }
}
```

Callback feature
---------------------------------------------------------------------------------------------------
[`sis.DecodingArgs`](https://github.com/sisimai/go-sisimai/blob/5-stable/sis/decoding-args.go) have
the `Callback0` and `Callback1` fields for keeping callback functions. The former is called at
`message.sift()` for dealitng email headers and entire message body. The latter is called at the
end of each email file processing inside of `sisimai.Rise()`.

The results generated by the callback functions are accessible via `Catch` field defined in
[sis.Fact](https://github.com/sisimai/go-sisimai/blob/5-stable/sis/fact.go).

### Callback0: For email headers and the body
The function set in args.Callback0 is called at `message.sift()`.

```go
package main
import "os"
import "fmt"
import "strings"
import "libsisimai.org/sisimai"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // sis.DecodingArgs{}

    args.Callback0 = func(arg *sisimai.CallbackArg0) (map[string]interface{}, error) {
        // - This function allows users to add custom processing to the email before parsing.
        // - For example, you can extract the delivery ID from the "X-Delivery-App-ID:" header
        //   and store it in the data map like this: data["x-delivery-app-id"] = "neko22-2".
        // - The library executes this function and assigns the return value to the Catch field of the Fact struct.
        // - Users can then retrieve and access the data from Catch by type assertion in the caller.
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
    if len(*sisi) > 0 {
        for _, e := range *sisi {
            // e is a sis.Fact struct
            re, as := e.Catch.(map[string]interface{})
            if as == false { continue }
            if ca, ok := re["x-delivery-app-id"].(string); ok {
                fmt.Printf("- Catch[X-Delivery-App-ID] = %s\n", ca)
            }
        }
    }
}
```

### Callback1: For each email file
The function set in `args.Callback1` is called at `sisimai.Rise()` function for dealing each email
file after decoding each bounce message.

```go
package main
import "os"
import "io/ioutil"
import "libsisimai.org/sisimai"

func main() {
    path := os.Args[1]     // go run ./sisid /path/to/mailbox or maildir/
    args := sisimai.Args() // sis.DecodingArgs{}

    args.Callback1 = func(arg *sisimai.CallbackArg1) (bool, error) {
        // - This function defines custom operations that the user wants to perform on the parsed email file.
        // - For example, you can write the contents of the parsed bounce email to a file in /tmp/.
        if nyaan := ioutil.WriteFile("/tmp/copy.eml", []byte(*arg.Mail), 0400); nyaan != nil {
            return false, nyaan
        }
        return true, nil
    }

    // sisi is a pointer to []sis.Fact
    sisi, nyaan := sisimai.Rise(path, args)
    if len(*sisi) > 0 {
        for _, e := range *sisi {
            // e is a sis.Fact struct
            ...
        }
    }
}
```

More information about the callback feature is available at
[Sisimai | How To Parse - Callback](https://libsisimai.org/en/usage/#callback) Page.

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
    "token": "5253e9da9dd67573851b057a89cbcf41293e99bf"
  }
]
```

Differences between Go and Others
===================================================================================================
The following table show the differences between the Go version of Sisimai and the other language 
versions: [p5-sisimai](https://github.com/sisimai/p5-sisimai/tree/5-stable) and
[rb-sisimai](https://github.com/sisimai/rb-sisimai/tree/5-stable).

Features
---------------------------------------------------------------------------------------------------
| Features                                     | Go             | Perl            | Ruby  / JRuby |
|----------------------------------------------|----------------|-----------------|---------------|
| System requirements                          | 1.17 -         | 5.26 -          | 2.4 - / 9.2 - |
| Dependencies (Except standard libs)          | 2 packages     | 2 modules       | 1 gem         |
| Supported character sets                     | UTF-8 only     | UTF-8,etc. [^2] | UTF-8,etc.[^3]|
| Source lines of code                         | 9,400 lines    | 9,900 lines     | 9,800 lines   |
| The number of tests                          | 141,200 tests  | 320,000 tests   | 410,000 tests |
| The number of bounce emails decoded/sec [^4] | 1200 emails    | 450 emails      | 340 emails    |
| License                                      | 2 Clause BSD   | 2 Caluse BSD    | 2 Clause BSD  |
| Commercial support                           | Coming soon    | Available       | Available     |

[^2]: Character sets supported by `Encode` and `Encode::Guess` modules
[^3]: Character sets supported by `String#encode` method
[^4]: macOS Monterey/1.6GHz Dual-Core Intel Core i5/16GB-RAM/Go 1.22/Perl 5.30

Contributing
===================================================================================================
Bug report
---------------------------------------------------------------------------------------------------
Please use the [issue tracker](https://github.com/sisimai/go-sisimai/issues) to report any bugs.

Emails could not be decoded
---------------------------------------------------------------------------------------------------
Bounce emails that couldn't be decoded by the latest version of sisimai are saved in the repository
[set-of-emails/to-be-debugged-because/sisimai-cannot-parse-yet](https://github.com/sisimai/set-of-emails/tree/master/to-be-debugged-because/sisimai-cannot-parse-yet). 
If you have found any bounce email cannot be decoded using sisimai, please add the email into the
directory and send Pull-Request to this repository.


Other Information
===================================================================================================
Related sites
---------------------------------------------------------------------------------------------------
* __@libsisimai__ | [Sisimai on Twitter (@libsisimai)](https://twitter.com/libsisimai)
* __LIBSISIMAI.ORG__ | [SISIMAI | MAIL ANALYZING INTERFACE | DECODING BOUNCES, BETTER AND FASTER.](https://libsisimai.org/)
* __Sisimai Blog__ | [blog.libsisimai.org](http://blog.libsisimai.org/)
* __Facebook Page__ | [facebook.com/libsisimai](https://www.facebook.com/libsisimai/)
* __GitHub__ | [github.com/sisimai/go-sisimai](https://github.com/sisimai/go-sisimai)
* __Perl verison__ | [Perl version of Sisimai](https://github.com/sisimai/p5-sisimai)
* __Ruby version__ | [Ruby version of Sisimai](https://github.com/sisimai/rb-sisimai)
* __Fixtures__ | [set-of-emails - Sample emails for "make test"](https://github.com/sisimai/set-of-emails)

See also
---------------------------------------------------------------------------------------------------
* [README-JA.md - README.md in Japanese(日本語)](https://github.com/sisimai/go-sisimai/blob/5-stable/README-JA.md)
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

