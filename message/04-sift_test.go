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
import "os"
import "net/mail"
import "libsisimai.org/sisimai/sis"

func TestSift(t *testing.T) {
	fn := "sisimai/message.sift"
	fs := "sis.BeforeFact"
	ae := "../set-of-emails/maildir/bsd/lhost-postfix-55.eml"
	cx := 0
	bf := new(sis.BeforeFact)
	bx, _ := os.ReadFile(ae); cx++; if len(bx) == 0 {
		t.Fatalf("os.ReadFile(%s) returns an empty string", ae)
	}
	et := string(bx); em, _ := mail.ReadMessage(strings.NewReader(et))
	eb, _ := io.ReadAll(em.Body); cx++; if len(eb) == 0 {
		t.Fatalf("io.ReadAll(%s) returns an empty string", ae)
	}

	bf.Sender  = "MAILER-DAEMON Fri Feb  2 18:30:22 2018"
	bf.Headers = makemap(&em.Header, false)
	bf.Payload = string(eb)

	cx++; if len(bf.Headers) == 0             { t.Errorf("makemap() returns empty headers") }
	cx++; if cv := sift(bf, nil); cv == false { t.Errorf("%s(bf) returns false", fn) }
	cx++; if len(bf.Digest)  == 0             { t.Errorf("bf.Digest is empty") }

	dx := bf.Digest[0]
	cx++; if dx.Action       != "failed"      { t.Errorf("%s.Digest.Action is not `failed`: %s", fs, dx.Action) }
	cx++; if dx.Agent        != "Postfix"     { t.Errorf("%s.Digest.Agent is not `Postfix`: %s", fs, dx.Agent) }
	cx++; if dx.Alias        == ""            { t.Errorf("%s.Digest.Alias is empty", fs) }
	cx++; if dx.Command      != "RCPT"        { t.Errorf("%s.Digest.Command is not `RCPT`: %s", fs, dx.Command) }
	cx++; if dx.FeedbackType != ""            { t.Errorf("%s.Digest.FeedbackType is not empty: %s", fs, dx.FeedbackType) }
	cx++; if dx.Date         == ""            { t.Errorf("%s.Digest.Date is empty", fs) }
	cx++; if dx.Diagnosis    == ""            { t.Errorf("%s.Digest.Diagnosis is empty", fs) }
	cx++; if dx.Lhost        == ""            { t.Errorf("%s.Digest.Lhost is empty", fs) }
	cx++; if dx.Reason       != ""            { t.Errorf("%s.Digest.Reason is not empty: %s", fs, dx.Reason) }
	cx++; if dx.Recipient    == ""            { t.Errorf("%s.Digest.Recipient is empty", fs) }
	cx++; if dx.ReplyCode    != "552"         { t.Errorf("%s.Digest.ReplyCode is not `552`: %s", fs, dx.ReplyCode) }
	cx++; if dx.Rhost        == ""            { t.Errorf("%s.Digest.Rhost is empty", fs) }
	cx++; if dx.Spec         != "SMTP"        { t.Errorf("%s.Digest.Spec is not `SMTP`: %s", fs, dx.Spec) }
	cx++; if dx.Status       != "5.0.0"       { t.Errorf("%s.Digest.Status is not `5.0.0`: %s", fs, dx.Status) }

	t.Logf("The number of tests = %d", cx)
}

