// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

import "sisimai/sis"

func init() {
	// Decode bounce messages from
	InquireFor["InterScanMSS"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }

        return sis.RisingUnderway{}
    }
}

//       _               _      _____       _            ____                  __  __ ____ ____  
//  _ __| |__   ___  ___| |_   / /_ _|_ __ | |_ ___ _ __/ ___|  ___ __ _ _ __ |  \/  / ___/ ___| 
// | '__| '_ \ / _ \/ __| __| / / | || '_ \| __/ _ \ '__\___ \ / __/ _` | '_ \| |\/| \___ \___ \ 
// | |  | | | | (_) \__ \ |_ / /  | || | | | ||  __/ |   ___) | (_| (_| | | | | |  | |___) |__) |
// |_|  |_| |_|\___/|___/\__/_/  |___|_| |_|\__\___|_|  |____/ \___\__,_|_| |_|_|  |_|____/____/ 
//                                                                                               
