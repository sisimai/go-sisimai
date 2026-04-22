// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _     
//   ___| |__  
//  / _ \ '_ \ 
// |  __/ |_) |
//  \___|_.__/ 

// Package "eb" provides constants for the email bounce.
package eb
const XeBYTE = 2000 * 1024 * 1024 * 1024 // 2GB: The maximum bytes of the email size in mail/lib.go
const GeFrom = "MAILER-DAEMON Fri Feb  2 18:30:22 2018"
var FeRFC822 = []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}

