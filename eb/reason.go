// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _       ______                            
//   ___| |__   / /  _ \ ___  __ _ ___  ___  _ __  
//  / _ \ '_ \ / /| |_) / _ \/ _` / __|/ _ \| '_ \ 
// |  __/ |_) / / |  _ <  __/ (_| \__ \ (_) | | | |
//  \___|_.__/_/  |_| \_\___|\__,_|___/\___/|_| |_|

package eb
const (
	// bounce reason names
	ReAUTH = "AuthFailure"
	ReREPU = "BadReputation"
	ReBLOC = "Blocked"
	ReBODY = "ContentError"
	ReSENT = "Delivered"
	ReSIZE = "EmailTooLarge"
	ReTIME = "Expired"
	ReTTLS = "FailedSTARTTLS"
	ReFEED = "Feedback"
	ReFILT = "Filtered"
	ReMOVE = "HasMoved"
	ReHOST = "HostUnknown"
	ReFULL = "MailboxFull"
	ReUNIX = "MailerError"
	ReNETW = "NetworkError"
	ReRELA = "NoRelaying"
	Re00MX = "NotAccept"
	ReNRFC = "NotCompliantRFC"
	Re___1 = "OnHold"
	RePOLI = "PolicyViolation"
	ReFROM = "Rejected"
	ReQPTR = "RequirePTR"
	ReRATE = "RateLimited"
	ReSECU = "SecurityError"
	ReSPAM = "SpamDetected"
	ReSUPP = "Suppressed"
	ReQUIT = "Suspend"
	ReCOMM = "SyntaxError"
	ReSYSE = "SystemError"
	ReSYSF = "SystemFull"
	Re___0 = "Undefined"
	ReUSER = "UserUnknown"
	ReAWAY = "Vacation"
	ReEXEC = "VirusDetected"
)

