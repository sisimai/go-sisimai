RELEASE NOTES for the Go version of Sisimai
===================================================================================================
- releases: "https://github.com/sisimai/go-sisimai/releases"
- document: "https://libsisimai.org/"
- packages: "https://pkg.go.dev/libsisimai.org/sisimai/v5"

v5.?.?
---------------------------------------------------------------------------------------------------
- release: ""
- version: ""
- changes:
  - #336 #346 Improvement in error message patterns.
    - Update error message patterns in `AuthFailure`, `ContentError`, `NoRelaying`, `UserUnknown`,
      `Rejected`, `SystemError`, `Suspend`, and `rhost/microsoft`.
    - Remove error message patterns already defined in `reason/` from some MTA moudles of `lhost/`.
    - Move error message patterns from some MTA modules of `lhost/` to `reason/`.
  - #337 Bug fix: The value of command should be `RCPT` when `RCPT first` in the error message.
  - #347 Update the order of `classorder` in `reason/lib.go`.
  - #350 Use a switch statement instead of if-else for better readability.
  - #355 List all the Zoho domains.
  - #361 Sisimai can be built with Go 1.26.0.
  - #358 #360 #365 Tiny code improvements around `append`, `strings.Trim` function.
  - #367 Code improvement for detecting bounce messages returned from Google Groups.
  - #369 #370 Fix and update comments in `smtp/status/lib.go`.
  - #372 Fix typo in the sample code at README files. thanks to @jcbf
  - #374 Fix bug in the boolean logic for lines starting with space or tab. Thanks to @jcbf
  - #388 `rfc791.FindIPv4Address()` is 3x faster.
  - #390 `rfc791.IsIPv4Address()` is 6x faster.

v5.6.0
---------------------------------------------------------------------------------------------------
- release: "Mon,  2 Feb 2026 18:30:22 +0900 (JST)"
- version: "5.6.0"
- changes:
  - **Changes in Bounce Reason Categorization**
    - #303 Update bounce status mappings in `smtp/status/lib.go`.
    - #307 Consolidate four bounce reasons into two:
      - #308 Merge `TooManyConn` and `Speeding` into `RateLimited`.
      - #309 Merge `ExceedLimit` and `MesgTooBig` into `EmailTooLarge`.
    - #325 Update the internal constant name for some reason names.
    - #327 Update assigned reasons in `rhost/for-*.go`.
    - #329 Merge similar error message patterns and remove ambiguous ones.
    - #333 Attachment-related errors have been moved from `PolicyViolation` to `ContentError`.
  - **Reorganize the internal status code** #322 
    - Temporary error `4.0.9**` has been changed to `4.9.***`.
    - Permanent error `5.0.9**` has been changed to `5.9.***`.
  - #319 `moji.Token()` has been moved to `fact/lib.go`.
  - Happy Birthday to Suzu (formerly known as "Neko-dono" Michitsuna).

