// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  ____  _____ ____ ____ _________  ____     ______        _         
// |  _ \|  ___/ ___| ___|___ /___ \|___ \   / /  _ \  __ _| |_ ___ _ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) | / /| | | |/ _` | __/ _ (_)
// |  _ <|  _|| |___ ___) |__) / __/ / __/ / / | |_| | (_| | ||  __/_ 
// |_| \_\_|   \____|____/____/_____|_____/_/  |____/ \__,_|\__\___(_)
/**************************************************************************************************
3.3.  Date and Time Specification

  Date and time values occur in several header fields. This section specifies the syntax for a full
  date and time specification. Though folding white space is permitted throughout the date-time
  specification, it is RECOMMENDED that a single space be used in each place that FWS appears
  (whether it is required or optional); some older implementations will not interpret longer se-
  quences of folding white space correctly.

    date-time       =   [ day-of-week "," ] date time [CFWS]
    day-of-week     =   ([FWS] day-name) / obs-day-of-week
    day-name        =   "Mon" / "Tue" / "Wed" / "Thu" / "Fri" / "Sat" / "Sun"
    date            =   day month year
    day             =   ([FWS] 1*2DIGIT FWS) / obs-day
    month           =   "Jan" / "Feb" / "Mar" / "Apr" / "May" / "Jun" / "Jul" / "Aug" /
                        "Sep" / "Oct" / "Nov" / "Dec"
    year            =   (FWS 4*DIGIT FWS) / obs-year
    time            =   time-of-day zone
    time-of-day     =   hour ":" minute [ ":" second ]
    hour            =   2DIGIT / obs-hour
    minute          =   2DIGIT / obs-minute
    second          =   2DIGIT / obs-second
    zone            =   (FWS ( "+" / "-" ) 4DIGIT) / obs-zone

  The day is the numeric day of the month. The year is any numeric year 1900 or later.

  The time-of-day specifies the number of hours, minutes, and optionally seconds since midnight of
  the date indicated.

  The date and time-of-day SHOULD express local time.

  The zone specifies the offset from Coordinated Universal Time (UTC, formerly referred to as GMT
  that the date and time-of-day represent. The "+" or "-" indicates whether the time-of-day is a-
  head of (i.e., east of) or behind (i.e., west of) Universal Time. The first two digits indicate
  the number of hours difference from Universal Time, and the last two digits indicate the number
  of additional minutes difference from Universal Time. (Hence, +hhmm means +(hh * 60 + mm) minutes,
  and -hhmm means -(hh * 60 + mm) minutes). The form "+0000" SHOULD be used to indicate a time zone
  at Universal Time. Though "-0000" also indicates Universal Time, it is used to indicate that the
  time was generated on a system that may be in a local time zone other than Universal Time and
  that the date-time contains no information about the local time zone.

  A date-time specification MUST be semantically valid. That is, the day-of-week (if included) MUST
  be the day implied by the date, the numeric day-of-month MUST be between 1 and the number of days
  allowed for the specified month (in the specified year), the time-of-day MUST be in the range
  00:00:00 through 23:59:60 (the number of seconds allowing for a leap second; see [RFC1305]), and
  the last two digits of the zone MUST be within the range 00 through 59.
**************************************************************************************************/
import "fmt"
import "strings"
import "strconv"
import sisimoji "libsisimai.org/sisimai/string"

var MonthName = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var DayOfWeek = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

