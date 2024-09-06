// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____       _ _                        _ 
// |  _ \  ___| (_)_   _____ _ __ ___  __| |
// | | | |/ _ \ | \ \ / / _ \ '__/ _ \/ _` |
// | |_| |  __/ | |\ V /  __/ | |  __/ (_| |
// |____/ \___|_|_| \_/ \___|_|  \___|\__,_|
//                                          
import "sisimai/sis"

func init() {
	Match["AuthFailure"] = func(argv1 string) bool { return false }
	Truth["AuthFailure"] = func(fo *sis.Fact) bool { return false }
}

