// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
//import "fmt"
//import "strings"

type EmailAddress struct {
	Address string // Email address
	User    string // Local part of the email addres
	Host    string // Domain part of the email address
	Verp    string // VERP
	Alias   string // Alias of the email address
	Name    string // Display name
	Comment string // (Comment)
}

/*
// Rise() is a constructor of Sisimai::Address
func Rise(argvs []string) (EmailAddress, error) {
	// @param    [[]string] argvs  Email address, name, and other elements
	// @return   [EmailAddress]    EmailAddress struct when the email address was not valid
	if len(argvs["address"]) == 0 { return EmailAddress{}, fmt.Errorf("Email address is empty") }



	return EmailAddress{}, nil
}
*/
