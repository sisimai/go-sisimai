// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

// Retry() returns the table of reason list which should be checked again
func Retry() map[string]bool {
	return map[string]bool{
		"undefined": true, "onhold": true,  "systemerror": true, "securityerror": true, "expired": true,
		"suspend": true, "networkerror": true, "hostunknown": true, "userunknown": true,
    }
}

