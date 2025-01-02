// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _                      _      
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       ___ ___   __| | ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / __/ _ \ / _` |/ _ \
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| (_| (_) | (_| |  __/
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \___\___/ \__,_|\___|
import "testing"
import "os"
import "io"
import "fmt"
import "errors"
import "strings"
import "path/filepath"
import "sisimai/sis"
import "sisimai/rfc1123"
import "sisimai/smtp/command"
import sisimbox "sisimai/mail"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

type IsExpected struct {
	Label      string // "01" or "1025"
	Index      uint8  // 1,2,3,...
	Status     string // "5.1.1" or empty
	ReplyCode  string // "550" or empty
	Reason     string // "userunknown"
	HardBounce bool   // true or false
	AnotherOne string // "Feedback-Type" or other value
}
var SampleRoot = "set-of-emails/"
var PublicDirs = "maildir/bsd/"
var SecretDirs = "private/"
var Alternates = map[string][]string{
	"Exchange2007": []string{"Office365"},
	"Exim":         []string{"MailRu", "MXLogic"},
	"qmail":        []string{"X4", "Yahoo"},
	"RFC3464":      []string{
		"Amavis", "Aol", "AmazonWorkMail", "Barracuda", "Bigfoot", "Facebook", "GSuite", "McAfee",
		"MessageLabs", "Outlook", "PowerMTA", "ReceivingSES", "SendGrid", "SurfControl", "Yandex",
		"X5",
	},
}

var ArgForRise = map[string]bool {"delivered": true, "vacation": true}
var TestReturn = map[string]interface{}{"neko-dono": []string{"Michitsuna", "Suzu"}}
var CallbackFn = func(arg *sis.CallbackArgs) map[string]interface{} { return TestReturn }

// EngineTest() is called from lhost/*_test.go, rhost/*_test.go, rfc3464/lib_test.go, arf/lib_test.go.
func EngineTest(t *testing.T, enginename string, isexpected [][]IsExpected, publictest bool) {
	// @param    *testing.T     t          Test object
	// @param    string         enginename MTA module name such as "OpenSMTPD"
	// @param    [][]IsExpected isexpected The list of results
	// @param    bool           publictest false if set-of-emails/private
	cx := 0
	//functionsi := "sisimai/lhost.Inquire[" + enginename + "]"
	prefixpath := "../../" + SampleRoot

	if publictest == true {
		// Public samples are in set-of-emails/maildir/bsd/lhost-*.eml
		prefixpath += PublicDirs + "lhost-" + strings.ToLower(enginename)

	} else {
		// Private samples are in set-of-emails/private/lhost-* directory
		prefixpath += SecretDirs + "lhost-" + strings.ToLower(enginename) + "/"
	}
	if len(isexpected) == 0 { t.Skip() }

	for _, e := range isexpected {
		// The element is a list of []IsExpected
		t.Run("", func(t *testing.T) {
			// Set a path to sample email
			ee := fmt.Sprintf("%s[%4s-00]", enginename, e[0].Label)
			ef := ""; if publictest == true {
				// Try to stat(2) for the public sample email
				ef = fmt.Sprintf("%s-%s.eml", prefixpath, e[0].Label)
				if _, nyaan := os.Stat(ef); nyaan != nil {
					t.Errorf("%s cannot read the sample email: %s", ee, nyaan)
				}; cx++
			} else {
				// Try to find the email file path in the set-of-emails/private, the file exists as
				// a file name such as private/lhost-opensmtpd/1012-933ce597.eml
				match, nyaan  := filepath.Glob(prefixpath + e[0].Label + "-*.eml")
				if nyaan      != nil { t.Errorf("%s something wrong: %s", ee, nyaan) }; cx++
				if len(match) == 0   { t.Errorf("%s email not found: %s", ee, nyaan) }; cx++
				for _, f := range match { ef = f; break }
			}

			emailthing, nyaan := sisimbox.Rise(ef)
			if nyaan != nil { t.Errorf("%s failed to load the sample email: %s", ee, nyaan) }; cx++

			for {
				if mesg, nyaan := emailthing.Read(); nyaan != nil {
					// Failed to read the email
					if errors.Is(nyaan, io.EOF) {
						// sisimai has reached to the end of email/directory
						break

					} else {
						// Something wrong, sisimai failed to read the email as a text
						t.Errorf("%s failed to read the sample email: %s", ee, nyaan)
						continue
					}; cx++
				} else {
					// Read and decode each email file as a string
					if emailthing.Size == 0 { t.Errorf("%s %s is empty", ee, ef); continue }; cx++

					mesg = sisimoji.ToLF(mesg)
					fact, nyaan := Rise(mesg, emailthing.Path, ArgForRise, CallbackFn)
					if nyaan     != nil { t.Errorf("%s %s", ee, nyaan[0].Error()) }; cx++
					if len(fact) == 0   { t.Errorf("%s failed to decode any bounce message in %s", ee, ef) }; cx++

					for j, fs := range fact {
						// Compare each decoded element with each expected value
						if j < len(e) {
							ev := e[j]
							ee  = fmt.Sprintf("%s[%4s-%02d]", enginename, ev.Label, ev.Index)

							if fs.DeliveryStatus != ev.Status {
								// DeliveryStatus
								t.Errorf("%s Status is (%s) but (%s)", ee, fs.DeliveryStatus, ev.Status)
							}; cx++

							if fs.ReplyCode != ev.ReplyCode {
								// ReplyCode
								t.Errorf("%s ReplyCode is (%s) but (%s)", ee, fs.ReplyCode, ev.ReplyCode)
							}; cx++

							if fs.Reason != ev.Reason {
								// Reason
								t.Errorf("%s Reason is (%s) but (%s)", ee, fs.Reason, ev.Reason)
							}; cx++

							if fs.HardBounce != ev.HardBounce {
								// HardBounce
								t.Errorf("%s HardBounce is (%t) but (%t)", ee, fs.HardBounce, ev.HardBounce)
							}; cx++

							if ev.AnotherOne != "" {
								// AnotherOne
								if fs.Reason == "feedback" && fs.FeedbackType != ev.AnotherOne {
									// For example "FeedbackType"
									t.Errorf("%s FeedbackType is (%s) but (%s)", ee, fs.FeedbackType, ev.AnotherOne)
								}; cx++
							}

							if fs.DecodedBy != enginename {
								// DecodedBy
								altdecoder := enginename
								if len(Alternates[fs.DecodedBy]) > 0 {
									// The MTA module in sisimai/lhost is a removed module
									// https://github.com/sisimai/go-sisimai/issues/7
									for _, as := range Alternates[fs.DecodedBy] {
										if as == enginename { altdecoder = fs.DecodedBy; break }
									}
									if enginename == "AmazonSES" {
										// AmazonSES or RFC3464
										if fs.DecodedBy == "AmazonSES" || fs.DecodedBy == "RFC3464" { altdecoder = fs.DecodedBy }
									}
								}

								if fs.DecodedBy != altdecoder {
									t.Errorf("%s DecodedBy is (%s) but (%s)", ee, fs.DecodedBy, altdecoder)
								}
							}; cx++

							/* Other fields except above */
							// Action
							if fs.Action == "" { t.Errorf("%s Action is empty", ee) }; cx++
							if ActionList[fs.Action] == false {
								t.Errorf("%s Action (%s) is an invalid value", ee, fs.Action)
							}; cx++

							// Addresser
							if fs.Addresser.Address == "" { t.Errorf("%s Addresser.Address is empty", ee) }; cx++
							if fs.Addresser.Alias   != "" { t.Errorf("%s Addresser.Alias is (%s) not empty", ee, fs.Addresser.Alias) }; cx++
							if fs.Addresser.Verp    != "" { t.Errorf("%s Addresser.Verp is (%s) not empty", ee, fs.Addresser.Verp) }; cx++
							if fs.Addresser.User    == "" { t.Errorf("%s Addresser.User is empty", ee) }; cx++
							if fs.Addresser.Host    == "" { t.Errorf("%s Addresser.User is empty", ee) }; cx++
							if fs.Addresser.Host    != fs.SenderDomain {
								// SenderDomain
								t.Errorf("%s Addresser.Host is (%s) but (%s)", ee, fs.Addresser.Host, fs.SenderDomain)
							}; cx++
							if rfc1123.IsInternetHost(fs.Addresser.Host) == false {
								// Is not a valid internet hostname
								t.Errorf("%s Addresser.Host (%s) is not a valid internet hostname", ee, fs.Addresser.Host)
							}
							if sisiaddr.IsEmailAddress(fs.Addresser.Address) == false {
								// Is not a valid email address
								t.Errorf("%s Addresser.Address (%s) is not a valid email address", ee, fs.Addresser.Address)
							}; cx++

							// Alias, Recipient, Destination
							if fs.Recipient.Address == "" { t.Errorf("%s Recipient.Address is empty", ee) }; cx++
							if fs.Recipient.User    == "" { t.Errorf("%s Recipient.User is empty", ee) }; cx++
							if fs.Recipient.Host    == "" { t.Errorf("%s Recipient.User is empty", ee) }; cx++
							if fs.Recipient.Host    != fs.Destination { 
								// Destination
								t.Errorf("%s Recipient.Host is (%s) but (%s)", ee, fs.Recipient.Host, fs.Destination)
							}; cx++
							if rfc1123.IsInternetHost(fs.Recipient.Host) == false {
								// Is not a valid internet hostname
								t.Errorf("%s Recipient.Host (%s) is not a valid internet hostname", ee, fs.Recipient.Host)
							}
							if sisiaddr.IsEmailAddress(fs.Recipient.Address) == false {
								// Is not a valid email address
								t.Errorf("%s Recipient.Address (%s) is not a valid email address", ee, fs.Recipient.Address)
							}; cx++
							if fs.Recipient.Verp != "" && sisiaddr.IsEmailAddress(fs.Recipient.Verp) == false {
								// Is not a valid email address
								t.Errorf("%s Recipient.Verp (%s) is not a valid email address", ee, fs.Recipient.Verp)
							}; cx++
							if fs.Recipient.Alias != "" && sisiaddr.IsEmailAddress(fs.Recipient.Alias) == false {
								// Is not a valid email address
								t.Errorf("%s Recipient.Alias (%s) is not a valid email address", ee, fs.Recipient.Alias)
							}; cx++

							// DiagnosticCode
							if fs.DiagnosticCode == "" {
								// When the value of DiagnosticCode is empty, the Reason should be "undefined"
								if fs.Reason != "undefined" {
									t.Errorf("%s DiagnosticCode is empty but reason is (%s)", ee, fs.Reason)
								}
							}; cx++

							// DiagnosticType
							if fs.DiagnosticType == "" { t.Errorf("%s DiagnosticType is empty", ee) }; cx++
							if strings.Contains(fs.DiagnosticType, " ") {
								t.Errorf("%s DiagnosticType includes space characters (%s)", ee, fs.DiagnosticType)
							}; cx++

							// FeedbackID
							if strings.Contains(fs.FeedbackID, " ") {
								t.Errorf("%s FeedbackID includes space characters (%s)", ee, fs.FeedbackID)
							}; cx++

							// Lhost, Rhost
							if strings.Contains(fs.Lhost, " ") { t.Errorf("%s Lhost includes space characters (%s)", ee, fs.Lhost) }; cx++
							if strings.Contains(fs.Rhost, " ") { t.Errorf("%s Rhost includes space characters (%s)", ee, fs.Rhost) }; cx++

							// ListID
							if strings.Contains(fs.ListID, " ") { t.Errorf("%s ListID includes space characters (%s)", ee, fs.ListID) }; cx++

							// MessageID
							if strings.Contains(fs.MessageID, " ") {
								t.Errorf("%s MessageID includes space characters (%s)", ee, fs.MessageID)
							}; cx++

							// Origin
							if fs.Origin == "" { t.Errorf("%s Origin is empty", ee) }; cx++

							// Command
							if fs.Command != "" && command.Test(fs.Command) == false {
								t.Errorf("%s Command (%s) is an invalid SMTP command", ee, fs.Command)
							}; cx++

							// Subject (not needed to test)
							// Timestamp (not needed to test)
							// TimezoneOffset (not needed to test)

							// Token
							if fs.Token      == "" { t.Errorf("%s Token is empty", ee) }; cx++
							if len(fs.Token) != 40 { t.Errorf("%s Token (%s) is not 40 characaters", ee, fs.Token) }; cx++

							if jx := fs.Dump(); len(jx) == 0 {
								// Dump()
								t.Errorf("%s Dump() returned an empty string", ee)

							} else {
								// Check the string as a JSON
								if strings.Contains(jx, "{") == false {
									t.Errorf("%s Dump() returned invalid JSON string (%s)", ee, jx[:20])
								}
							}; cx++
						} else{
							// THe number of fact is greater than the number of expected values
							t.Errorf("%s missing the expected values", ee); cx++
						}
					}
				}
			}
		})
	}
	t.Logf("The number of tests = %d", cx)
}

