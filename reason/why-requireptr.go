// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                  _          ____ _____ ____  
// |  _ \ ___  __ _ _   _(_)_ __ ___|  _ \_   _|  _ \ 
// | |_) / _ \/ _` | | | | | '__/ _ \ |_) || | | |_) |
// |  _ <  __/ (_| | |_| | | | |  __/  __/ | | |  _ < 
// |_| \_\___|\__, |\__,_|_|_|  \___|_|    |_| |_| \_\
//               |_|                                  

package reason
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/smtp/status"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["RequirePTR"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"access denied. ip name lookup failed",
			"all mail servers must have a ptr record with a valid reverse dns entry",
			"bad dns ptr resource record",
			"cannot find your hostname",
			"cannot resolve your address.",
			"client host rejected: cannot find your hostname", // Yahoo!
			"fix reverse dns for ",
			"ips with missing ptr records",
			"no ptr record found.",
			"please get a custom reverse dns name from your isp for your host",
			"ptr record setup",
			"reverse dns failed",
			"reverse dns required",
			"sender ip reverse lookup rejected",
			"the ip address sending this message does not have a ptr record setup", // Google
			"the corresponding forward dns entry does not point to the sending ip", // Google
			"this system will not accept messages from servers/devices with no reverse dns",
			"unresolvable relay host name",
			"we do not accept mail from hosts with dynamic ip or generic dns ptr-records",
		}
		pairs := [][]string{
			[]string{"domain "," mismatches client ip"},
			[]string{"dns lookup failure: ", " try again later"},
			[]string{"reverse dns lookup for host ", " failed permanently"},
			[]string{"server access ", " forbidden by invalid rdns record of your mail server"},
			[]string{"service permits ", " unverifyable sending ips"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "requireptr" or not
	ProbesInto["RequirePTR"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is requireptr, false: is not requireptr
		if fo == nil                                      { return false }
		if fo.Reason == "requireptr"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "requireptr" { return true  }
		return IncludedIn["RequirePTR"](strings.ToLower(fo.DiagnosticCode))
	}
}

