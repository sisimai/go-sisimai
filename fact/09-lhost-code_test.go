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
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rhost"
import "libsisimai.org/sisimai/v5/address"
import "libsisimai.org/sisimai/v5/rfc1123"
import "libsisimai.org/sisimai/v5/rfc1894"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/smtp/command"
import sisimbox "libsisimai.org/sisimai/v5/mail"

type IsExpected struct {
	Label      string // "01" or "1025"
	Index      uint8  // 1,2,3,...
	Status     string // "5.1.1" or empty
	ReplyCode  string // "550" or empty
	Reason     string // "userunknown"
	HardBounce bool   // true or false
	AnotherOne string // "Feedback-Type" or other value
}
var SampleRoot = "set-of-emails"
var PublicDirs = filepath.Join("maildir", "bsd")
var SecretDirs = "private"
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

var TestReturn = map[string]interface{}{"neko-dono": []string{"Michitsuna", "Suzu"}}
var CallbackFn = func(arg *sis.CallbackArg0) (map[string]interface{}, error) { return TestReturn, nil }
var ArgForRise = &sis.DecodingArgs{Delivered: true, Vacation: true, Callback0: CallbackFn}

// EngineTest is called from lhost/*_test.go, rhost/*_test.go, rfc3464/lib_test.go, arf/lib_test.go.
//   Arguments:
//     - t (*testing.T):             Test object
//     - enginename (string)         MTA module name such as "OpenSMTPD"
//     - isexpected ([][]IsExpected) The list of results
//     - publictest (bool)           false if set-of-emails/private
func EngineTest(t *testing.T, enginename string, isexpected [][]IsExpected, publictest bool) {
	cx         := 0
	prefixpath := filepath.Join("..", SampleRoot)
	hostprefix := ""
	remotehost := false
	rhostclass := ""

	if strings.HasPrefix(t.Name(), "TestLhost") == true { hostprefix = "lhost-" }
	if strings.HasPrefix(t.Name(), "TestRhost") == true {
		// TestRhost***, rhost-***-01.eml
		hostprefix = "rhost-"
		remotehost = true
		rhostclass = strings.Replace(t.Name(), "TestRhost", "", 1)
	}

	if publictest == true {
		// Public samples are in set-of-emails/maildir/bsd/lhost-*.eml
		prefixpath += string(os.PathSeparator) + filepath.Join(PublicDirs, hostprefix + strings.ToLower(enginename))

	} else {
		// Private samples are in set-of-emails/private/lhost-* directory
		prefixpath += filepath.Join(SecretDirs, hostprefix, strings.ToLower(enginename))
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
				cx++; if _, nyaan := os.Stat(ef); nyaan != nil {
					t.Errorf("%s cannot read the sample email: %s", ee, nyaan)
				}
			} else {
				// Try to find the email file path in the set-of-emails/private, the file exists as
				// a file name such as private/lhost-opensmtpd/1012-933ce597.eml
				if _, nyaan := os.Stat(prefixpath); nyaan == nil {
					match, nyaan  := filepath.Glob(prefixpath + e[0].Label + "-*.eml")
					cx++; if nyaan      != nil { t.Errorf("%s something wrong: %s", ee, nyaan) }
					cx++; if len(match) == 0   { t.Errorf("%s email not found: %s", ee, nyaan) }
					for _, f := range match { ef = f; break }

				} else {
					// set-of-emails/private directory exists only in the developer's PC
					return
				}
			}

			emailthing, nyaan := sisimbox.Rise(ef)
			cx++; if nyaan != nil {
				// No sample email specified in fact/*-test.go
				t.Errorf("%s failed to load the sample email: %s", ee, nyaan)

			} else {
				// Check each value in sis.Fact{}
				sisi := []sis.Fact{}; for {
					if mesg, nyaan := emailthing.Read(); nyaan != nil {
						// Failed to read the email
						cx++; if errors.Is(nyaan, io.EOF) {
							// sisimai has reached to the end of email/directory
							break

						} else {
							// Something wrong, sisimai failed to read the email as a text
							t.Errorf("%s failed to read the sample email: %s", ee, nyaan)
							continue
						}
					} else {
						// Read and decode each email file as a string
						cx++; if emailthing.Size == 0 { t.Errorf("%s %s is empty", ee, ef); continue }

						moji.ToLF(mesg); facts, nyaan := Rise(mesg, emailthing.Path, ArgForRise)
						cx++; if nyaan != nil && len(nyaan) > 0 { t.Logf("%s %s", ee, nyaan[0].Error()) }
						if facts != nil && len(facts) != 0 { sisi = append(sisi, facts...) }
					}
				}
				cx++; if len(sisi) == 0 { t.Errorf("%s failed to decode any bounce message in %s", ee, ef) }

				cx++; for j, fs := range sisi {
					// Compare each decoded element with each expected value
					cx++; if j < len(e) {
						ev := e[j]
						ee  = fmt.Sprintf("%s[%4s-%02d]", enginename, ev.Label, ev.Index)

						cx++; if fs.DeliveryStatus != ev.Status {
							// DeliveryStatus
							t.Errorf("%s Status is (%s) but (%s)", ee, fs.DeliveryStatus, ev.Status)
						}
						cx++; if strings.Contains(fs.DeliveryStatus, "\n") { t.Errorf("%s DeliveryStatus includes 'LF'", ee) }
						cx++; if strings.Contains(fs.DeliveryStatus, "\r") { t.Errorf("%s DeliveryStatus includes 'CR'", ee) }

						cx++; if fs.ReplyCode != ev.ReplyCode {
							// ReplyCode
							t.Errorf("%s ReplyCode is (%s) but (%s)", ee, fs.ReplyCode, ev.ReplyCode)
						}
						cx++; if strings.Contains(fs.ReplyCode, "\n") { t.Errorf("%s ReplyCode includes 'LF'", ee) }
						cx++; if strings.Contains(fs.ReplyCode, "\r") { t.Errorf("%s ReplyCode includes 'CR'", ee) }

						cx++; if fs.Reason != ev.Reason {
							// Reason
							t.Errorf("%s Reason is (%s) but (%s)", ee, fs.Reason, ev.Reason)
						}
						cx++; if strings.Contains(fs.Reason, "\n") { t.Errorf("%s Reason includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Reason, "\r") { t.Errorf("%s Reason includes 'CR'", ee) }

						cx++; if fs.HardBounce != ev.HardBounce {
							// HardBounce
							t.Errorf("%s HardBounce is (%t) but (%t)", ee, fs.HardBounce, ev.HardBounce)
						}

						cx++; if ev.AnotherOne != "" {
							// AnotherOne
							cx++; if fs.Reason == "feedback" && fs.FeedbackType != ev.AnotherOne {
								// For example "FeedbackType"
								t.Errorf("%s FeedbackType is (%s) but (%s)", ee, fs.FeedbackType, ev.AnotherOne)
							}
						}

						if fs.DecodedBy != enginename {
							// DecodedBy
							altdecoder := enginename
							if len(Alternates[fs.DecodedBy]) > 0 {
								// The MTA module in lhost/ is a removed module
								// https://github.com/sisimai/go-sisimai/issues/7
								for _, as := range Alternates[fs.DecodedBy] {
									if as == enginename { altdecoder = fs.DecodedBy; break }
								}
								if enginename == "AmazonSES" {
									// AmazonSES or RFC3464
									if fs.DecodedBy == "AmazonSES" || fs.DecodedBy == "RFC3464" { altdecoder = fs.DecodedBy }
								}
							}

							cx++; if remotehost == true {
								// Check the rhost, lhost, and destination value
								cv := rhost.Name(&fs)
								if cv == "" { t.Errorf("%s rhost.Name returns (empty) but (%s)", ee, rhostclass) }

							} else {
								// Test for lhost-*.eml
								if fs.DecodedBy != altdecoder {
									t.Errorf("%s DecodedBy is (%s) but (%s)", ee, fs.DecodedBy, altdecoder)
								}
							}
						}
						cx++; if strings.Contains(fs.DecodedBy, "\n") { t.Errorf("%s DecodedBy includes 'LF'", ee) }
						cx++; if strings.Contains(fs.DecodedBy, "\r") { t.Errorf("%s DecodedBy includes 'CR'", ee) }

						/* Other fields except above */
						// Action
						cx++; if fs.Reason != "feedback" && fs.Reason != "vacation" {
							// Action is empty when the bounce mesage is a feedback loop
							cx++; if fs.Action == "" { t.Errorf("%s Action is empty", ee) }
							cx++; if rfc1894.ActionList[fs.Action] == false {
								t.Errorf("%s Action (%s) is an invalid value", ee, fs.Action)
							}
						}
						cx++; if strings.Contains(fs.Action, "\n") { t.Errorf("%s Action includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Action, "\r") { t.Errorf("%s Action includes 'CR'", ee) }

						// Addresser
						cx++; if fs.Addresser.Address == "" { t.Errorf("%s Addresser.Address is empty", ee) }
						cx++; if rfc5322.IsQuotedAddress(fs.Addresser.Address) == false {
							cx++; if fs.Addresser.Alias != "" && strings.Contains(fs.Addresser.Address, "+") == false {
								t.Errorf("%s Addresser.Alias is (%s) not empty", ee, fs.Addresser.Alias)
							}
							cx++; if fs.Addresser.Verp  != "" && strings.Contains(fs.Addresser.Address, "=") == false {
								t.Errorf("%s Addresser.Verp is (%s) not empty", ee, fs.Addresser.Verp)
							}
						}
						cx++; if strings.Contains(fs.Addresser.Address, "\n") { t.Errorf("%s Addresser.Address includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Addresser.Address, "\r") { t.Errorf("%s Addresser.Address includes 'CR'", ee) }

						cx++; if address.IsMailerDaemon(fs.Addresser.Address) == false {
							cx++; if fs.Addresser.User == "" { t.Errorf("%s Addresser.User is empty", ee) }
							cx++; if fs.Addresser.Host == "" { t.Errorf("%s Addresser.Host is empty", ee) }
							cx++; if rfc1123.IsInternetHost(fs.Addresser.Host) == false {
								// Is not a valid internet hostname
								t.Errorf("%s Addresser.Host (%s) is not a valid internet hostname", ee, fs.Addresser.Host)
							}
							cx++; if rfc5322.IsEmailAddress(fs.Addresser.Address) == false {
								// Is not a valid email address
								t.Errorf("%s Addresser.Address (%s) is not a valid email address", ee, fs.Addresser.Address)
							}
						}
						cx++; if fs.Addresser.Host != fs.SenderDomain {
							// SenderDomain
							t.Errorf("%s Addresser.Host is (%s) but (%s)", ee, fs.Addresser.Host, fs.SenderDomain)
						}

						// Alias, Recipient, Destination
						cx++; if fs.Recipient.Address == "" { t.Errorf("%s Recipient.Address is empty", ee) }
						cx++; if fs.Recipient.User    == "" { t.Errorf("%s Recipient.User is empty", ee)    }
						cx++; if fs.Recipient.Host    == "" { t.Errorf("%s Recipient.User is empty", ee)    }
						cx++; if fs.Recipient.Host    != fs.Destination {
							// Destination
							t.Errorf("%s Recipient.Host is (%s) but (%s)", ee, fs.Recipient.Host, fs.Destination)
						}
						cx++; if rfc1123.IsInternetHost(fs.Recipient.Host) == false {
							// Is not a valid internet hostname
							t.Errorf("%s Recipient.Host (%s) is not a valid internet hostname", ee, fs.Recipient.Host)
						}
						cx++; if rfc5322.IsEmailAddress(fs.Recipient.Address) == false {
							// Is not a valid email address
							t.Errorf("%s Recipient.Address (%s) is not a valid email address", ee, fs.Recipient.Address)
						}
						cx++; if fs.Recipient.Verp != "" && rfc5322.IsEmailAddress(fs.Recipient.Verp) == false {
							// Is not a valid email address
							t.Errorf("%s Recipient.Verp (%s) is not a valid email address", ee, fs.Recipient.Verp)
						}
						cx++; if fs.Recipient.Alias != "" && rfc5322.IsEmailAddress(fs.Recipient.Alias) == false {
							// Is not a valid email address
							t.Errorf("%s Recipient.Alias (%s) is not a valid email address", ee, fs.Recipient.Alias)
						}
						cx++; if strings.Contains(fs.Recipient.Address, "\n") { t.Errorf("%s Recipient.Address includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Recipient.Address, "\r") { t.Errorf("%s Recipient.Address includes 'CR'", ee) }

						// DiagnosticCode
						cx++; if strings.Contains(fs.DiagnosticCode, "\n") { t.Errorf("%s DiagnosticCode includes 'LF'", ee) }
						cx++; if strings.Contains(fs.DiagnosticCode, "\r") { t.Errorf("%s DiagnosticCode includes 'CR'", ee) }

						// DiagnosticType
						cx++; if fs.Reason != "feedback" && fs.Reason != "vacation" {
							// DiagnosticType is empty when the bounce mesage is a feedback loop
							cx++; if fs.DiagnosticType == "" { t.Errorf("%s DiagnosticType is empty", ee) }
						}
						cx++; if strings.Contains(fs.DiagnosticType, " ") {
							t.Errorf("%s DiagnosticType includes space characters (%s)", ee, fs.DiagnosticType)
						}
						cx++; if strings.Contains(fs.DiagnosticType, "\n") { t.Errorf("%s DiagnosticType includes 'LF'", ee) }
						cx++; if strings.Contains(fs.DiagnosticType, "\r") { t.Errorf("%s DiagnosticType includes 'CR'", ee) }

						// FeedbackID
						cx++; if strings.Contains(fs.FeedbackID, " ") {
							t.Errorf("%s FeedbackID includes space characters (%s)", ee, fs.FeedbackID)
						}
						cx++; if strings.Contains(fs.FeedbackID, "\n") { t.Errorf("%s FeedbackID includes 'LF'", ee) }
						cx++; if strings.Contains(fs.FeedbackID, "\r") { t.Errorf("%s FeedbackID includes 'CR'", ee) }

						// Lhost, Rhost
						cx++; if strings.Contains(fs.Lhost, " ") { t.Errorf("%s Lhost includes space characters (%s)", ee, fs.Lhost) }
						cx++; if strings.Contains(fs.Rhost, " ") { t.Errorf("%s Rhost includes space characters (%s)", ee, fs.Rhost) }
						cx++; if strings.Contains(fs.Lhost, "\n"){ t.Errorf("%s Lhost includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Lhost, "\r"){ t.Errorf("%s Lhost includes 'CR'", ee) }
						cx++; if strings.Contains(fs.Rhost, "\n"){ t.Errorf("%s Rhost includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Rhost, "\r"){ t.Errorf("%s Rhost includes 'CR'", ee) }

						// ListID
						cx++; if strings.Contains(fs.ListID, " ") { t.Errorf("%s ListID includes space characters (%s)", ee, fs.ListID) }
						cx++; if strings.Contains(fs.ListID, "\n"){ t.Errorf("%s ListID includes 'LF'", ee) }
						cx++; if strings.Contains(fs.ListID, "\r"){ t.Errorf("%s ListID includes 'CR'", ee) }

						// MessageID
						cx++; if strings.Contains(fs.MessageID, " ") {
							t.Errorf("%s MessageID includes space characters (%s)", ee, fs.MessageID)
						}
						cx++; if strings.Contains(fs.MessageID, "\n"){ t.Errorf("%s MessageID includes 'LF'", ee) }
						cx++; if strings.Contains(fs.MessageID, "\r"){ t.Errorf("%s MessageID includes 'CR'", ee) }

						// Origin
						cx++; if fs.Origin == "" { t.Errorf("%s Origin is empty", ee) }
						cx++; if strings.Contains(fs.Origin, "\n"){ t.Errorf("%s Origin includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Origin, "\r"){ t.Errorf("%s Origin includes 'CR'", ee) }

						// Command
						cx++; if fs.Command != "" && command.Test(fs.Command) == false {
							t.Errorf("%s Command (%s) is an invalid SMTP command", ee, fs.Command)
						}
						cx++; if strings.Contains(fs.Command, "\n"){ t.Errorf("%s Command includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Command, "\r"){ t.Errorf("%s Command includes 'CR'", ee) }

						// Subject (not needed to test)
						// Timestamp
						cx++; if fs.Timestamp.Day() == 0 {
							t.Errorf("%s Timestamp (%s) is an invalid string", ee, fs.Timestamp)
						}

						// TimezoneOffset
						cx++; if strings.Contains(fs.TimezoneOffset, " ") {
							t.Errorf("%s TimezoneOffset (%s) includes space characters", ee, fs.TimezoneOffset)
						}
						cx++; if strings.Contains(fs.TimezoneOffset, "\n"){ t.Errorf("%s TimezoneOffset includes 'LF'", ee) }
						cx++; if strings.Contains(fs.TimezoneOffset, "\r"){ t.Errorf("%s TimezoneOffset includes 'CR'", ee) }

						// Token
						cx++; if fs.Token      == "" { t.Errorf("%s Token is empty", ee) }
						cx++; if len(fs.Token) != 40 { t.Errorf("%s Token (%s) is not 40 characaters", ee, fs.Token) }
						cx++; if strings.Contains(fs.Token, "\n"){ t.Errorf("%s Token includes 'LF'", ee) }
						cx++; if strings.Contains(fs.Token, "\r"){ t.Errorf("%s Token includes 'CR'", ee) }

						cx++; if jx, je := fs.Dump(); len(jx) == 0 || je != nil {
							// Dump()
							t.Errorf("%s Dump() returned an empty string or error: %s", ee, je)

						} else {
							// Check the string as a JSON
							cx++; if strings.Contains(jx, "\n"){ t.Errorf("%s JSON includes 'LF'", jx) }
							cx++; if strings.Contains(jx, "\r"){ t.Errorf("%s JSON includes 'CR'", jx) }
							cx++; if strings.Contains(jx, "{") == false {
								t.Errorf("%s Dump() returned invalid JSON string (%s)", ee, jx[:20])
							}
						}
					} else{
						// THe number of fact is greater than the number of expected values
						t.Errorf("%s missing the expected values", ee)
					}
				}
			}
		})
	}
	t.Logf("The number of tests = %d", cx)
}

