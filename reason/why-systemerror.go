// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _                 _____                     
// / ___| _   _ ___| |_ ___ _ __ ___ | ____|_ __ _ __ ___  _ __ 
// \___ \| | | / __| __/ _ \ '_ ` _ \|  _| | '__| '__/ _ \| '__|
//  ___) | |_| \__ \ ||  __/ | | | | | |___| |  | | | (_) | |   
// |____/ \__, |___/\__\___|_| |_| |_|_____|_|  |_|  \___/|_|   
//        |___/                                                 

package reason
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.RePROC] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"aliasing/forwarding loop broken",
			"automatic homedir creator crashed", // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"can't create user output file",
			"cannot send e-mail to yourself",
			"could not load ",
			"input/output error",
			"interrupted system call",
			"it encountered an error while being processed",
			"it would create a mail loop",
			"ldap attribute", // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"ldap lookup",    // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"ldap server",    // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"loop was found in the mail exchanger",
			"loops back to myself",
			"mail transport unavailable",
			"no such file or directory",
			"error while executing qmail-forward", // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"queue file write error",
			"recipient deferred because there is no mdb",
			"remote server is misconfigured",
			"service currently unavailable",
			"several matches found in domino directory", // Donimo
			"temporary local problem",
			"timeout waiting for input",
			"too many results returned but needs to be unique", // qmail-ldap-1.03-20040101.patch:19817 - 19866
			"transaction failed ",
		}
		pairs := [][]string{
			[]string{"config", " error"},
			[]string{"internal ", "error"},
			[]string{"local ", "error"},
			[]string{"proxy", "broken pipe"},
			[]string{"unable to connect ", "daemon"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.RePROC] = func(fo *siba.Fact) bool { return false }
}

