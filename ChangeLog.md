RELEASE NOTES for the Go version of Sisimai
===================================================================================================
- releases: "https://github.com/sisimai/go-sisimai/releases"
- document: "https://libsisimai.org/"
- packages: "https://pkg.go.dev/libsisimai.org/sisimai/v5"

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
    - Implement `smtp/reply.AssosiatedWith()`
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

