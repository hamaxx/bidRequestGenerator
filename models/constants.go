package models

var Genders = []ChoiceString{
	ChoiceString{100, ""},
	ChoiceString{50, "male"},
	ChoiceString{55, "female"},
	ChoiceString{1, "other"},
}
var AgeRanges = []ChoiceString{
	ChoiceString{100, ""},
	ChoiceString{10, "12-17"},
	ChoiceString{20, "18-24"},
	ChoiceString{30, "25-34"},
	ChoiceString{20, "35-44"},
	ChoiceString{15, "45-54"},
	ChoiceString{10, "55-64"},
	ChoiceString{5, "65+"},
}

var DeviceTypes = []string{"mobile", "desktop", "tablet"}
var DeviceOs = map[string][]string{
	"desktop": []string{"windows 10", "windows 8", "windows 7", "macos", "linux"},
	"mobile":  []string{"ios", "android"},
	"tablet":  []string{"ios", "android"},
}

var Countries = []string{
	"US",
	"GB",
	"FR",
	"JP",
	"DE",
	"AU",
	"ES",
	"IT",
	"CA",
	"BR",
	"IN",
	"TR",
	"RU",
	"NO",
	"CH",
	"IL",
	"NZ",
	"ID",
	"MX",
	"BE",
	"ZA",
	"NL",
	"PH",
	"KR",
	"SE",
	"TH",
	"IE",
	"AR",
	"MY",
	"PL",
	"PT",
	"RO",
	"IQ",
	"AT",
	"SG",
	"DZ",
	"MA",
	"GR",
	"FI",
	"AE",
	"SA",
	"PK",
	"HU",
	"CO",
	"EG",
	"DK",
	"UA",
	"BY",
	"BG",
	"CL",
	"VN",
	"TW",
	"NG",
	"TN",
	"SI",
	"HK",
	"RS",
	"RE",
	"CZ",
	"VE",
	"KE",
	"HR",
	"AL",
	"CI",
	"BD",
	"SY",
	"LU",
	"PE",
	"BA",
	"CM",
	"CN",
	"QA",
	"SN",
	"GH",
	"SK",
	"DO",
	"KW",
	"LB",
	"JO",
	"NP",
	"CY",
	"LT",
	"LK",
	"OM",
	"MK",
	"IR",
	"MT",
	"ZW",
	"MD",
	"EE",
	"PY",
	"KZ",
	"MU",
	"EC",
	"GE",
	"MQ",
	"BH",
	"PR",
	"MM",
	"LV",
	"IS",
	"TT",
	"UY",
	"PA",
	"MG",
	"GP",
	"SD",
	"YE",
	"BO",
	"JM",
	"TZ",
	"LA",
	"ME",
	"JE",
	"BJ",
	"GA",
	"ML",
	"AM",
}

var Regions = map[string][]string{
	"US": []string{
		"CA",
		"FL",
		"TX",
		"NY",
		"GA",
		"PA",
		"NC",
		"OH",
		"IL",
		"MI",
		"NJ",
		"VA",
		"MA",
		"MD",
		"SC",
		"TN",
		"IN",
		"WA",
		"MO",
		"AZ",
		"WI",
		"AL",
		"CT",
		"MN",
		"CO",
		"KY",
		"LA",
		"OR",
		"DC",
		"OK",
		"DE",
		"NV",
		"IA",
		"UT",
		"MS",
		"KS",
		"NH",
		"WV",
		"AR",
		"HI",
		"NE",
		"ME",
		"RI",
		"NM",
		"ID",
		"VT",
		"AK",
		"ND",
		"MT",
		"SD",
	},
	"GB": []string{
		"ENG",
		"SCT",
		"WLS",
		"NIR",
		"LND",
		"BIR",
		"HRT",
		"SRY",
		"KEN",
		"HAM",
		"ESS",
		"LAN",
		"MAN",
		"NGM",
		"LDS",
		"WSX",
		"NFK",
		"CAM",
		"STS",
		"GLG",
		"DBY",
		"TWH",
		"OXF",
		"WIL",
		"SOM",
		"LIV",
	},
	"FR": []string{
		"IDF",
		"ARA",
		"OCC",
		"NAQ",
		"GES",
		"HDF",
		"J",
		"PAC",
		"NOR",
		"PDL",
		"BRE",
		"BFC",
		"CVL",
		"U",
		"R",
		"E",
		"F",
		"B",
	},
}
