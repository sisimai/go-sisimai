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

// TailDeliveryMatter returns the last element pointer of DeliveryMatter struct.
//   Arguments:
//     - de ([]DeliveryMatter): []DeliveryMatter
//   Returns:
//     - (*DeliveryMatter): The last element pointer of DeliveryMatter struct.
func TailDeliveryMatter(de []DeliveryMatter) *DeliveryMatter {
	width := len(de); if width == 0 { return nil }
	return &(de)[width - 1]
}

// NextDeliveryMatter appends a new element and returns the last element pointer.
//   Arguments:
//     - de (*[]DeliveryMatter): The pointer to []DeliveryMatter.
//   Returns:
//     - (*DeliveryMatter): The last element pointer of DeliveryMatter struct.
func NextDeliveryMatter(de *[]DeliveryMatter) *DeliveryMatter {
	*de = append(*de, DeliveryMatter{})
	return &(*de)[len(*de) - 1]
}

// *DeliveryMatter.Select returns the current value of the sis.DeliveryMatter instance.
//   Arguments:
//     - field (string): Lower-cased member name of sis.DeliveryMatter.
//   Returns:
//     - (string): The value of the member name specified at "field".
func(de *DeliveryMatter) Select(field string) string {
	switch field {
		case "action":       return de.Action
		case "agent":        return de.Agent
		case "alias":        return de.Alias
		case "command":      return de.Command
		case "date":         return de.Date
		case "diagnosis":    return de.Diagnosis
		case "feedbacktype": return de.FeedbackType
		case "lhost":        return de.Lhost
		case "reason":       return de.Reason
		case "recipient":    return de.Recipient
		case "replycode":    return de.ReplyCode
		case "rhost":        return de.Rhost
		case "spec":         return de.Spec
		case "status":       return de.Status
		default:             return ""
	}
}

// *DeliveryMatter.Update set the argument into the member of sis.DeliveryMatter instance.
//   Arguments:
//     - field (string): Lower-cased member name of sis.DeliveryMatter.
//     - value (string): New value to be updated.
//   Returns:
//     - (bool): true if it has updated successfully.
func(de *DeliveryMatter) Update(field string, value string) bool {
	if field == "" || value == "" { return false }

	//actionlist := []string{"delayed", "delivered", "expanded", "failed", "relayed"}
	feedbacklo := []string{"abuse", "dkim", "fraud", "miscategorized", "not-spam", "opt-out", "virus", "other"}

	switch field {
		default: return false
		case "action":       if rfc1894.ActionList[value] { de.Action = value }    // Only valid values are accepted
		case "agent":        de.Agent = value                                      // Any value is accepted
		case "alias":        if rfc5322.IsEmailAddress(value) { de.Alias = value } // Only valid email addresses are accepted
		case "command":      if command.Test(value) { de.Command = value }         // Only valid values are accepted
		case "date":         de.Date = value                                       // Any value is accepted
		case "diagnosis":    de.Diagnosis = value                                  // Any value is accepted
		case "feedbacktype": if slices.Contains(feedbacklo, value) { de.FeedbackType = value      } // Only valid values are accepted
		case "lhost":        if rfc1123.IsInternetHost(value) { de.Lhost = strings.ToLower(value) } // Only valid hostnames are accepted
		case "reason":       de.Reason = strings.ToLower(value)
		case "recipient":    if rfc5322.IsEmailAddress(value) { de.Recipient = value } // Only valid email addresses are accepted
		case "replycode":    if reply.Test(value) { de.ReplyCode = value }             // Only valid SMTP reply codes are accepted
		case "rhost":        if rfc1123.IsInternetHost(value) { de.Rhost = strings.ToLower(value) } // Only valid hostnames are accepted
		case "spec":         de.Spec = value                             // Any value is accepted
		case "status":       if status.Test(value) { de.Status = value } // Only valid SMTP status codes are accepted
	}
	return true
}

// *DeliveryMatter.AsRFC1894 returns a lower-cased member name converted from a field name defined in RFC1894.
//   Arguments:
//     - field (string): Field name defined in RFC1894.
//   Returns:
//     - (string): Member name of sis.DeliveryMatter struct.
func(de *DeliveryMatter) AsRFC1894(field string) string {
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
	if field == "" { return "" }

	switch field {
		default:                                         return ""
		case "action", "status":                         return field
		case "arrival-date", "last-attempt-date":        return "date"
		case "diagnostic-code":                          return "diagnosis"
		case "final-recipient":                          return "recipient"
		case "original-recipient", "x-actual-recipient": return "alias"
		case "reporting-mta":                            return "lhost"
		case "remote-mta", "received-from-mta":          return "rhost"
	}
}

