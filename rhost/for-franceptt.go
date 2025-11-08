// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      _______                         ____ _____ _____ 
//  _ __| |__   ___  ___| |_   / /  ___| __ __ _ _ __   ___ ___|  _ \_   _|_   _|
// | '__| '_ \ / _ \/ __| __| / /| |_ | '__/ _` | '_ \ / __/ _ \ |_) || |   | |  
// | |  | | | | (_) \__ \ |_ / / |  _|| | | (_| | | | | (_|  __/  __/ | |   | |  
// |_|  |_| |_|\___/|___/\__/_/  |_|  |_|  \__,_|_| |_|\___\___|_|    |_|   |_|  

package rhost
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["FrancePTT"] = func(fo *siba.Fact) string {
		// - https://www.postmastery.com/orange-postmaster-smtp-error-codes-ofr/
		// - https://smtpfieldmanual.com/provider/orange
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		errorcodes := map[string]string{
			// - 550 5.7.1 Service unavailable; client [192.0.2.1] blocked using Spamhaus
			//   Les emails envoyes vers la messagerie Laposte.net ont ete bloques par nos services.
			//   Afin de regulariser votre situation, nous vous invitons a cliquer sur le lien ci-dessous
			//   et a suivre la procedure.
			// - The emails sent to the mail host Laposte.net were blocked by our services. To regularize
			//   your situation please click on the link below and follow the procedure
			//   https://www.spamhaus.org/lookup/ LPNAAA_101 (in reply to RCPT TO command))
			"101": eb.ReBLOC,

			// - 550 mwinf5c04 ME Adresse IP source bloquee pour incident de spam.
			// - Client host blocked for spamming issues. OFR006_102 Ref http://csi.cloudmark.com ...
			// - 550 5.5.0 Les emails envoyes vers la messagerie Laposte.net ont ete bloques par nos
			//   services. Afin de regulariser votre situation, nous vous invitons a cliquer sur le lien
			//   ci-dessous et a suivre la procedure.
			// - The emails sent to the mail host Laposte.net were blocked by our services. To regularize
			//   your situation please click on the link below and follow the procedure
			//   https://senderscore.org/blacklistlookup/  LPN007_102
			"102": eb.ReBLOC,

			// - 550 mwinf5c10 ME Service refuse. Veuillez essayer plus tard.
			// - Service refused, please try later. OFR006_103 192.0.2.1 [103]
			"103": eb.ReBLOC,

			// - 421 mwinf5c79 ME Trop de connexions, veuillez verifier votre configuration.
			// - Too many connections, slow down. OFR005_104 [104]
			// - Too many connections, slow down. LPN105_104
			"104": eb.ReCONN,

			"105": "", // Veuillez essayer plus tard.
			"107": "", // Service refused, please try later. LPN006_107
			"108": "", // service refused, please try later. LPN001_108
			"109": "", // Veuillez essayer plus tard. LPN003_109
			"201": "", // Veuillez essayer plus tard. OFR004_201

			// - 550 5.7.0 Code d'authentification invalide OFR_305
			"305": eb.ReSECU,

			// - 550 5.5.0 SPF: *** is not allowed to send mail. LPN004_401
			"401": eb.ReAUTH,

			// - 550 5.5.0 Authentification requise. Authentication Required. LPN105_402
			"402": eb.ReSECU,

			// - 5.0.1 Emetteur invalide. Invalid Sender.
			"403": eb.ReFROM,

			// - 5.0.1 Emetteur invalide. Invalid Sender. LPN105_405
			// - 501 5.1.0 Emetteur invalide. Invalid Sender. OFR004_405 [405] (in reply to MAIL FROM command))
			"405": eb.ReFROM,

			// Emetteur invalide. Invalid Sender. OFR_415
			"415": eb.ReFROM,

			// - 550 5.1.1 Adresse d au moins un destinataire invalide.
			// - Invalid recipient. LPN416 (in reply to RCPT TO command)
			// - Invalid recipient. OFR_416 [416] (in reply to RCPT TO command)
			"416": eb.ReUSER,

			// - 552 5.1.1 Boite du destinataire pleine.
			// - Recipient overquota. OFR_417 [417] (in reply to RCPT TO command))
			"417": eb.ReFULL,

			// - Adresse d au moins un destinataire invalide

			// - 550 5.5.0 Boite du destinataire archivee.
			// - Archived recipient. LPN007_420 (in reply to RCPT TO command)
			"420": eb.ReQUIT,

			// - 5.5.3 Mail from not owned by user. LPN105_421.
			"421": eb.ReFROM,

			"423": "", // Service refused, please try later. LPN105_423
			"424": "", // Veuillez essayer plus tard. LPN105_424

			// - 550 5.5.0 Le compte du destinataire est bloque. The recipient account isblocked.
			//   LPN007_426 (in reply to RCPT TO command)
			"426": eb.ReQUIT,

			// - 421 4.2.0 Service refuse. Veuillez essayer plus tard. Service refused, please try later.
			//   OFR005_505 [505] (in reply to end of DATA command)
			// - 421 4.2.1 Service refuse. Veuillez essayer plus tard. Service refused, please try later.
			//   LPN007_505 (in reply to end of DATA command)
			"505": eb.ReSYSE,

			// - Mail rejete. Mail rejected. OFR_506 [506]
			"506": eb.ReSPAM,

			// - 550 5.5.0 Service refuse. Veuillez essayer plus tard. service refused, please try later.
			//   LPN005_510 (in reply to end of DATA command)
			"510": eb.ReBLOC,

			"513": "", // Mail rejete. Mail rejected. OUK_513

			// - Taille limite du message atteinte
			"514": eb.ReSIZE,

			// - 571 5.7.1 Message refused, DMARC verification Failed.
			// - Message refuse, verification DMARC en echec LPN007_517
			"517": eb.ReAUTH,

			// - 554 5.7.1 Client host rejected LPN000_630
			"630": eb.RePOLI,

			// - 421 mwinf5c77 ME Service refuse. Veuillez essayer plus tard. Service refused, please try
			//   later. OFR_999 [999]
			"999": eb.ReBLOC,
		}
		messagesof := map[string][]string{
			eb.ReAUTH: []string{
				// - 421 smtp.orange.fr [192.0.2.1] Emetteur invalide, Veuillez verifier la configuration
				//   SPF/DNS de votre nom de domaine. Invalid Sender. SPF check failed, please verify the
				//   SPF/DNS configuration for your domain name.
				"spf/dns de votre nom de domaine",
			},
		}
		issuedcode := strings.ToLower(strings.ReplaceAll(fo.DiagnosticCode, "_", "-"))
		labelindex := -1
		errorlabel := ""

		for _, e := range []string{"lpn", "lpnaaa", "ofr", "ouk"} {
			// Try to find an error code prefix like "LPN"
			labelindex = strings.LastIndex(issuedcode, e);  if labelindex < 0 { continue }
			errorlabel = e; if strings.Contains(issuedcode, e + "-") { errorlabel += "-" }
			break
		}

		if errorlabel != "" {
			// There is a label (like "LPN") in the error message
			lhs, _, _ := strings.Cut(issuedcode[labelindex + len(errorlabel) - 1:], " ")
			for e := range errorcodes {
				// The key is a code number like "525"
				if strings.HasSuffix(lhs, e) { return errorcodes[e] }
			}
		}

		// There is no error label in the error message
		for e := range messagesof {
			// The key name is a bounce reason name
			if moji.ContainsAny(issuedcode, messagesof[e]) { return e }
		}
		return ""
	}
}