v5.5.0
---------------------------------------------------------------------------------------------------
- release: "Fri,  5 Dec 2025 17:25:22 +0900 (JST)"
- version: "5.5.0"
- changes:
  - #173 Add some SMTP reply codes that uniquely identify a bounce reason at `smtp/reply/lib.go`.
  - #174 #186 Implement `IsAmbiguous()` function at `smtp/status/lib.go`.
  - #185 Sisimai partially supports the media types described in RFC6533 such as `message/global`.
  - #189 #190 Refactor: Remove redundant header normalization code blocks in `rfc2045` package.
  - #203 #204 Fix the index out of range bug in multipart blocks.
  - #223 #226 The following MTA modules have been updated for decoding more bounce emails:
    - #153 #200 `lhost/via-mimecast.go` for decoding bounce mails from Mimecast.
    - #209 #214 `lhost/via-mfilter.go`
    - #211 #216 `lhost/via-interscanmss.go` has been renamed to `lhost/via-trendmicro.go`. #233
    - #217 #218 `lhost/via-postfix.go`
    - #210 #215 `lhost/via-x1.go`
    - #219 #222 `lhost/via-x2.go`
    - #234 #241 `lhost/via-mailmarshalsmtp.go` has been renamed to `via-mailmarshal.go`.
  - #229 Add new error message pattern for Spamhaus in `reason/why-blocked.go`.
  - #230 #238 Package `sis` has been renamed to `siba`: Sisimai Internal Bounce Abstraction.
  - #202 #243 Detect `Suspend` reason from a auto replied message by `rfc3834/lib.go`.
  - #242 Fixed an issue where auto-reply messages were not decoded by `Rise()` function when the
    `vacation` option was specified.
  - #247 #248 #250 #251 #255 Fix wrong reason names and values of the `Action:` field.
  - #244 #245 #257 Use constants to define SMTP commands and bounce reason names, and Action values.
  - #249 `Message-ID` related errors are classified as `NotCompliantRFC`.
  - #261 #262 #264 #273 Code, test, and coverages improved.
  - #263 Fix bug in `IsEmailAddress()` function of `rfc5322/address.go`.
  - #265 Fix bug: wrong `groupindex` value in `address/find.go`.
  - #279 #28 Check that the input file or MIME decoded part is a binary or not.
  - #283 #285 #288 #291 Update SMTP error codes and messages in the following `rhost` packages:
    - `for-apple.go` `inactive email address` indicates `Suspend`
    - `for-facebook.go` Update `AuthFailure` and `MailboxFull`
    - `for-franceptt.go` Update `AuthFailure`, `Blocked`, `Rejected`, and `TooManyConn`
    - `for-godaddy.go` Add `TooManyConn`
    - `for-messagelabs.go` Add many error message patterns
    - `for-tencent.go` Add `Speeding`
    - `for-microsoft.go` Implement SMTP error codes described in Outlook Postmaster/Troubleshooting
    - `for-zoho.go` Impelment as a new module
  - #293 Sisimai can be built with Go 1.25.4.
  - #296 Some error message patterns in `PolicyViolation` have been moved to `Rejected`.
  - #299 EXPERIMENTAL: Implement `Toxic` field in the decoded results.

v5.4.1
---------------------------------------------------------------------------------------------------
- release: "Sun, 31 Aug 2025 09:22:25 +0900 (JST)"
- version: "5.4.1"
- changes:
  - #155 #161 Support a bounce mail returned from `privaterelay.appleid.com`.
  - #158 #160 Update SMTP error and status codes of Gmail updated in August 2025.
    - Gmail SMTP errors and codes https://support.google.com/a/answer/3726730
    - https://github.com/azumakuniyuki/feb-2024-no-auth-no-entry/commit/364214227
    - #163 #164 Sisimai can be built with Go 1.25.

