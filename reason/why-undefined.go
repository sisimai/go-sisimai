// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _           _       __ _                _ 
// | | | |_ __   __| | ___ / _(_)_ __   ___  __| |
// | | | | '_ \ / _` |/ _ \ |_| | '_ \ / _ \/ _` |
// | |_| | | | | (_| |  __/  _| | | | |  __/ (_| |
//  \___/|_| |_|\__,_|\___|_| |_|_| |_|\___|\__,_|
import "libsisimai.org/sisimai/sis"

func init() {
	IncludedIn["Undefined"] = func(argv1 string) bool { return false }
	ProbesInto["Undefined"] = func(fo *sis.Fact) bool { return false }
}