// Date() tidies up and converts the date string to the format 
func Date(argv1 string) string {
	// @param    string argv1  Date string
	// @return   string        Tidied date string
	// @see      https://www.rfc-editor.org/rfc/rfc5322#section-3.3
	// @see      https://www.ietf.org/rfc/rfc3339.txt
	// @see      https://en.wikipedia.org/wiki/ISO_8601
	// @example  Tidy up the date string and convert to RFC822-formatted date string
	//   rfc5322.Date("2018-02-02T18:30:22 Fri") => Fri, 2 Feb 2018 18:30:22 +0000
	//   rfc5322.Date("Fri, Feb 2 2018 2:2:2")   => Fri, 2 Feb 2018 02:02:02 +0000
	if argv1 == "" { return "" }

	datestring := sisimoji.Sweep(strings.ReplaceAll(argv1, ",", ", ")) // "Thu,22" -> "Thu, 22"
	year2digit := uint8(0) // 2-digit year such as 22, 97
	p          := [6]string{
		"", // [0] Year (2018)
		"", // [1] Month (Feb)
		"", // [2] Day (2)
		"", // [3] Day of week (Fri)
		"", // [4] Time (18:30:22)
		"", // [5] Timezone Offset (0)
	}

	for _, e := range strings.Split(datestring, " ") {
		// Check, convert each piece and store it to p[*]
		cw := len(e); if cw == 0 { continue }
		if cw < 3 {
			// This piece might be a day such as 1, or 02, or 31
			cv, nyaan := strconv.ParseUint(e, 10, 8); if nyaan != nil {
				// Failed to parse as a integer
				return ""

			} else {
				// Successfully parsed and convertd to an interger
				if cv > 31 || cv == 0 {
					// 2-digit year ?
					year2digit = uint8(cv)

				} else {
					// Deal as a day of month
					p[2] = fmt.Sprintf("%02d", cv)
				}
			}
		} else if cw == 3 || (cw == 4 && strings.HasSuffix(e, ",")) {
			// 3 characters: "Feb" or "Thu" or "Thu,", or 3-digit date like "029"
			if sisimoji.ContainsOnlyNumbers(e) && strings.HasPrefix(e, "0") {
				// Tue, 029 Apr 2019 23:34:45 -0800 (PST)
				p[2] = e[1:]

			} else {
				upperfirst := strings.ToUpper(e[0:1]) + strings.ToLower(e[1:3])
				if sisimoji.EqualsAny(upperfirst, MonthName) { p[1] = upperfirst; continue }
				if sisimoji.EqualsAny(upperfirst, DayOfWeek) { p[3] = upperfirst; continue }
			}

		} else if cw == 4 {
			// This piece might be a 4-digit year such as 1997, 2018
			cv, nyaan := strconv.ParseUint(e, 10, 16); if nyaan != nil { continue }
			p[0] = fmt.Sprintf("%04d", cv)

		} else if cw == 5 && (strings.HasPrefix(e, "+") || strings.HasPrefix(e, "-")) {
			// This piece might be a timezone offset such as "+0900", "-0400"
			cv, nyaan := strconv.ParseUint(e[1:5], 10, 16); if nyaan != nil { continue }
			p[5] = fmt.Sprintf("%s%04d", e[0:1], cv)

		} else if cw > 5 {
			// Time string such as "18:30:22" or other formatted string
			if strings.Count(e, ":") == 2 {
				// This piece might be a time such as "18:30:22", "3:1:4"
				ct := []uint8{}
				for _, f := range strings.Split(e, ":") {
					// Each element(integer) should be greater equal 0 and less equal 60.
					cv, nyaan := strconv.ParseUint(f, 10, 8); if nyaan != nil || cv > 60 {
						// This piece does not seem to a time string
						return ""
					}
					ct = append(ct, uint8(cv))
				}
				if len(ct) != 3 {
					// This piece does not seem to a time string
					return ""
				}
				p[4] = fmt.Sprintf("%02d:%02d:%02d", ct[0], ct[1], ct[2])

			} else {
				// NOTE: Should we support other formatted date string like the followings?
				// - "Sun, 29 May 2014 1:2 +0900"
				// - "4/29/01 11:34:45 PM",
				// - "2014-03-26 00-01-19",
				// - "29-04-2017 22:22",
				continue
			}
		}
	}

	if p[0] == "" && year2digit > 0 {
		// 4-digit year string is empty, try to use 2-digit year instead.
		yy  := "20"; if year2digit > 81 { yy = "19" } // RFC822 is published August, 1982
		p[0] = fmt.Sprintf("%s%2d", yy, int(year2digit))
	}
	if p[3] == "" { p[3] = "Thu"   } // Set "Thu" for now
	if p[5] == "" { p[5] = "+0000" } // Deal as UTC

	// Date: Thu, 22 Feb 2022 22:22:22 +0200
	if p[0] == "" || p[1] == "" || p[2] == "" || p[4] == "" { return "" }
	return fmt.Sprintf("%s, %s %s %s %s %s", p[3], p[2], p[1], p[0], p[4], p[5])
}