v5.4.0
---------------------------------------------------------------------------------------------------
- release: "Tue,  1 Jul 2025 20:22:22 +0900 (JST)"
- version: "5.4.0"
- changes:
  - **BREAKING CHANGES**
    - `sisimai.Rise()` function now return `[]sis.Fact` and `[]sis.NotDecoded` instead of pointers.
      Thanks to @corny #119 #145 #148
    - The minimum Go version required to run Sisimai is now **Go 1.24** #112 #127
      - #110 Fixed out of bounds read at `EmailEntity.setNewLine()` function in `mail/lib.go` using
        the build-in `min()` function. Thanks to @VolkerLieber
      - #116 #117 Use `strings.Cut()` instead of `strings.SplitN(v,s,2)` and `v[n:strings.Index(v,s)]`
      - #118 #122 Use `slices.Contains()` instead of `moji.EqualsAny()`
      - #123 #128 Use `strings.Lines()` instead of `strings.Split(v, "\n")`
  - #99 Use `error.Error()` instead of `fmt.Sprintf()` for stringify an error message
  - #101 #102 #103 Use `make()` to initialise a slice
  - #104 #152 Implement new status codes of Google: `4.7.40` and `5.7.32` as `authfailure`
  - #105 #106 Tiny code improvement in `moji.Select()` function
  - #107 #108 #149 #150 #151 Implement the new status code `5.7.515` as `authfailure` and other
    undocumented status codes begin with `4.4.` as `systemerror` of Microsoft in `sisimai/rhost`
  - #109 #115 Fix spell errors in some documents
  - #111 #113 Set a pointer to `sis.DecodingArgs` struct when the 2nd argument of `sisimai.Rise()`
    is `nil`. Thanks to @VolkerLieber
  - #120 #121 Use golangci-lint with the minimum linters. Thanks to @corny
  - #124 #125 `moji.Squeeze()` and `mail.setNewLine()` no longer returns any value
  - #131 #132 Use iota instead of `lhost.DeliveryStatus` hash map
  - #133 #134 use the `switch` statement without a condition instead of the infinite `for` loop and
    the `if` statement
  - #135 Use `moji.IsContained()` instead of `string.Contains()` in some loops
  - #136 #137 Implement `moji.AlignedAny()` function in `moji/any.go`
  - #139 #141 `reason.GetRetried` has been replaced with `reason.ShouldBeRetried()` function
  - #140 Tiny code improvements in `reason.IsExplicit()` function
  - #142 #143 Change the order of fields in some structs to improve memory alignment
  - #146 #147 Adopt idiomatic Go style by returning direct slices instead of pointers to slices in
    some functions except `sisimai.Rise()`

v5.3.0
---------------------------------------------------------------------------------------------------
- release: "Sat, 29 Mar 2025 05:13:32 +0900 (JST)"
- version: "5.3.0"
- changes:
  - **Corrected the completely and utterly broken module path.** The module path has been replaced
    with `libsisimai.org/sisimai/v5`. Thanks to @inboxsphere #96 #98
  - sisimai can be built with Go 1.24
  - Performance has improved by approximately 1.7x compared to v5.2.1.
    - Reduce struct and string copies, and the return types of several functions have been modified
      to return pointers to strings and structs #76 #80 #83 #84 #85 #86 #87 
    - #78 `fmt.Sprintf()` usage has been decreased to improve performance.
    - #88 #89 Optimize the order of MTA modules.
  - #75 Refactor comments for all the functions, follow the godoc as possible.
  - #79 Fix bug at the argument of `fmt.Sprintf()` in `rfc3834/lib.go`
  - #82 `sis.Beforefact.Empty()` has been renamed to `IsEmpty()`
  - #90 #92 SMTP reply code improvements
    - Update the list of SMTP status codes in `smtp/reply/lib.go`
    - Implement `smtp/reply.AssociatedWith()`
    - Implement `smtp/status.IsExplicit()`

v5.2.1
---------------------------------------------------------------------------------------------------
- release: "Wed, 12 Mar 2025 06:22:25 +0900 (JST)
- version: "5.2.1"
- changes:
  - #64 NTT DOCOMO (Major Japanese mobile carrier) no longer rejects an email message due to domain
    rejection or similar email settings, but instead were being delivered to the spam folder after
    March 13th. #70
  - Implement `rhost/for-cloudflare.go` for Cloudflare Email Routing #66
  - Implement `libsisimai.org/sisimai.Reason()` function #71
  - Implement test codes for `rhost/for-*.go` #41 #62
  - `string` package has been renamed to `moji` #68 #69
  - Memory assingment improvements #72 #73
  - Comments for all the functions refactored #74 #75
  - #42 Refactored to remove all dependencies on third-party modules in golang.org/x/*, except for
    the Go standard library, resulting in a 22% binary size reduction.

v5.2.0 - The first release
---------------------------------------------------------------------------------------------------
- release: "Tue, 25 Feb 2025 10:48:25 +0900 (JST)"
- version: "5.2.0"
- changes:
  - The first release of the Go version of sisimai
  - The Go version of sisimai available at `go get libsisimai.org/sisimai`

