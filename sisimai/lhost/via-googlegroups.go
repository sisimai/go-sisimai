// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                   _       ____                           
// | | |__   ___  ___| |_   / / ___| ___   ___   __ _| | ___ / ___|_ __ ___  _   _ _ __  ___ 
// | | '_ \ / _ \/ __| __| / / |  _ / _ \ / _ \ / _` | |/ _ \ |  _| '__/ _ \| | | | '_ \/ __|
// | | | | | (_) \__ \ |_ / /| |_| | (_) | (_) | (_| | |  __/ |_| | | | (_) | |_| | |_) \__ \
// |_|_| |_|\___/|___/\__/_/  \____|\___/ \___/ \__, |_|\___|\____|_|  \___/ \__,_| .__/|___/
//                                              |___/                             |_|        
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Google Groups: https://groups.google.com
	InquireFor["GoogleGroups"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true                        { return sis.RisingUnderway{} }
		if strings.Contains(bf.Payload, "Google Groups") == false { return sis.RisingUnderway{} }
		if len(bf.Headers["x-failed-recipients"])        == 0     { return sis.RisingUnderway{} }
		if len(bf.Headers["x-google-smtp-source"])       == 0     { return sis.RisingUnderway{} }

		// X-Google-Smtp-Source: APXvYqx67WVONuSclAC3HckRuO768rET6VCNXk6xYv7cW5I1l9kkn35pT4zE29miuroXfMsHzqeVDrOoIjb8hdt7tjtNL2XAomNl7FA=
		// From: Mail Delivery Subsystem <mailer-daemon@googlemail.com>
		// Subject: Delivery Status Notification (Failure)
		// X-Failed-Recipients: libsisimai@googlegroups.com
		if strings.Contains(bf.Headers["from"][0], "<mailer-daemon@googlemail.com>")  == false { return sis.RisingUnderway{} }
		if strings.Contains(bf.Headers["subject"][0], "Delivery Status Notification") == false { return sis.RisingUnderway{} }

		// Hello kijitora@libsisimai.org,
		//
		// We're writing to let you know that the group you tried to contact (group-name)
		// may not exist, or you may not have permission to post messages to the group.
		// A few more details on why you weren't able to post:
		//
		//  * You might have spelled or formatted the group name incorrectly.
		//  * The owner of the group may have removed this group.
		//  * You may need to join the group before receiving permission to post.
		//  * This group may not be open to posting.
		//
		// If you have questions related to this or any other Google Group,
		// visit the Help Center at https://groups.google.com/support/.
		//
		// Thanks,
		//
		// Google Groups
		boundaries := []string{"----- Original message -----", "Content-Type: message/rfc822"}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		entiremesg := strings.SplitN(emailparts[0], "\n\n", 5); entiremesg[len(entiremesg) - 1] = ""
		issuedcode := strings.ReplaceAll(strings.Join(entiremesg, " "), "\n", " ")
		receivedby := []string{""}; if len(bf.Headers["received"]) > 0 { receivedby = bf.Headers["received"] }
		recordwide := [3]string{
			rfc5322.Received(receivedby[0])[1], // rhost
			"onhold",                           // reason
			sisimoji.Sweep(issuedcode),         // diagnosis
		}

		for {
			// * You might have spelled or formatted the group name incorrectly.
			// * The owner of the group may have removed this group.
			// * You may need to join the group before receiving permission to post.
			// * This group may not be open to posting.
			if strings.Count(emailparts[0], "\n *") == 4 { recordwide[1] = "rejected"; break }
			if strings.Count(emailparts[0], "\n*")  == 4 { recordwide[1] = "rejected"; break }
			break
		}

		for _, e := range strings.Split(bf.Headers["x-failed-recipients"][0], ",") {
			// X-Failed-Recipients: neko@example.jp, cat@example.org, ...
			if sisiaddr.IsEmailAddress(e) == false { continue }
			if len(v.Recipient) > 0 {
				// There are multiple recipient addresses in the message body.
				dscontents = append(dscontents, sis.DeliveryMatter{})
				v = &(dscontents[len(dscontents) - 1])
			}
			v.Recipient = sisiaddr.S3S4(e)
			recipients += 1
			v.Rhost     = recordwide[0]
			v.Reason    = recordwide[1]
			v.Diagnosis = recordwide[2]
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

