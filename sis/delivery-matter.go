// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _ _                      __  __       _   _            
// |  _ \  ___| (_)_   _____ _ __ _   _|  \/  | __ _| |_| |_ ___ _ __ 
// | | | |/ _ \ | \ \ / / _ \ '__| | | | |\/| |/ _` | __| __/ _ \ '__|
// | |_| |  __/ | |\ V /  __/ |  | |_| | |  | | (_| | |_| ||  __/ |   
// |____/ \___|_|_| \_/ \___|_|   \__, |_|  |_|\__,_|\__|\__\___|_|   
//                                |___/                               

package sis
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/rfc1894"
import "libsisimai.org/sisimai/v5/rfc1123"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/smtp/reply"
import "libsisimai.org/sisimai/v5/smtp/status"
import "libsisimai.org/sisimai/v5/smtp/command"

type DeliveryMatter struct {
	Action       string     // The value of Action header
	Agent        string     // MTA name
	Alias        string     // The value of alias entry(RHS)
	Command      string     // SMTP command in the message body
	Date         string     // The value of Last-Attempt-Date header
	Diagnosis    string     // The value of Diagnostic-Code header
	FeedbackType string     // Feedback type
	Lhost        string     // The value of Received-From-MTA header
	Reason       string     // Temporary reason of bounce
	Recipient    string     // The value of Final-Recipient header
	ReplyCode    string     // SMTP Reply Code
	Rhost        string     // The value of Remote-MTA header
	Spec         string     // Protocl specification
	Status       string     // The value of Status header
}

// TailDeliveryMatter returns the last element pointer of DeliveryMatter struct
//   Arguments:
//     - argv0 (*[]DeliveryMatter): The pointer to []DeliveryMatter
//   Returns:
//     - (*DeliveryMatter):         The last element pointer of DeliveryMatter struct
func TailDeliveryMatter(argv0 *[]DeliveryMatter) *DeliveryMatter {
	width := len(*argv0); if width == 0 { return nil }
	return &(*argv0)[width - 1]
}

// NextDeliveryMatter appends a new element and returns the last element pointer
//   Arguments:
//     - argv0 (*[]DeliveryMatter): The pointer to []DeliveryMatter
//   Returns:
//     - (*DeliveryMatter):         The last element pointer of DeliveryMatter struct
func NextDeliveryMatter(argv0 *[]DeliveryMatter) *DeliveryMatter {
	*argv0 = append(*argv0, DeliveryMatter{})
	return &(*argv0)[len(*argv0) - 1]
}

// *DeliveryMatter.Select returns the current value of the sis.DeliveryMatter instance.
//   Arguments:
//     - argv0 (string): Lower-cased member name of sis.DeliveryMatter
//   Returns:
//     - (string):       The value of the member name specified at argv0
func(this *DeliveryMatter) Select(argv0 string) string {
	switch argv0 {
		case "action":       return this.Action
		case "agent":        return this.Agent
		case "alias":        return this.Alias
		case "command":      return this.Command
		case "date":         return this.Date
		case "diagnosis":    return this.Diagnosis
		case "feedbacktype": return this.FeedbackType
		case "lhost":        return this.Lhost
		case "reason":       return this.Reason
		case "recipient":    return this.Recipient
		case "replycode":    return this.ReplyCode
		case "rhost":        return this.Rhost
		case "spec":         return this.Spec
		case "status":       return this.Status
		default:             return ""
	}
}

// *DeliveryMatter.Update set the argument into the member of sis.DeliveryMatter instance.
//   Arguments:
//     - argv0 (string): Lower-cased member name of sis.DeliveryMatter
//     - argv1 (string): New value to be updated
//   Returns:
//     - (bool):         true if it has updated successfully
func(this *DeliveryMatter) Update(argv0 string, argv1 string) bool {
	if argv0 == "" || argv1 == "" { return false }

	//actionlist := []string{"delayed", "delivered", "expanded", "failed", "relayed"}
	feedbacklo := []string{"abuse", "dkim", "fraud", "miscategorized", "not-spam", "opt-out", "virus", "other"}

	switch argv0 {
		default: return false
		case "action":       if rfc1894.ActionList[argv1] { this.Action = argv1 }    // Only valid values are accepted
		case "agent":        this.Agent = argv1                                      // Any value is accepted
		case "alias":        if rfc5322.IsEmailAddress(argv1) { this.Alias = argv1 } // Only valid email addresses are accepted
		case "command":      if command.Test(argv1) { this.Command = argv1 }         // Only valid values are accepted
		case "date":         this.Date = argv1                                       // Any value is accepted
		case "diagnosis":    this.Diagnosis = argv1                                  // Any value is accepted
		case "feedbacktype": if slices.Contains(feedbacklo, argv1) { this.FeedbackType = argv1      } // Only valid values are accepted
		case "lhost":        if rfc1123.IsInternetHost(argv1) { this.Lhost = strings.ToLower(argv1) } // Only valid hostnames are accepted
		case "reason":       this.Reason = strings.ToLower(argv1)
		case "recipient":    if rfc5322.IsEmailAddress(argv1) { this.Recipient = argv1 } // Only valid email addresses are accepted
		case "replycode":    if reply.Test(argv1) { this.ReplyCode = argv1 }             // Only valid SMTP reply codes are accepted
		case "rhost":        if rfc1123.IsInternetHost(argv1) { this.Rhost = strings.ToLower(argv1) } // Only valid hostnames are accepted
		case "spec":         this.Spec = argv1                             // Any value is accepted
		case "status":       if status.Test(argv1) { this.Status = argv1 } // Only valid SMTP status codes are accepted
	}
	return true
}

// *DeliveryMatter.AsRFC1894 returns a lower-cased member name converted from a field name defined in RFC1894.
//   Arguments:
//     - argv1 (string): Field name defined in RFC1894
//   Returns:
//     - (string):       Member name of sis.DeliveryMatter struct
func(this *DeliveryMatter) AsRFC1894(argv1 string) string {
	// Available values are the followings:
	// - "action":             Action    (list)
	// - "arrival-date":       Date      (date)
	// - "diagnostic-code":    Diagnosis (code)
	// - "final-recipient":    Recipient (addr)
	// - "last-attempt-date":  Date      (date)
	// - "original-recipient": Alias     (addr)
	// - "received-from-mta":  Lhost     (host)
	// - "remote-mta":         Rhost     (host)
	// - "reporting-mta":      Lhost     (host)
	// - "status":             Status    (stat)
	// - "x-actual-recipient": Alias     (addr)
	if argv1 == "" { return "" }

	switch argv1 {
		default:                                         return ""
		case "action", "status":                         return argv1
		case "arrival-date", "last-attempt-date":        return "date"
		case "diagnostic-code":                          return "diagnosis"
		case "final-recipient":                          return "recipient"
		case "original-recipient", "x-actual-recipient": return "alias"
		case "reporting-mta":                            return "lhost"
		case "remote-mta", "received-from-mta":          return "rhost"
	}
}

