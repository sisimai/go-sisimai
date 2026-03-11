// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "",        "",    "vacation",        false, 0, ""}},
		{{"02",   1, "",        "",    "vacation",        false, 0, ""}},
		{{"03",   1, "",        "",    "vacation",        false, 0, ""}},
		{{"04",   1, "",        "",    "vacation",        false, 0, ""}},
		{{"05",   1, "",        "",    "vacation",        false, 0, ""}},
		{{"06",   1, "5.9.221", "",    "suspend",         false, 1, ""}},
	}; EngineTest(t, "RFC3834", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1002", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1003", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1004", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1005", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1006", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1007", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1008", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1009", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1010", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1011", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1012", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1013", 1, "",        "",    "vacation",        false, 0, ""}},
		{{"1014", 1, "5.9.221", "",    "suspend",         false, 1, ""}},
	}; EngineTest(t, "RFC3834", secretlist, false)
}

