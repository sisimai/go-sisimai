// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ___                                   ____  _____ ____  
// | | |__   ___  ___| |_   / / \   _ __ ___   __ _ _______  _ __ / ___|| ____/ ___| 
// | | '_ \ / _ \/ __| __| / / _ \ | '_ ` _ \ / _` |_  / _ \| '_ \\___ \|  _| \___ \ 
// | | | | | (_) \__ \ |_ / / ___ \| | | | | | (_| |/ / (_) | | | |___) | |___ ___) |
// |_|_| |_|\___/|___/\__/_/_/   \_\_| |_| |_|\__,_/___\___/|_| |_|____/|_____|____/ 
import "os"
import "fmt"
import "errors"
import "strings"
import "encoding/json"
import "sisimai/sis"
import "sisimai/rfc1123"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Amazon SES(Sending): https://aws.amazon.com/ses/
	InquireFor["AmazonSES"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email(JSON)
		// @return   RisingUnderway      RisingUnderway structure
		// @see https://docs.aws.amazon.com/ses/latest/dg/notification-contents.html
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := false
		emailparts := []string{bf.Body, ""}
		for {
			// Remote the following string begins with "--"
			// --
			// If you wish to stop receiving notifications from this topic, please click or visit the link below to unsubscribe:
			// https://sns.us-west-2.amazonaws.com/unsubscribe.html?SubscriptionArn=arn:aws:sns:us-west-2:1...
			nt := "notificationType"
			p1 := strings.Index(bf.Body, "\n\n--\n")
			if p1 > 0 { emailparts[0] = bf.Body[:p1] }
			if strings.Contains(emailparts[0], "!\n ") { emailparts[0] = strings.ReplaceAll(emailparts[0], "!\n ", "") }
			p2 := strings.Index(emailparts[0], `"Message"`)

			if p2 > 0 {
				// The JSON included in the email is a format like the following:
				// {
				//  "Type" : "Notification",
				//  "MessageId" : "02f86d9b-eecf-573d-b47d-3d1850750c30",
				//  "TopicArn" : "arn:aws:sns:us-west-2:123456789012:SES-EJ-B",
				//  "Message" : "{\"notificationType\"...
				if strings.Contains(emailparts[0], "\\") { emailparts[0] = strings.ReplaceAll(emailparts[0], "\\",   "") }
				p3 := sisimoji.IndexOnTheWay(emailparts[0], "{",  p2 + 9)
				p4 := sisimoji.IndexOnTheWay(emailparts[0], "\n", p2 + 9)
				emailparts[0] = emailparts[0][p3:p4]
				emailparts[0] = strings.TrimRight(emailparts[0], ",")
				emailparts[0] = strings.TrimRight(emailparts[0], `"`)
			}

			if strings.Contains(emailparts[0], nt)   == false { break }
			if strings.HasPrefix(emailparts[0], "{") == false { break }
			if strings.HasSuffix(emailparts[0], "}") == false { break }
			proceedsto = true; break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		// https://docs.aws.amazon.com/en_us/ses/latest/DeveloperGuide/notification-contents.html
		type eachHeader struct {
			Name  string // "MIME-Version"
			Value string // "1.0"
		}
		type commonHead struct {
			From      []string        // ["Sender Name <sender@example.com>"]
			To        []string        // ["Recipient Name <recipient@example.com>"]
			Date      string          // "Mon, 08 Oct 2018 14:05:45 +0000"
			MessageID string          // "custom-message-ID"
			Subject   string          // "Message sent using Amazon SES"
		}
		type mailObject struct {
			Timestamp        string   // "2018-10-08T14:05:45 +0000"
			MessageID        string   // "000001378603177f-7a5433e7-8edb-42ae-af10-f0181f34d6ee-000000"
			Source           string   // "sender@example.com"
			SourceARN        string   // "arn:aws:ses:us-east-1:888888888888:identity/example.com"
			SourceIP         string   // "127.0.3.0"
			SendingAccountID string   // "123456789012"
			CallerIdentity   string   // ?
			Destination      []string // ["recipient@example.com"]
			HeadersTruncated bool
			Headers          []eachHeader
			CommonHeaders    commonHead
		}
		RFC822Head := func(it *mailObject) string {
			// RFC822Head() returns the original email message (headers only)
			allheaders := ""
			for _, e := range it.Headers { allheaders += fmt.Sprintf("%s: %s\n", e.Name, e.Value) }
			if it.CommonHeaders.Date    != "" { allheaders += fmt.Sprintf("Date: %s\n", it.CommonHeaders.Date)       }
			if it.CommonHeaders.Subject != "" { allheaders += fmt.Sprintf("Subject: %s\n", it.CommonHeaders.Subject) }
			return allheaders
		}

		//-----------------------------------------------------------------------------------------
		// "notificationType": "Bounce"
		// https://docs.aws.amazon.com/ses/latest/dg/notification-contents.html#bounce-object
		//
		// Bounce types
		//   The bounce object contains a bounce type of Undetermined, Permanent, or Transient. The
		//   Permanent and Transient bounce types can also contain one of several bounce subtypes.
		//
		//   When you receive a bounce notification with a bounce type of Transient, you might be
		//   able to send email to that recipient in the future if the issue that caused the message
		//   to bounce is resolved.
		//
		//   When you receive a bounce notification with a bounce type of Permanent, it's unlikely
		//   that you'll be able to send email to that recipient in the future. For this reason, you
		//   should immediately remove the recipient whose address produced the bounce from your
		//   mailing lists.
		//
		// "bounceType"/"bounceSubType" "Desription"
		// Undetermined/Undetermined -- The bounce message didn't contain enough information for
		//                              Amazon SES to determine the reason for the bounce.
		//
		// Permanent/General ---------- When you receive this type of bounce notification, you should
		//                              immediately remove the recipient's email address from your
		//                              mailing list.
		// Permanent/NoEmail ---------- It was not possible to retrieve the recipient email address
		//                              from the bounce message.
		// Permanent/Suppressed ------- The recipient's email address is on the Amazon SES suppression
		//                              list because it has a recent history of producing hard bounces.
		// Permanent/OnAccountSuppressionList
		//                              Amazon SES has suppressed sending to this address because it
		//                              is on the account-level suppression list.
		//
		// Transient/General ---------- You might be able to send a message to the same recipient
		//                              in the future if the issue that caused the message to bounce
		//                              is resolved.
		// Transient/MailboxFull ------ the recipient's inbox was full.
		// Transient/MessageTooLarge -- message you sent was too large
		// Transient/ContentRejected -- message you sent contains content that the provider doesn't allow
		// Transient/AttachmentRejected the message contained an unacceptable attachment
		/*
		reasonmaps := map[string]string {
			"Supressed":                "undefined", // "suppressed" will be assigned (new reason name)
			"OnAccountSuppressionList": "undefined", // "suppressed" will be assigned (new reason name)
			"General":                  "onhold",
			"MailboxFull":              "mailboxfull",
			"MessageTooLarge":          "mesgtoobig",
			"ContentRejected":          "contenterror",
			"AttachmentRejected":       "securityerror",
		}
		*/
		type failedRCPT struct {
			EmailAddress   string     // "bounce@simulator.amazonses.com",
			DiagnosticCode string     // "smtp; 550 5.1.1 user unknown"
			Action         string     // "failed"
			Status         string     // "5.1.1"
		}
		type bounceBack struct {
			BounceType        string  // "Undetermined", "Permanent", "Transient"
			BounceSubType     string  // "General", "Suppressed", "MailboxFull", and so on
			BouncedRecipients []failedRCPT
			Timestamp         string  // "2016-10-21T06:58:02.245Z"
			FeedbackID        string  // "01010157e6083d17-38cf01f3-852d-4401-8e8a-84e67a3e51d8-000000"
			RemoteMTAIP       string  // "127.0.2.0"
			ReportingMTA      string  // "dsn; a27-33.smtp-out.us-west-2.amazonses.com"
		}
		type ReturnedTo struct {
			NotificationType string
			Mail   mailObject
			Bounce bounceBack // This field is present only if the notificationType is Bounce
		}

		//-----------------------------------------------------------------------------------------
		// "notificationType": "Complaint"
		// https://docs.aws.amazon.com/ses/latest/dg/notification-contents.html#complaint-object
		//
		// Complaint types
		//   You may see the following complaint types in the complaintFeedbackType field as assigned
		//   by the reporting ISP, according to the Internet Assigned Numbers Authority website:
		//
		// - abuse:        Indicates unsolicited email or some other kind of email abuse.
		// - auth-failure: Email authentication failure report.
		// - fraud:        Indicates some kind of fraud or phishing activity.
		// - not-spam:     Indicates that the entity providing the report does not consider the message
		//                 to be spam. This may be used to correct a message that was incorrectly tagged
		//                 or categorized as spam.
		// - other:        Indicates any other feedback that does not fit into other registered types.
		// - virus:        Reports that a virus is found in the originating message.
		type complainBy struct { EmailAddress string } // [{"emailAddress": "complaint@simulator.amazonses.com"}]
		type complaints struct {
			ComplainedRecipients []complainBy
			Timestamp             string // "2016-10-21T06:58:02.245Z"
			FeedbackID            string // "01010157e6083d17-38cf01f3-852d-4401-8e8a-84e67a3e51d8-000000"
			ComplaintSubType      string // The value of the complaintSubType field can either be null or OnAccountSuppressionList
			UserAgent             string // "Amazon SES Mailbox Simulator",
			ComplaintFeedbackType string // "abuse"
			ArrivalDate           string // The value of the Arrival-Date or Received-Date field
		}
		type Complained struct {
			NotificationType string
			Mail      mailObject
			Complaint complaints // This field is present only if the notificationType is Complaint 
		}

		//-----------------------------------------------------------------------------------------
		// "notificationType": "Delivery"
		// https://docs.aws.amazon.com/ses/latest/dg/notification-contents.html#delivery-object
		type sentstatus struct {
			Timestamp             string   // "2016-10-21T06:58:02.245Z"
			ProcessingTimeMillis  int      // 5753
			Recipients            []string // ["complaint@simulator.amazonses.com"]
			SMTPResponse          string   // "250 2.6.0 Message received"
			RemoteMTAIP           string   // "127.0.2.0"
			ReportingMTA          string   // "dsn; a27-33.smtp-out.us-west-2.amazonses.com"
		}
		type Deliveries struct {
			NotificationType string
			Mail      mailObject
			Delivery  sentstatus // This field is present only if the notificationType is Complaint 
		}

		//-----------------------------------------------------------------------------------------
		type NotifiedTo struct {
			returnedto *ReturnedTo // "notificationType":"Bounce"
			complained *Complained // "notificationType":"Complaint"
			deliveries *Deliveries // "notificationType":"Delivery"
		}

		//-----------------------------------------------------------------------------------------
		var whatnotify string      // The first character of "notificationType": "B", "C" or "D"
		var notifiedto NotifiedTo  // This instance have 3 types: ReturnedTo, Deliveries, Complained
		var mailinside *mailObject // The pointer to mailObject struct
		var jsonerrors error  = errors.New("Invalid JSON format")
		var jsonstring []byte = []byte(emailparts[0])

		for json.Valid(jsonstring) == true {
			// The JSON string should contain one of the followings in "notificationType" field
			// - "Bounce"
			// - "Complaint"
			// - "Delivery"
			if strings.Contains(emailparts[0], `"notificationType":"Bounce"`) {
				// {"notificationType":"Bounce","bounce":{"bounceType":"Permanent",...
				var p ReturnedTo
				jsonerrors = json.Unmarshal(jsonstring, &p); if jsonerrors != nil { break }
				notifiedto.returnedto = &p
				mailinside = &p.Mail
				whatnotify = "B"

			} else if strings.Contains(emailparts[0], `"notificationType":"Complaint"`) {
				// {"notificationType":"Complaint","complaint":{"complainedRecipients":[{"e...
				var p Complained
				jsonerrors = json.Unmarshal(jsonstring, &p); if jsonerrors != nil { break }
				notifiedto.complained = &p
				mailinside = &p.Mail
				whatnotify = "C"

			} else if strings.Contains(emailparts[0], `"notificationType":"Delivery"`) {
				// {"notificationType":"Delivery","mail":{"timestamp":...
				var p Deliveries
				jsonerrors = json.Unmarshal(jsonstring, &p); if jsonerrors != nil { break }
				notifiedto.deliveries = &p
				mailinside = &p.Mail
				whatnotify = "D"

			} else {
				// There is no "notificationType" field or unknown type of "notificationType" field
				// in the JSON string in the message body
				jsonerrors = errors.New("There is no notificationType field or unknown type of notificationType field")
			}
			break
		}
		if whatnotify == "" {
			// Failed to loadl/decode JSON
			fmt.Fprintf(os.Stderr, " ***warning: %s\n", jsonerrors)
			return sis.RisingUnderway{}
		}

		dscontents := []sis.DeliveryMatter{{}}
//		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		if whatnotify == "B" {
			// "notificationType":"Bounce"
			o := &notifiedto.returnedto.Bounce
			for _, e := range (*o).BouncedRecipients {
				// {"emailAddress":"neko@example.jp", "action":"failed", "status":"5.1.1", "diagnosticCode": "..."}
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e.EmailAddress)
				v.Diagnosis = sisimoji.Sweep(e.DiagnosticCode)
				v.Command   = command.Find(v.Diagnosis)
				v.Action    = e.Action
				v.Status    = status.Find(e.Status, "")
				v.ReplyCode = reply.Find(v.Diagnosis, v.Status)
				v.Date      = (*o).Timestamp
				v.Lhost     = rfc1123.Find((*o).ReportingMTA)
			}
		} else if whatnotify == "C" {
			// "notificationType":"Complaint"
			o := &notifiedto.complained.Complaint
			for _, e := range (*o).ComplainedRecipients {
				// {"emailAddress":"neko@example.jp"}
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient    = sisiaddr.S3S4(e.EmailAddress)
				v.Reason       = "feedback"
				v.FeedbackType = (*o).ComplaintFeedbackType
				v.Date         = (*o).Timestamp
				v.Diagnosis    = fmt.Sprintf(`{"feedbackid":"%s", "useragent":"%s"}`, (*o).FeedbackID, (*o).UserAgent)
			}
		} else if whatnotify == "D" {
			// "notificationType":"Delivery"
			o := &notifiedto.deliveries.Delivery
			for _, e := range (*o).Recipients {
				// {"emailAddress":"neko@example.jp"}
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e)
				v.Reason    = "delivered"
				v.Date      = (*o).Timestamp
				v.Lhost     = (*o).ReportingMTA
				v.Diagnosis = (*o).SMTPResponse
				v.Status    = status.Find(v.Diagnosis, "")
				v.ReplyCode = reply.Find(v.Diagnosis, v.Status)
			}
		}
		emailparts[1] = RFC822Head(mailinside)
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

