RELEASE NOTES for the Go version of Sisimai
===================================================================================================
- releases: "https://github.com/sisimai/go-sisimai/releases"
- document: "https://libsisimai.org/"

v5.2.0p3
---------------------------------------------------------------------------------------------------
- release: ""
- version: ""
- changes:
  - #64 NTT DOCOMO (Major Japanese mobile carrier) no longer returns a bounce mail due to domain
    rejection or similar email settings, but instead were being delivered to the spam folder after
    March 13th. #70
  - Implement `rhost/for-cloudflare.go` for Cloudflare Email Routing #66
  - Implement `libsisimai.org/sisimai.Reason()` function #71
  - Implement test codes for `rhost/for-*.go` #41 #62
  - `string` package has been renamed to `moji` #68 #69
  - Memory assingment improvements #72 #73

v5.2.0 - The first release
---------------------------------------------------------------------------------------------------
- release: "Tue, 25 Feb 2025 10:48:25 +0900 (JST)"
- version: "5.2.0"
- changes:
  - The first release of the Go version of sisimai
  - The Go version of sisimai available at `go get libsisimai.org/sisimai`

