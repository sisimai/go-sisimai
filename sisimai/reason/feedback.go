// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____             _ _                _    
// |  ___|__  ___  __| | |__   __ _  ___| | __
// | |_ / _ \/ _ \/ _` | '_ \ / _` |/ __| |/ /
// |  _|  __/  __/ (_| | |_) | (_| | (__|   < 
// |_|  \___|\___|\__,_|_.__/ \__,_|\___|_|\_\
import "sisimai/sis"

func init() {
	Match["Feedback"] = func(argv1 string) bool { return false }
	Truth["Feedback"] = func(fo *sis.Fact) bool { return false }
}

