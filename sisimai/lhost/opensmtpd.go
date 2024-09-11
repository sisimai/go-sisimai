// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

import "sisimai/sis"

func init() {
	// Decode bounce messages from
	InquireFor["OpenSMTPD"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }

        return sis.RisingUnderway{}
    }
}

//       _               _      _____                   ____  __  __ _____ ____  ____  
//  _ __| |__   ___  ___| |_   / / _ \ _ __   ___ _ __ / ___||  \/  |_   _|  _ \|  _ \ 
// | '__| '_ \ / _ \/ __| __| / / | | | '_ \ / _ \ '_ \\___ \| |\/| | | | | |_) | | | |
// | |  | | | | (_) \__ \ |_ / /| |_| | |_) |  __/ | | |___) | |  | | | | |  __/| |_| |
// |_|  |_| |_|\___/|___/\__/_/  \___/| .__/ \___|_| |_|____/|_|  |_| |_| |_|   |____/ 
//                                    |_|                                              
