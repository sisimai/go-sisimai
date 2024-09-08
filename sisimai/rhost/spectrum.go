// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

import "sisimai/sis"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["NAME"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		return ""
	}
}

