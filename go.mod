module libsisimai.org/sisimai

replace (
	sisimai/address => ./sisimai/address
	sisimai/fact => ./sisimai/fact
	sisimai/lhost => ./sisimai/lhost
	sisimai/mail => ./sisimai/mail
	sisimai/rfc1894 => ./sisimai/rfc1894
	sisimai/rfc2045 => ./sisimai/rfc2045
	sisimai/rfc5322 => ./sisimai/rfc5322
	sisimai/smtp/command => ./sisimai/smtp/command
	sisimai/smtp/reply => ./sisimai/smtp/reply
	sisimai/smtp/status => ./sisimai/smtp/status
	sisimai/smtp/transcript => ./sisimai/smtp/transcript
	sisimai/string => ./sisimai/string
)

go 1.20

require (
	sisimai/address v0.0.0-00010101000000-000000000000
	sisimai/fact v0.0.0-00010101000000-000000000000
	sisimai/lhost v0.0.0-00010101000000-000000000000 // indirect
	sisimai/mail v0.0.0-00010101000000-000000000000
	sisimai/rfc1894 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc2045 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/rfc5322 v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/command v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/reply v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/status v0.0.0-00010101000000-000000000000 // indirect
	sisimai/smtp/transcript v0.0.0-00010101000000-000000000000 // indirect
	sisimai/string v0.0.0-00010101000000-000000000000 // indirect
)
