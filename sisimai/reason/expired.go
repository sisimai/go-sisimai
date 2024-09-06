// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____            _              _ 
// | ____|_  ___ __ (_)_ __ ___  __| |
// |  _| \ \/ / '_ \| | '__/ _ \/ _` |
// | |___ >  <| |_) | | | |  __/ (_| |
// |_____/_/\_\ .__/|_|_|  \___|\__,_|
//            |_|                     
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Try to match that the given text and message patterns
	Match["Expired"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"connection timed out",
			"could not find a gateway for",
			"delivery attempts will continue to be",
			"delivery expired",
			"delivery time expired",
			"failed to deliver to domain ",
			"giving up on",
			"have been failing for a long time",
			"has been delayed",
			"it has not been collected after",
			"message expired after sitting in queue for",
			"message expired, cannot connect to remote server",
			"message expired, connection refulsed",
			"message timed out",
			"retry time not reached for any host after a long failure period",
			"server did not respond",
			"this message has been in the queue too long",
			"unable to deliver message after multiple retries",
			"was not reachable within the allowed queue period",
			"your message could not be delivered for more than",
		}
		pairs := [][]string{
			[]string{"could not be delivered for", " days"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "expired" or not
	Truth["Expired"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is expired, false: is not expired
		return false
	}
}

