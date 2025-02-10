// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____       _ _                        _ 
// |  _ \  ___| (_)_   _____ _ __ ___  __| |
// | | | |/ _ \ | \ \ / / _ \ '__/ _ \/ _` |
// | |_| |  __/ | |\ V /  __/ | |  __/ (_| |
// |____/ \___|_|_| \_/ \___|_|  \___|\__,_|
import "libsisimai.org/sisimai/sis"

func init() {
	IncludedIn["Delivered"] = func(argv1 string) bool { return false }
	ProbesInto["Delivered"] = func(fo *sis.Fact) bool { return false }
}

