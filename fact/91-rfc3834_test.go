// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ______  _____ ____ _____  ___ _____ _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ / ( _ )___ /| || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |     |_ \ / _ \ |_ \| || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) | (_) |__) |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/ \___/____/   |_|  
import "testing"

func TestRFC3834(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "",        "",    "vacation",        false, ""}},
		{{"02",   1, "",        "",    "vacation",        false, ""}},
		{{"03",   1, "",        "",    "vacation",        false, ""}},
		{{"04",   1, "",        "",    "vacation",        false, ""}},
		{{"05",   1, "",        "",    "vacation",        false, ""}},
	}; EngineTest(t, "RFC3834", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "",        "",    "vacation",        false, ""}},
		{{"1002", 1, "",        "",    "vacation",        false, ""}},
		{{"1003", 1, "",        "",    "vacation",        false, ""}},
		{{"1004", 1, "",        "",    "vacation",        false, ""}},
		{{"1005", 1, "",        "",    "vacation",        false, ""}},
		{{"1006", 1, "",        "",    "vacation",        false, ""}},
		{{"1007", 1, "",        "",    "vacation",        false, ""}},
		{{"1008", 1, "",        "",    "vacation",        false, ""}},
		{{"1009", 1, "",        "",    "vacation",        false, ""}},
		{{"1010", 1, "",        "",    "vacation",        false, ""}},
		{{"1011", 1, "",        "",    "vacation",        false, ""}},
		{{"1012", 1, "",        "",    "vacation",        false, ""}},
		{{"1013", 1, "",        "",    "vacation",        false, ""}},

	}; EngineTest(t, "RFC3834", secretlist, false)
}

