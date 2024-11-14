// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ___                                   ____  _____ ____  
// | | |__   ___  ___| |_   / / \   _ __ ___   __ _ _______  _ __ / ___|| ____/ ___| 
// | | '_ \ / _ \/ __| __| / / _ \ | '_ ` _ \ / _` |_  / _ \| '_ \\___ \|  _| \___ \ 
// | | | | | (_) \__ \ |_ / / ___ \| | | | | | (_| |/ / (_) | | | |___) | |___ ___) |
// |_|_| |_|\___/|___/\__/_/_/   \_\_| |_| |_|\__,_/___\___/|_| |_|____/|_____|____/ 
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from Amazon SES(Sending): https://aws.amazon.com/ses/
	InquireFor["AmazonSES"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email(JSON)
		// @return   RisingUnderway      RisingUnderway structure
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
			p1 := strings.Index(bf.Body, "\n\n--\n"); if p1 > 0 { emailparts[0] = bf.Body[:p1] }
			p2 := strings.Index(emailparts[0], `"Message"`)

			if p2 > 0 {
				// The JSON included in the email is a format like the following:
				// {
				//  "Type" : "Notification",
				//  "MessageId" : "02f86d9b-eecf-573d-b47d-3d1850750c30",
				//  "TopicArn" : "arn:aws:sns:us-west-2:123456789012:SES-EJ-B",
				//  "Message" : "{\"notificationType\"...
				if strings.Contains(emailparts[0], "!\n ") { emailparts[0] = strings.ReplaceAll(emailparts[0], "!\n ", "") }
				if strings.Contains(emailparts[0], "\\")   { emailparts[0] = strings.ReplaceAll(emailparts[0], "\\",   "") }
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

		dscontents := []sis.DeliveryMatter{{}}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

