// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _       ______                                          _ 
//   ___| |__   / / ___|___  _ __ ___  _ __ ___   __ _ _ __   __| |
//  / _ \ '_ \ / / |   / _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// |  __/ |_) / /| |__| (_) | | | | | | | | | | | (_| | | | | (_| |
//  \___|_.__/_/  \____\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|

package eb
const (
	CeHELO = "HELO"
	CeEHLO = "EHLO"
	CeMAIL = "MAIL"
	CeRCPT = "RCPT"
	CeDATA = "DATA"
	CeQUIT = "QUIT"
	CeRSET = "RSET"
	CeNOOP = "NOOP"
	CeVRFY = "VRFY"
	CeETRN = "ETRN"
	CeEXPN = "EXPN"
	CeHELP = "HELP"
	CeAUTH = "AUTH"
	CeTTLS = "STARTTLS"
	CeXFWD = "XFORWARD"
	CeCONN = "CONN" // CONN is a pseudo SMTP command used only in Sisimai
)

