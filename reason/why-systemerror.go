// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _                 _____                     
// / ___| _   _ ___| |_ ___ _ __ ___ | ____|_ __ _ __ ___  _ __ 
// \___ \| | | / __| __/ _ \ '_ ` _ \|  _| | '__| '__/ _ \| '__|
//  ___) | |_| \__ \ ||  __/ | | | | | |___| |  | | | (_) | |   
// |____/ \__, |___/\__\___|_| |_| |_|_____|_|  |_|  \___/|_|   
//        |___/                                                 

package reason
import "strings"
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["SystemError"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"aliasing/forwarding loop broken",
			"can't create user output file",
			"could not load drd for domain",
			"internal error reading data", // Microsoft
			"internal server error: operation now in progress", // Microsoft
			"interrupted system call",
			"it encountered an error while being processed",
			"it would create a mail loop",
			"local configuration error",
			"local error in processing",
			"loop was found in the mail exchanger",
			"loops back to myself",
			"mail system configuration error",
			"queue file write error",
			"recipient deferred because there is no mdb",
			"remote server is misconfigured",
			"server configuration error",
			"service currently unavailable",
			"system config error",
			"temporary local problem",
			"timeout waiting for input",
			"transaction failed ",
		}
		pairs := [][]string{
			[]string{"unable to connect ", "daemon"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "systemerror" or not
	ProbesInto["SystemError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is systemerror, false: is not systemerror
		return false
	}
}

