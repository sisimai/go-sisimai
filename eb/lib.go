// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _     
//   ___| |__  
//  / _ \ '_ \ 
// |  __/ |_) |
//  \___|_.__/ 

// Package "eb" provides constants for the email bounce.
package eb

// Max depth of MIME parts
// - sendmail-8.18.1/sendmail/conf.h:#define MAXMIMENESTING	20	/* max MIME multipart nesting */
// - postfix-3.7.2/proto/postconf.proto:%PARAM mime_nesting_limit 100
const XdMIME = 64
const XeBYTE = 128 * 1024 * 1024 // 128MB: The maximum size of each email file in Maildir/
const GeFrom = "MAILER-DAEMON Fri Feb  2 18:30:22 2018"
var FeRFC822 = []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}

