module libsisimai.org/sisimai

go 1.22

replace (
	sisimai/address => ./sisimai/address
	sisimai/fact => ./sisimai/fact
	sisimai/lhost => ./sisimai/lhost
	sisimai/mail => ./sisimai/mail
	sisimai/message => ./sisimai/message
	sisimai/reason => ./sisimai/reason
	sisimai/rfc1123 => ./sisimai/rfc1123
	sisimai/rfc1894 => ./sisimai/rfc1894
	sisimai/rfc2045 => ./sisimai/rfc2045
	sisimai/rfc3464 => ./sisimai/rfc3464
	sisimai/rfc5322 => ./sisimai/rfc5322
	sisimai/rfc5965 => ./sisimai/rfc5965
	sisimai/rhost => ./sisimai/rhost
	sisimai/sis => ./sisimai/sis
	sisimai/smtp/command => ./sisimai/smtp/command
	sisimai/smtp/failure => ./sisimai/smtp/failure
	sisimai/smtp/reply => ./sisimai/smtp/reply
	sisimai/smtp/status => ./sisimai/smtp/status
	sisimai/smtp/transcript => ./sisimai/smtp/transcript
	sisimai/string => ./sisimai/string
)

require (
	sisimai/address v0.0.0-00010101000000-000000000000
	sisimai/fact v0.0.0-00010101000000-000000000000
	sisimai/mail v0.0.0-00010101000000-000000000000
	sisimai/string v0.0.0-00010101000000-000000000000
)

require (
	sisimai/lhost v0.0.0-00010101000000-000000000000 // indirect
	sisimai/message v0.0.0-00010101000000-000000000000 // indirect
	sisimai/reason v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc1123 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc1894 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc2045 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc3464 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc5322 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc5965 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rhost v0.0.0-00010101000000-000000000000 // indirect
	sisimai/sis v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/command v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/failure v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/reply v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/status v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/transcript v0.0.0-00010101000000-000000000000 // indirect
)
