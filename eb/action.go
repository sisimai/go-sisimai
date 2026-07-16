// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _       ___        _   _             
//   ___| |__   / / \   ___| |_(_) ___  _ __  
//  / _ \ '_ \ / / _ \ / __| __| |/ _ \| '_ \ 
// |  __/ |_) / / ___ \ (__| |_| | (_) | | | |
//  \___|_.__/_/_/   \_\___|\__|_|\___/|_| |_|

package eb
// https://datatracker.ietf.org/doc/html/rfc3464#page-16
// 2.3.3 Action field
//   The Action field indicates the action performed by the Reporting-MTA as a result of its attempt
//   to deliver the message to this recipient address. This field MUST be present for each recipient
//   named in the DSN.
//
//   The syntax for the action-field is:
//     action-field = "Action" ":" action-value
//     action-value = "failed" / "delayed" / "delivered" / "relayed" / "expanded"
//
//   The action-value may be spelled in any combination of upper and lower case characters.
//
//     "failed"    indicates that the message could not be delivered to the recipient.
//                 The Reporting MTA has abandoned any attempts to deliver the message to this
//                 recipient. No further notifications should be expected.
//
//     "delayed"   indicates that the Reporting MTA has so far been unable to deliver or relay the
//                 message, but it will continue to attempt to do so. Additional notification
//                 messages may be issued as the message is further delayed or successfully
//                 delivered, or if delivery attempts are later abandoned.
//
//     "delivered" indicates that the message was successfully delivered to the recipient address
//                 specified by the sender, which includes "delivery" to a mailing list exploder.
//                 It does not indicate that the message has been read. This is a terminal state
//                 and no further DSN for this recipient should be expected.
//
//     "relayed"   indicates that the message has been relayed or gatewayed into an environment
//                 that does not accept responsibility for generating DSNs upon successful delivery.
//                 This action-value SHOULD NOT be used unless the sender has requested notification
//                 of successful delivery for this recipient.
//
//     "expanded"  indicates that the message has been successfully delivered to the recipient
//                 address as specified by the sender, and forwarded by the Reporting-MTA beyond
//                 that destination to multiple additional recipient addresses. An action-value of
//                 "expanded" differs from "delivered" in that "expanded" is not a terminal state.
//                 Further "failed" and/or "delayed" notifications may be provided.
const (
	AeFAIL = "failed"
	AeSTAY = "delayed"
	AeSENT = "delivered"
	AePASS = "relayed"
	AeEXPN = "expanded"
)

