// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      _______      _           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   / / ____|_  _(_)_ __ ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __| / /|  _| \ \/ / | '_ ` _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ / / | |___ >  <| | | | | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__/_/  |_____/_/\_\_|_| |_| |_|
//                                                                           
import "testing"

func TestLhostExim(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.7.0",   "550", "blocked",         false, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""},
		 {"02",   2, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.7.0",   "554", "policyviolation", false, ""}},
		{{"04",   1, "5.7.0",   "550", "blocked",         false, ""}},
		{{"05",   1, "5.1.1",   "553", "userunknown",      true, ""}},
		{{"06",   1, "4.0.947", "",    "expired",         false, ""}},
		{{"08",   1, "4.0.947", "",    "expired",         false, ""}},
		{{"29",   1, "5.0.0",   "550", "authfailure",     false, ""}},
		{{"30",   1, "5.7.1",   "554", "userunknown",      true, ""}},
		{{"31",   1, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"32",   1, "5.0.973", "",    "requireptr",      false, ""}},
		{{"33",   1, "5.0.973", "554", "requireptr",      false, ""}},
		{{"34",   1, "5.7.1",   "554", "requireptr",      false, ""}},
		{{"35",   1, "5.0.971", "550", "blocked",         false, ""}},
		{{"36",   1, "5.0.901", "550", "rejected",        false, ""}},
		{{"37",   1, "5.0.912", "553", "hostunknown",      true, ""}},
		{{"38",   1, "4.0.901", "450", "requireptr",      false, ""}},
		{{"39",   1, "5.0.973", "550", "requireptr",      false, ""}},
		{{"40",   1, "5.0.901", "551", "requireptr",      false, ""}},
		{{"41",   1, "4.0.901", "450", "requireptr",      false, ""}},
		{{"42",   1, "5.7.1",   "554", "requireptr",      false, ""}},
		{{"43",   1, "5.7.1",   "550", "requireptr",      false, ""}},
		{{"44",   1, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"45",   1, "5.2.0",   "550", "rejected",        false, ""}},
		{{"46",   1, "5.7.1",   "554", "blocked",         false, ""}},
		{{"47",   1, "5.0.971", "550", "blocked",         false, ""}},
		{{"48",   1, "5.7.1",   "550", "requireptr",      false, ""}},
		{{"49",   1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"50",   1, "5.1.7",   "550", "rejected",        false, ""}},
		{{"51",   1, "5.1.0",   "553", "rejected",        false, ""}},
		{{"52",   1, "5.0.902", "",    "syntaxerror",     false, ""}},
		{{"53",   1, "5.0.939", "",    "mailererror",     false, ""}},
		{{"54",   1, "5.0.901", "550", "blocked",         false, ""}},
		{{"55",   1, "5.7.0",   "554", "spamdetected",    false, ""}},
		{{"56",   1, "5.0.971", "554", "blocked",         false, ""}},
		{{"57",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"58",   1, "5.0.934", "500", "mesgtoobig",      false, ""}},
		{{"59",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"60",   1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"61",   1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Exim", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.7.0",   "554", "policyviolation", false, ""}},
		{{"1002", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1003", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1004", 1, "5.7.0",   "550", "blocked",         false, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1005", 2, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1006", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1007", 1, "5.7.0",   "554", "policyviolation", false, ""}},
		{{"1008", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1009", 1, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"1010", 1, "5.7.0",   "550", "blocked",         false, ""}},
		{{"1011", 1, "5.1.1",   "553", "userunknown",      true, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1013", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1014", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1015", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1016", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1017", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1018", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1019", 1, "5.1.1",   "553", "userunknown",      true, ""}},
		{{"1020", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1022", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1023", 1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1024", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1025", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1026", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1027", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1028", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1029", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1031", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1032", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1033", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1034", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1035", 1, "5.1.8",   "550", "rejected",        false, ""}},
		{{"1036", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1037", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1038", 1, "5.7.0",   "550", "blocked",         false, ""}},
		{{"1039", 1, "4.0.922", "",    "mailboxfull",     false, ""}},
		{{"1040", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1041", 1, "4.0.947", "451", "spamdetected",    false, ""}},
		{{"1042", 1, "5.0.944", "",    "networkerror",    false, ""}},
		{{"1043", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1044", 1, "5.0.944", "",    "networkerror",    false, ""}},
		{{"1045", 1, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"1046", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1047", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1049", 1, "5.0.921", "554", "suspend",         false, ""}},
		{{"1050", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1051", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1053", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1054", 1, "5.0.921", "554", "suspend",         false, ""}},
		{{"1055", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1056", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1057", 1, "5.0.921", "554", "suspend",         false, ""}},
		{{"1058", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1059", 1, "5.0.901", "550", "onhold",          false, ""}},
		{{"1060", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1061", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1062", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1063", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1064", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1065", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1066", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1067", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1068", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1069", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1070", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1071", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1072", 1, "5.2.1",   "554", "userunknown",      true, ""}},
		{{"1073", 1, "5.0.921", "554", "suspend",         false, ""}},
		{{"1074", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1075", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1076", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1077", 1, "5.0.921", "554", "suspend",         false, ""}},
		{{"1078", 1, "5.0.900", "",    "undefined",       false, ""}},
		{{"1079", 1, "5.0.0",   "",    "hostunknown",      true, ""},
		 {"1079", 2, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1080", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1081", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1082", 1, "5.0.901", "",    "onhold",          false, ""}},
		{{"1083", 1, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1084", 1, "5.0.0",   "550", "systemerror",     false, ""},
		 {"1084", 2, "5.0.0",   "550", "systemerror",     false, ""}},
		{{"1085", 1, "5.0.0",   "550", "blocked",         false, ""},
		 {"1085", 2, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1086", 1, "5.0.0",   "",    "onhold",          false, ""},
		 {"1086", 2, "5.0.0",   "",    "onhold",          false, ""},
		 {"1086", 3, "5.0.0",   "",    "onhold",          false, ""},
		 {"1086", 4, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1087", 1, "5.0.0",   "550", "onhold",          false, ""}},
		{{"1088", 1, "5.0.901", "550", "onhold",          false, ""},
		 {"1088", 2, "5.0.0",   "550", "onhold",          false, ""}},
		{{"1089", 1, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1089", 2, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1090", 1, "5.0.0",   "",    "onhold",          false, ""},
		 {"1090", 2, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1091", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1092", 1, "5.0.0",   "",    "undefined",       false, ""}},
		{{"1094", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1095", 1, "5.0.0",   "",    "undefined",       false, ""}},
		{{"1098", 1, "4.0.947", "",    "expired",         false, ""},
		 {"1098", 2, "4.0.947", "",    "expired",         false, ""}},
		{{"1099", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1100", 1, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1100", 2, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1100", 3, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1100", 4, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1100", 5, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1101", 1, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1103", 1, "5.0.900", "",    "undefined",       false, ""},
		 {"1103", 2, "5.0.900", "",    "undefined",       false, ""},
		 {"1103", 3, "5.0.0",   "",    "undefined",       false, ""}},
		{{"1104", 1, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1104", 2, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1105", 1, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1106", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1107", 1, "5.0.980", "",    "spamdetected",    false, ""}},
		{{"1109", 1, "5.7.1",   "554", "userunknown",      true, ""}},
		{{"1110", 1, "5.0.912", "",    "hostunknown",      true, ""},
		 {"1110", 2, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"1111", 1, "5.0.973", "",    "requireptr",      false, ""}},
		{{"1112", 1, "5.0.973", "554", "requireptr",      false, ""}},
		{{"1113", 1, "5.7.1",   "554", "requireptr",      false, ""}},
		{{"1114", 1, "5.0.971", "550", "blocked",         false, ""}},
		{{"1115", 1, "5.0.901", "550", "rejected",        false, ""}},
		{{"1116", 1, "5.0.912", "553", "hostunknown",      true, ""}},
		{{"1117", 1, "4.0.901", "450", "requireptr",      false, ""}},
		{{"1118", 1, "5.0.973", "550", "requireptr",      false, ""}},
		{{"1119", 1, "5.0.901", "551", "requireptr",      false, ""}},
		{{"1120", 1, "4.0.901", "450", "requireptr",      false, ""}},
		{{"1121", 1, "5.7.1",   "554", "requireptr",      false, ""}},
		{{"1122", 1, "5.7.1",   "550", "requireptr",      false, ""}},
		{{"1123", 1, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1124", 1, "5.2.0",   "550", "rejected",        false, ""}},
		{{"1125", 1, "5.7.1",   "554", "blocked",         false, ""}},
		{{"1126", 1, "5.0.971", "550", "blocked",         false, ""}},
		{{"1127", 1, "5.7.1",   "550", "requireptr",      false, ""}},
		{{"1128", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1129", 1, "5.1.7",   "550", "rejected",        false, ""}},
		{{"1130", 1, "5.1.0",   "553", "rejected",        false, ""}},
		{{"1131", 1, "5.0.902", "",    "syntaxerror",     false, ""}},
		{{"1132", 1, "5.0.939", "",    "mailererror",     false, ""}},
		{{"1133", 1, "5.0.901", "550", "blocked",         false, ""}},
		{{"1134", 1, "5.7.0",   "554", "spamdetected",    false, ""}},
		{{"1135", 1, "5.0.971", "554", "blocked",         false, ""}},
		{{"1136", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1137", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1138", 1, "5.0.901", "550", "blocked",         false, ""}},
		{{"1139", 1, "5.0.918", "550", "rejected",        false, ""}},
		{{"1140", 1, "5.0.945", "",    "toomanyconn",     false, ""}},
		{{"1141", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1142", 1, "5.0.981", "",    "virusdetected",   false, ""}},
		{{"1143", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1145", 1, "5.0.934", "500", "mesgtoobig",      false, ""}},
		{{"1146", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1147", 1, "5.0.901", "551", "blocked",         false, ""}},
		{{"1148", 1, "5.0.980", "550", "spamdetected",    false, ""}},
		{{"1149", 1, "5.0.901", "550", "rejected",        false, ""}},
		{{"1150", 1, "5.7.1",   "553", "blocked",         false, ""}},
		{{"1151", 1, "5.0.0",   "550", "suspend",         false, ""}},
		{{"1152", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1153", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1154", 1, "5.7.1",   "550", "blocked",         false, ""}},
		{{"1155", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1156", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1157", 1, "5.0.0",   "",    "spamdetected",    false, ""}},
		{{"1158", 1, "5.0.0",   "",    "filtered",        false, ""}},
		{{"1159", 1, "5.0.0",   "",    "spamdetected",    false, ""}},
		{{"1161", 1, "5.3.4",   "552", "mesgtoobig",      false, ""},
		 {"1161", 2, "5.3.4",   "552", "mesgtoobig",      false, ""},
		 {"1161", 3, "5.3.4",   "552", "mesgtoobig",      false, ""},
		 {"1161", 4, "5.3.4",   "552", "mesgtoobig",      false, ""}},
		{{"1162", 1, "5.7.1",   "550", "requireptr",      false, ""}},
		{{"1163", 1, "5.1.1",   "550", "mailboxfull",     false, ""}},
		{{"1164", 1, "5.7.1",   "553", "authfailure",     false, ""}},
		{{"1165", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1168", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1169", 1, "5.4.3",   "",    "systemerror",     false, ""}},
		{{"1170", 1, "5.0.0",   "",    "systemerror",     false, ""},
		 {"1170", 2, "5.0.0",   "",    "systemerror",     false, ""}},
		{{"1171", 1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"1172", 1, "5.0.0",   "",    "hostunknown",      true, ""},
		 {"1172", 2, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1173", 1, "5.0.0",   "",    "networkerror",    false, ""}},
		{{"1175", 1, "5.0.0",   "",    "expired",         false, ""},
		 {"1175", 2, "5.0.0",   "",    "expired",         false, ""},
		 {"1175", 3, "5.0.0",   "",    "expired",         false, ""}},
		{{"1176", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1177", 1, "5.0.0",   "",    "filtered",        false, ""},
		 {"1177", 2, "5.0.0",   "",    "filtered",        false, ""}},
		{{"1178", 1, "4.0.947", "",    "expired",         false, ""}},
		{{"1179", 1, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1179", 2, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1181", 1, "5.0.0",   "",    "mailererror",     false, ""},
		 {"1181", 2, "5.0.939", "",    "mailererror",     false, ""},
		 {"1181", 3, "5.0.0",   "",    "mailererror",     false, ""}},
		{{"1182", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1183", 1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"1184", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1185", 1, "5.0.0",   "554", "suspend",         false, ""}},
		{{"1186", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1187", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1188", 1, "5.2.0",   "550", "spamdetected",    false, ""}},
		{{"1189", 1, "5.0.0",   "",    "expired",         false, ""}},
		{{"1190", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1191", 1, "5.0.0",   "550", "suspend",         false, ""}},
	}; EngineTest(t, "Exim", secretlist, false)
}

