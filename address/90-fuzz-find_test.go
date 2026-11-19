//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ___       _     _                   
// |_   _|__  ___| |_   / / \   __| | __| |_ __ ___  ___ ___ 
//   | |/ _ \/ __| __| / / _ \ / _` |/ _` | '__/ _ \/ __/ __|
//   | |  __/\__ \ |_ / / ___ \ (_| | (_| | | |  __/\__ \__ \
//   |_|\___||___/\__/_/_/   \_\__,_|\__,_|_|  \___||___/___/
//                                                           
package address
import "testing"

// % go test -tags=fuzz -fuzz=FuzzFind -fuzztime=1m ./address
func FuzzFind(f *testing.F) {
	emailseeds := []string{
		"------ pipe to |/usr/local/neko/bin/cat kijitora@example.com /home/neko/.cat",
		"   * kijitora@example.co.jp",
		"<<< 550 5.1.1 <kijitora@example.jp>... Invalid recipient",
		"/home/neko/.cat/kijitora@example.com: No such file or directory",
		"Authentication-Results: aneyakoji.example.jp; sender-id=pass header.from=kijitora@example.jp; auth=pass (CRAM-MD5);",
		"by 00000000 (envelope-from <kijitora@example.or.jp>) with ESMTP",
		"Final-Recipient: RFC822; <kijitora@example.net>",
		"for kijitora@example.co.jp; Thu, 29 Apr 2006 23:34:45 +0900",
		"https://example.jp/unsubscribe.html?SubscriptionArn=arn:*&Endpoint=kijitora@example.jp",
		"Kijitora<mailto:kijitora@neko.example.com>",
		"kijitora@example.org: [192.0.2.244], 554 Service currently unavailable",
		"Reason: 542 kijitora@example.com Rejected",
		"Recipient Address:      kijitora@example.com",
		"Received >>> 550 5.1.1 <kijitora@example.co.jp>... user unknown",
		"Return-Path: <>",
		"Sent <<< RCPT TO:<kijitora@example.co.jp>",
		"smtp.mailfrom=3Dkijitora@outlook.example.com;=20",
		"This user doesn't have a example.com account (kijitora@example.com) [0]",
		`To: "" <kijitora@libsisimai.org>`,
		"To:      <kijitora@nyaan.example.net>",
		"To: =?utf-8?Q?kijitora=40example=2Ejp?= <kijitora@example.jp>",
		"To: &quot;kijitora@example.jp&quot; &lt;kijitora@example.jp&gt;",
		"To: Kijitora <kijitora@cat.example.ne.jp>",
		"TO:kijitora-neko-nyaan@ntt.example.ne.jp",
		"X-Display-Name: kijitora@cr.neko.nyaan.jp",
		"X-HmXmrOriginalRecipient: kijitora@example.com",
		"X-Original-To: kijitora-nyaan@em.example.com",
		"X-Postfix-Sender: rfc822; kijitora-nyaan@em.example.com",
		"Your message to kijitora@example.com couldn't be delivered.",
	}
	for _, e := range emailseeds { f.Add(e + "\n")          }
	f.Fuzz(func(t *testing.T, seed string) { _ = Find(seed) })
}

