// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____              _             _____                     
// / ___| _   _ _ __ | |_ __ ___  _| ____|_ __ _ __ ___  _ __ 
// \___ \| | | | '_ \| __/ _` \ \/ /  _| | '__| '__/ _ \| '__|
//  ___) | |_| | | | | || (_| |>  <| |___| |  | | | (_) | |   
// |____/ \__, |_| |_|\__\__,_/_/\_\_____|_|  |_|  \___/|_|   
//        |___/                                               

package reason
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReCOMM] = func(mesg string) bool { return false }

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReCOMM] = func(fo *siba.Fact) bool { return false }
}

