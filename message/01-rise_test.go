// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _____         _      __                                        
// |_   _|__  ___| |_   / / __ ___   ___  ___ ___  __ _  __ _  ___ 
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
//   | |  __/\__ \ |_ / /| | | | | |  __/\__ \__ \ (_| | (_| |  __/
//   |_|\___||___/\__/_/ |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                                      |___/      
import "testing"
import "strings"
import "io"
import "errors"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc791"
import "libsisimai.org/sisimai/rfc1123"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import "libsisimai.org/sisimai/smtp/command"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimbox "libsisimai.org/sisimai/mail"

func TestRise(t *testing.T) {
	fn := "sisimai/message.Rise"
	fs := "BeforeFact"
	cx := 0
	en := 0
	ae := "../../set-of-emails/mailbox/mbox-0"
	c1 := func(arg *sis.CallbackArgs) (map[string]interface{}, error) {
		data := make(map[string]interface{})
		head := []string{"X-Mailer", "Return-Path"}
		for _, e := range head {
			if arg.Payload != nil && len(*arg.Payload) > 0 {
				cb := *arg.Payload
				ce := "\n" + e + ":"
				if j := strings.Index(cb, ce); j > 0 {
					right := cb[j + len(ce):]
					if u := strings.Index(right, "\n"); u > 0 {
						cv := strings.TrimSpace(right[:u])
						data[strings.ToLower(e)] = cv
					}
				}
			}
		}
		data["nekochan"] = []string{"michitsuna", "suzu"}
		return data, nil
	}

	eo, _ := sisimbox.Rise(ae)
	cx++; if eo.Size == 0 { t.Errorf("sisimai/mail.Rise(%s) returns empty string", ae) }

	for {
		if ef, ee := eo.Read(); ef != nil || ee == nil {
			en += 1
			cv := Rise(ef, c1)
			cx++; if len(cv.Errors) > 0 { t.Errorf("%s() returns error: %v", fn, cv.Errors) }
			cx++; if cv.Void() == true  { t.Errorf("%s.Void() returns true", fs) }

			cx++; if  cv.Sender == "" { t.Errorf("%s.Sender is empty", fs)  }
			cx++; for cv.Sender != "" {
				cx++; if strings.Contains(strings.ToUpper(cv.Sender), "MAILER-DAEMON") == false &&
						 strings.Contains(cv.Sender, "@")             == false {
							 t.Errorf("%s.Sender include neither mailer-daemon nor @: %s", fs, cv.Sender)
				}
				break
			}

			cx++; if cv.Payload      == "" { t.Errorf("%s.Payload is empty", fs) }
			cx++; if len(cv.Headers) == 0  { t.Errorf("%s.Headers is empty", fs) }
			cx++; if len(cv.RFC822)  == 0  { t.Errorf("%s.RFC822 is empty", fs) }
			cx++; if len(cv.Digest)  == 0  { t.Errorf("%s.Digest is empty", fs) }

			for _, e := range []string{"to", "from", "subject", "date"} {
				cx++; if len(cv.Headers[e]) == 0  { t.Errorf("%s.Headers[%s] have no element", fs, e); continue }
				cx++; if cv.Headers[e][0]   == "" { t.Errorf("%s.Headers[%s][0] is empty", fs, e) }
			}
			for _, e := range []string{"to", "from"} {
				cx++; if len(cv.RFC822[e])  == 0  { t.Errorf("%s.RFC822[%s] have no element", fs, e); continue }
				cx++; if cv.RFC822[e][0]    == "" { t.Errorf("%s.RFC822[%s][0] is empty", fs, e) }
			}

			for _, e := range cv.Digest {
				cx++; for e.Action != "" {
					cx++; if strings.HasPrefix(e.Action, " ") { t.Errorf("%s.Digest.Action starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Action, " ") { t.Errorf("%s.Digest.Action ends with ` `", fs)   }
					cx++; if e.Action != "failed"             { t.Errorf("%s.Digest.Action is %s", fs, e.Action) }
					break
				}
				cx++; if  e.Agent == "" { t.Errorf("%s.Digest.Agent is empty", fs) }
				cx++; for e.Alias != "" {
					cx++; if e.Alias == e.Recipient { t.Errorf("%s.Digest.Alias is %s", fs, e.Alias) }
					cx++; if sisiaddr.IsEmailAddress(e.Alias) == false && strings.HasPrefix(e.Alias, "|/") == false {
						t.Errorf("%s.Alias is an invalid value: %s", fs, e.Alias)
					}
					break
				}

				cx++; for e.Command != "" {
					cx++; if strings.HasPrefix(e.Command, " ") { t.Errorf("%s.Digest.Command starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Command, " ") { t.Errorf("%s.Digest.Command ends with ` `", fs)   }
					cx++; if command.Test(e.Command) == false  {
						t.Errorf("%s.Digest.Command is an invalid value: %s", fs, e.Command)
					}
					break
				}

				cx++; for e.Date != ""   {
					cx++; if strings.HasPrefix(e.Date, " ") { t.Errorf("%s.Digest.Date starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Date, " ") { t.Errorf("%s.Digest.Date ends with ` `", fs)   }
					cx++; if rfc5322.Date(e.Date) == ""     { t.Errorf("%s.Digest.Date is invalid format: %s", fs, e.Date) }
					break
				}
				cx++; if e.Diagnosis    == "" { t.Errorf("%s.Digest.Diagnosis is empty", fs) }
				cx++; if e.FeedbackType != "" { t.Errorf("%s.Digest.FeedbackType is %s", fs, e.FeedbackType) }
				cx++; for e.Lhost       != "" {
					cx++; if strings.HasPrefix(e.Lhost, " ") { t.Errorf("%s.Digest.Lhost starts with ` `", fs)  }
					cx++; if strings.HasSuffix(e.Lhost, " ") { t.Errorf("%s.Digest.Lhost ends with ` `", fs)    }
					cx++; if rfc1123.IsInternetHost(e.Lhost) == false && rfc791.IsIPv4Address(e.Lhost) == false {
						t.Errorf("%s.Digest.Lhost is an invalid internet host: %s", fs, e.Lhost)
					}
					break
				}
				cx++; for e.Reason != "" {
					cx++; if strings.HasPrefix(e.Reason, " ") { t.Errorf("%s.Digest.Reason starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Reason, " ") { t.Errorf("%s.Digest.Reason ends with ` `", fs)   }
					cx++; if e.Reason != "suspend" { t.Errorf("%s.Digest.Reason is %s", fs, e.Reason) }
					break
				}
				cx++; if  e.Recipient == "" { t.Errorf("%s.Recipient is empty", fs) }
				cx++; for e.Recipient != "" {
					cx++; if e.Recipient == e.Alias { t.Errorf("%s.Recipient is the same value with Alias", fs) }
					cx++; if sisiaddr.IsEmailAddress(e.Recipient) == false {
						t.Errorf("%s.Recipient is an invalid email address: %s", fs, e.Recipient)
					}
					break
				}
				cx++; for e.ReplyCode != "" {
					cx++; if strings.HasPrefix(e.ReplyCode, " ") { t.Errorf("%s.Digest.ReplyCode starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.ReplyCode, " ") { t.Errorf("%s.Digest.ReplyCode ends with ` `", fs)   }
					cx++; if reply.Test(e.ReplyCode) == false {
						t.Errorf("%s.ReplyCode is an invalid SMTP reply code: %s", fs, e.ReplyCode)
					}
					break
				}
				cx++; for e.Rhost != "" {
					cx++; if strings.HasPrefix(e.Rhost, " ") { t.Errorf("%s.Digest.Rhost starts with ` `", fs)  }
					cx++; if strings.HasSuffix(e.Rhost, " ") { t.Errorf("%s.Digest.Rhost ends with ` `", fs)    }
					cx++; if rfc1123.IsInternetHost(e.Rhost) == false && rfc791.IsIPv4Address(e.Rhost) == false {
						t.Errorf("%s.Digest.Rhost is an invalid internet host: %s", fs, e.Rhost)
					}
					break
				}
				cx++; for e.Spec  != "" {
					cx++; if strings.HasPrefix(e.Spec, " ") { t.Errorf("%s.Digest.Spec starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Spec, " ") { t.Errorf("%s.Digest.Spec ends with ` `", fs)   }
					cx++; if strings.EqualFold(e.Spec, "SMTP") == false && strings.EqualFold(e.Spec, "X-UNIX") == false {
						t.Errorf("%s.Spec is an invalid value: %s", fs, e.Spec)
					}
					break
				}
				cx++; if e.Status != "" {
					cx++; if strings.HasPrefix(e.Status, " ") { t.Errorf("%s.Digest.Status starts with ` `", fs) }
					cx++; if strings.HasSuffix(e.Status, " ") { t.Errorf("%s.Digest.Status ends with ` `", fs)   }
					cx++; if status.Test(e.Status) == false {
						t.Errorf("%s.Status is an invalid SMTP status code: %s", fs, e.Status)
					}
					break
				}
			}

			if re, as := cv.Catch.(map[string]interface{}); as {
				if ca, ok := re["nekochan"].([]string); ok {
					cx++; if ca[0] != "michitsuna" { t.Errorf("%s.Catch[nekochan][0] is not michitsuna", fs) }
					cx++; if ca[1] != "suzu"       { t.Errorf("%s.Catch[nekochan][1] is not suzu", fs) }
				}
				if xe, ok := re["x-mailer"].(string);    ok {
					cx++; if xe == "" { t.Errorf("%s.Catch[x-mailer] is empty", fs) }
				}
				if xe, ok := re["return-path"].(string); ok {
					cx++; if xe == "" { t.Errorf("%s.Catch[return-path] is empty", fs) }
				}
			}
		} else {
			if errors.Is(ee, io.EOF) { break }
			continue
		}
	}
	if en != 37 { t.Errorf("%s() returns %d emails", fn, en) }

	et := ""
	ev := Rise(&et, nil)
	cx++; if ev.Void()  == false { t.Errorf("%s.Void() returns false", fs) }
	cx++; if ev.Sender  != ""    { t.Errorf("%s.Sender is not empty: %s", fs, ev.Sender)   }
	cx++; if ev.Payload != ""    { t.Errorf("%s.Payload is not empty: %s", fs, ev.Payload) }
	cx++; if len(ev.Errors) > 0  { t.Errorf("%s.Errors is not empty: %v", fs, ev.Errors)   }

	t.Logf("The number of tests = %d", cx)
}

