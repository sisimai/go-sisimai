// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

import "sisimai/sis"

func init() {
	// Decode bounce messages from
	InquireFor["MailMarshalSMTP"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }

        return sis.RisingUnderway{}
    }
}

//       _               _      ____  __       _ _ __  __                _           _ ____  __  __ _____ ____  
//  _ __| |__   ___  ___| |_   / /  \/  | __ _(_) |  \/  | __ _ _ __ ___| |__   __ _| / ___||  \/  |_   _|  _ \ 
// | '__| '_ \ / _ \/ __| __| / /| |\/| |/ _` | | | |\/| |/ _` | '__/ __| '_ \ / _` | \___ \| |\/| | | | | |_) |
// | |  | | | | (_) \__ \ |_ / / | |  | | (_| | | | |  | | (_| | |  \__ \ | | | (_| | |___) | |  | | | | |  __/ 
// |_|  |_| |_|\___/|___/\__/_/  |_|  |_|\__,_|_|_|_|  |_|\__,_|_|  |___/_| |_|\__,_|_|____/|_|  |_| |_| |_|    
//                                                                                                              