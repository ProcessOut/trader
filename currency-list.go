package trader

// CurrencyInformation contains all the information relevant to a currency
type CurrencyInformation struct {
	// Number is the iso 4217 number of the currency
	Number uint
	// Places is the number of places after decimal separator
	Places int
	// FullName is the full name of the currency
	FullName string
	// Countries is a list of country names using said currency
	Countries []string
}

// Verify returns whether a currency is valid according to the ISO 4217
func (c CurrencyCode) Verify() bool {
	_, ok := validCurrencies[c.format()]
	return ok
}

// Information returns the information about said currency (see CurrencyInformation)
// If said currency doesn't exist, nil is retured. (Might want to call Verify() first)
func (c CurrencyCode) Information() *CurrencyInformation {
	new, ok := validCurrencies[c.format()]
	if !ok {
		return nil
	}
	return &new
}

// ValidCurrencies according to the ISO 4217, though it is recommended
// to pass through Verify() or Information() to access this map.
// This was parsed from https://en.wikipedia.org/wiki/ISO_4217#cite_note-ReferenceA-6
func ValidCurrencies() map[CurrencyCode]CurrencyInformation {
	return validCurrencies
}

// Unofficial currencies have an iso 4217 number of 0 (BTC, XBT, ETH)
var (
	validCurrencies = map[CurrencyCode]CurrencyInformation{
		"AED": CurrencyInformation{784, 2, "United Arab Emirates dirham", []string{"United Arab Emirates"}},
		"AFN": CurrencyInformation{971, 2, "Afghan afghani", []string{"Afghanistan"}},
		"ALL": CurrencyInformation{8, 2, "Albanian lek", []string{"Albania"}},
		"AMD": CurrencyInformation{51, 2, "Armenian dram", []string{"Armenia"}},
		"ANG": CurrencyInformation{532, 2, "Netherlands Antillean guilder", []string{"Curaçao (CW)", "Sint Maarten (SX)"}},
		"AOA": CurrencyInformation{973, 2, "Angolan kwanza", []string{"Angola"}},
		"ARS": CurrencyInformation{32, 2, "Argentine peso", []string{"Argentina"}},
		"AUD": CurrencyInformation{36, 2, "Australian dollar", []string{"Australia", "Christmas Island (CX)", "Cocos (Keeling) Islands (CC)", "Heard Island and McDonald Islands (HM)", "Kiribati (KI)", "Nauru (NR)", "Norfolk Island (NF)", "Tuvalu (TV)", "Australian Antarctic Territory"}},
		"AWG": CurrencyInformation{533, 2, "Aruban florin", []string{"Aruba"}},
		"AZN": CurrencyInformation{944, 2, "Azerbaijani manat", []string{"Azerbaijan"}},
		"BAM": CurrencyInformation{977, 2, "Bosnia and Herzegovina convertible mark", []string{"Bosnia and Herzegovina"}},
		"BBD": CurrencyInformation{52, 2, "Barbados dollar", []string{"Barbados"}},
		"BDT": CurrencyInformation{50, 2, "Bangladeshi taka", []string{"Bangladesh"}},
		"BGN": CurrencyInformation{975, 2, "Bulgarian lev", []string{"Bulgaria"}},
		"BHD": CurrencyInformation{48, 3, "Bahraini dinar", []string{"Bahrain"}},
		"BIF": CurrencyInformation{108, 0, "Burundian franc", []string{"Burundi"}},
		"BMD": CurrencyInformation{60, 2, "Bermudian dollar", []string{"Bermuda"}},
		"BND": CurrencyInformation{96, 2, "Brunei dollar", []string{"Brunei", "auxiliary in Singapore (SG)"}},
		"BOB": CurrencyInformation{68, 2, "Boliviano", []string{"Bolivia"}},
		"BOV": CurrencyInformation{984, 2, "Bolivian Mvdol (funds code)", []string{"Bolivia"}},
		"BRL": CurrencyInformation{986, 2, "Brazilian real", []string{"Brazil"}},
		"BSD": CurrencyInformation{44, 2, "Bahamian dollar", []string{"Bahamas"}},
		"BTC": CurrencyInformation{0, 8, "Bitcoin", []string{}},
		"XBT": CurrencyInformation{0, 8, "Bitcoin", []string{}},
		"BTN": CurrencyInformation{64, 2, "Bhutanese ngultrum", []string{"Bhutan"}},
		"BWP": CurrencyInformation{72, 2, "Botswana pula", []string{"Botswana"}},
		"BYN": CurrencyInformation{933, 2, "Belarusian ruble", []string{"Belarus"}},
		"BYR": CurrencyInformation{974, 0, "Belarusian ruble", []string{"Belarus"}},
		"BZD": CurrencyInformation{84, 2, "Belize dollar", []string{"Belize"}},
		"CAD": CurrencyInformation{124, 2, "Canadian dollar", []string{"Canada"}},
		"CDF": CurrencyInformation{976, 2, "Congolese franc", []string{"Democratic Republic of the Congo"}},
		"CHE": CurrencyInformation{947, 2, "WIR Euro (complementary currency)", []string{"Switzerland"}},
		"CHF": CurrencyInformation{756, 2, "Swiss franc", []string{"Switzerland", "Liechtenstein (LI)"}},
		"CHW": CurrencyInformation{948, 2, "WIR Franc (complementary currency)", []string{"Switzerland"}},
		"CLF": CurrencyInformation{990, 4, "Unidad de Fomento (funds code)", []string{"Chile"}},
		"CLP": CurrencyInformation{152, 0, "Chilean peso", []string{"Chile"}},
		"CNY": CurrencyInformation{156, 2, "Chinese yuan", []string{"China"}},
		"COP": CurrencyInformation{170, 2, "Colombian peso", []string{"Colombia"}},
		"COU": CurrencyInformation{970, 2, "Unidad de Valor Real (UVR) (funds code)", []string{"Colombia"}},
		"CRC": CurrencyInformation{188, 2, "Costa Rican colon", []string{"Costa Rica"}},
		"CUC": CurrencyInformation{931, 2, "Cuban convertible peso", []string{"Cuba"}},
		"CUP": CurrencyInformation{192, 2, "Cuban peso", []string{"Cuba"}},
		"CVE": CurrencyInformation{132, 0, "Cape Verde escudo", []string{"Cape Verde"}},
		"CZK": CurrencyInformation{203, 2, "Czech koruna", []string{"Czech Republic"}},
		"DJF": CurrencyInformation{262, 0, "Djiboutian franc", []string{"Djibouti"}},
		"DKK": CurrencyInformation{208, 2, "Danish krone", []string{"Denmark", "Faroe Islands (FO)", "Greenland (GL)"}},
		"DOP": CurrencyInformation{214, 2, "Dominican peso", []string{"Dominican Republic"}},
		"DZD": CurrencyInformation{12, 2, "Algerian dinar", []string{"Algeria"}},
		"EGP": CurrencyInformation{818, 2, "Egyptian pound", []string{"Egypt", "auxiliary in Gaza Strip"}},
		"ERN": CurrencyInformation{232, 2, "Eritrean nakfa", []string{"Eritrea"}},
		"ETB": CurrencyInformation{230, 2, "Ethiopian birr", []string{"Ethiopia"}},
		"ETH": CurrencyInformation{0, 2, "Ether", []string{}},
		"EUR": CurrencyInformation{978, 2, "Euro", []string{"Akrotiri and Dhekelia", "Andorra (AD)", "Austria (AT)", "Belgium (BE)", "Cyprus (CY)", "Estonia (EE)", "Finland (FI)", "France (FR)", "Germany (DE)", "Greece (GR)", "Guadeloupe (GP)", "Ireland (IE)", "Italy (IT)", "Kosovo", "Latvia (LV)", "Lithuania (LT)", "Luxembourg (LU)", "Malta (MT)", "Martinique (MQ)", "Mayotte (YT)", "Monaco (MC)", "Montenegro (ME)", "Netherlands (NL)", "Portugal (PT)", "Réunion (RE)", "Saint Barthélemy (BL)", "Saint Pierre and Miquelon (PM)", "San Marino (SM)", "Slovakia (SK)", "Slovenia (SI)", "Spain (ES)", "Vatican City (VA); see Eurozone"}},
		"FJD": CurrencyInformation{242, 2, "Fiji dollar", []string{"Fiji"}},
		"FKP": CurrencyInformation{238, 2, "Falkland Islands pound", []string{"Falkland Islands (pegged to GBP 1:1)"}},
		"GBP": CurrencyInformation{826, 2, "Pound sterling", []string{"United Kingdom", "the Isle of Man (IM", "see Manx pound)", "Jersey (JE", "see Jersey pound)", "Guernsey (GG", "see Guernsey pound)", "South Georgia and the South Sandwich Islands (GS)", "British Indian Ocean Territory (IO) (also uses USD)", "Tristan da Cunha (SH-TA)", "and British Antarctic Territory"}},
		"GEL": CurrencyInformation{981, 2, "Georgian lari", []string{"Georgia (except Abkhazia (GE-AB) and South Ossetia)"}},
		"GHS": CurrencyInformation{936, 2, "Ghanaian cedi", []string{"Ghana"}},
		"GIP": CurrencyInformation{292, 2, "Gibraltar pound", []string{"Gibraltar (pegged to GBP 1:1)"}},
		"GMD": CurrencyInformation{270, 2, "Gambian dalasi", []string{"Gambia"}},
		"GNF": CurrencyInformation{324, 0, "Guinean franc", []string{"Guinea"}},
		"GTQ": CurrencyInformation{320, 2, "Guatemalan quetzal", []string{"Guatemala"}},
		"GYD": CurrencyInformation{328, 2, "Guyanese dollar", []string{"Guyana"}},
		"HKD": CurrencyInformation{344, 2, "Hong Kong dollar", []string{"Hong Kong", "Macao (MO)"}},
		"HNL": CurrencyInformation{340, 2, "Honduran lempira", []string{"Honduras"}},
		"HRK": CurrencyInformation{191, 2, "Croatian kuna", []string{"Croatia"}},
		"HTG": CurrencyInformation{332, 2, "Haitian gourde", []string{"Haiti"}},
		"HUF": CurrencyInformation{348, 2, "Hungarian forint", []string{"Hungary"}},
		"IDR": CurrencyInformation{360, 2, "Indonesian rupiah", []string{"Indonesia"}},
		"ILS": CurrencyInformation{376, 2, "Israeli new shekel", []string{"Israel", "State of Palestine (PS)"}},
		"INR": CurrencyInformation{356, 2, "Indian rupee", []string{"India", "Bhutan", "Nepal", "Zimbabwe"}},
		"IQD": CurrencyInformation{368, 3, "Iraqi dinar", []string{"Iraq"}},
		"IRR": CurrencyInformation{364, 2, "Iranian rial", []string{"Iran"}},
		"ISK": CurrencyInformation{352, 0, "Icelandic króna", []string{"Iceland"}},
		"JMD": CurrencyInformation{388, 2, "Jamaican dollar", []string{"Jamaica"}},
		"JOD": CurrencyInformation{400, 3, "Jordanian dinar", []string{"Jordan", "auxiliary in West Bank"}},
		"JPY": CurrencyInformation{392, 0, "Japanese yen", []string{"Japan"}},
		"KES": CurrencyInformation{404, 2, "Kenyan shilling", []string{"Kenya"}},
		"KGS": CurrencyInformation{417, 2, "Kyrgyzstani som", []string{"Kyrgyzstan"}},
		"KHR": CurrencyInformation{116, 2, "Cambodian riel", []string{"Cambodia"}},
		"KMF": CurrencyInformation{174, 0, "Comoro franc", []string{"Comoros"}},
		"KPW": CurrencyInformation{408, 2, "North Korean won", []string{"North Korea"}},
		"KRW": CurrencyInformation{410, 0, "South Korean won", []string{"South Korea"}},
		"KWD": CurrencyInformation{414, 3, "Kuwaiti dinar", []string{"Kuwait"}},
		"KYD": CurrencyInformation{136, 2, "Cayman Islands dollar", []string{"Cayman Islands"}},
		"KZT": CurrencyInformation{398, 2, "Kazakhstani tenge", []string{"Kazakhstan"}},
		"LAK": CurrencyInformation{418, 2, "Lao kip", []string{"Laos"}},
		"LBP": CurrencyInformation{422, 2, "Lebanese pound", []string{"Lebanon"}},
		"LKR": CurrencyInformation{144, 2, "Sri Lankan rupee", []string{"Sri Lanka"}},
		"LRD": CurrencyInformation{430, 2, "Liberian dollar", []string{"Liberia"}},
		"LSL": CurrencyInformation{426, 2, "Lesotho loti", []string{"Lesotho"}},
		"LYD": CurrencyInformation{434, 3, "Libyan dinar", []string{"Libya"}},
		"MAD": CurrencyInformation{504, 2, "Moroccan dirham", []string{"Morocco"}},
		"MDL": CurrencyInformation{498, 2, "Moldovan leu", []string{"Moldova (except Transnistria)"}},
		"MGA": CurrencyInformation{969, 1, "Malagasy ariary", []string{"Madagascar"}},
		"MKD": CurrencyInformation{807, 2, "Macedonian denar", []string{"Macedonia"}},
		"MMK": CurrencyInformation{104, 2, "Myanmar kyat", []string{"Myanmar"}},
		"MNT": CurrencyInformation{496, 2, "Mongolian tögrög", []string{"Mongolia"}},
		"MOP": CurrencyInformation{446, 2, "Macanese pataca", []string{"Macao"}},
		"MRO": CurrencyInformation{478, 1, "Mauritanian ouguiya", []string{"Mauritania"}},
		"MUR": CurrencyInformation{480, 2, "Mauritian rupee", []string{"Mauritius"}},
		"MVR": CurrencyInformation{462, 2, "Maldivian rufiyaa", []string{"Maldives"}},
		"MWK": CurrencyInformation{454, 2, "Malawian kwacha", []string{"Malawi"}},
		"MXN": CurrencyInformation{484, 2, "Mexican peso", []string{"Mexico"}},
		"MXV": CurrencyInformation{979, 2, "Mexican Unidad de Inversion (UDI) (funds code)", []string{"Mexico"}},
		"MYR": CurrencyInformation{458, 2, "Malaysian ringgit", []string{"Malaysia"}},
		"MZN": CurrencyInformation{943, 2, "Mozambican metical", []string{"Mozambique"}},
		"NAD": CurrencyInformation{516, 2, "Namibian dollar", []string{"Namibia"}},
		"NGN": CurrencyInformation{566, 2, "Nigerian naira", []string{"Nigeria"}},
		"NIO": CurrencyInformation{558, 2, "Nicaraguan córdoba", []string{"Nicaragua"}},
		"NOK": CurrencyInformation{578, 2, "Norwegian krone", []string{"Norway", "Svalbard and Jan Mayen (SJ)", "Bouvet Island (BV)", "Queen Maud Land", "Peter I Island"}},
		"NPR": CurrencyInformation{524, 2, "Nepalese rupee", []string{"Nepal"}},
		"NZD": CurrencyInformation{554, 2, "New Zealand dollar", []string{"New Zealand", "Cook Islands (CK)", "Niue (NU)", "Pitcairn Islands (PN; see also Pitcairn Islands dollar)", "Tokelau (TK)", "Ross Dependency"}},
		"OMR": CurrencyInformation{512, 3, "Omani rial", []string{"Oman"}},
		"PAB": CurrencyInformation{590, 2, "Panamanian balboa", []string{"Panama"}},
		"PEN": CurrencyInformation{604, 2, "Peruvian Sol", []string{"Peru"}},
		"PGK": CurrencyInformation{598, 2, "Papua New Guinean kina", []string{"Papua New Guinea"}},
		"PHP": CurrencyInformation{608, 2, "Philippine peso", []string{"Philippines"}},
		"PKR": CurrencyInformation{586, 2, "Pakistani rupee", []string{"Pakistan"}},
		"PLN": CurrencyInformation{985, 2, "Polish złoty", []string{"Poland"}},
		"PYG": CurrencyInformation{600, 0, "Paraguayan guaraní", []string{"Paraguay"}},
		"QAR": CurrencyInformation{634, 2, "Qatari riyal", []string{"Qatar"}},
		"RON": CurrencyInformation{946, 2, "Romanian leu", []string{"Romania"}},
		"RSD": CurrencyInformation{941, 2, "Serbian dinar", []string{"Serbia"}},
		"RUB": CurrencyInformation{643, 2, "Russian ruble", []string{"Russia", "Abkhazia (GE-AB)", "South Ossetia", "Crimea"}},
		"RWF": CurrencyInformation{646, 0, "Rwandan franc", []string{"Rwanda"}},
		"SAR": CurrencyInformation{682, 2, "Saudi riyal", []string{"Saudi Arabia"}},
		"SBD": CurrencyInformation{90, 2, "Solomon Islands dollar", []string{"Solomon Islands"}},
		"SCR": CurrencyInformation{690, 2, "Seychelles rupee", []string{"Seychelles"}},
		"SDG": CurrencyInformation{938, 2, "Sudanese pound", []string{"Sudan"}},
		"SEK": CurrencyInformation{752, 2, "Swedish krona/kronor", []string{"Sweden"}},
		"SGD": CurrencyInformation{702, 2, "Singapore dollar", []string{"Singapore", "auxiliary in Brunei (BN)"}},
		"SHP": CurrencyInformation{654, 2, "Saint Helena pound", []string{"Saint Helena (SH-SH)", "Ascension Island (SH-AC) (pegged to GBP 1:1)"}},
		"SLL": CurrencyInformation{694, 2, "Sierra Leonean leone", []string{"Sierra Leone"}},
		"SOS": CurrencyInformation{706, 2, "Somali shilling", []string{"Somalia (except Somaliland)"}},
		"SRD": CurrencyInformation{968, 2, "Surinamese dollar", []string{"Suriname"}},
		"SSP": CurrencyInformation{728, 2, "South Sudanese pound", []string{"South Sudan"}},
		"STD": CurrencyInformation{678, 2, "São Tomé and Príncipe dobra", []string{"São Tomé and Príncipe"}},
		"SVC": CurrencyInformation{222, 2, "Salvadoran colón", []string{"El Salvador"}},
		"SYP": CurrencyInformation{760, 2, "Syrian pound", []string{"Syria"}},
		"SZL": CurrencyInformation{748, 2, "Swazi lilangeni", []string{"Swaziland"}},
		"THB": CurrencyInformation{764, 2, "Thai baht", []string{"Thailand", "Cambodia", "Myanmar", "Laos"}},
		"TJS": CurrencyInformation{972, 2, "Tajikistani somoni", []string{"Tajikistan"}},
		"TMT": CurrencyInformation{934, 2, "Turkmenistani manat", []string{"Turkmenistan"}},
		"TND": CurrencyInformation{788, 3, "Tunisian dinar", []string{"Tunisia"}},
		"TOP": CurrencyInformation{776, 2, "Tongan paʻanga", []string{"Tonga"}},
		"TRY": CurrencyInformation{949, 2, "Turkish lira", []string{"Turkey", "Northern Cyprus"}},
		"TTD": CurrencyInformation{780, 2, "Trinidad and Tobago dollar", []string{"Trinidad and Tobago"}},
		"TWD": CurrencyInformation{901, 2, "New Taiwan dollar", []string{"Taiwan"}},
		"TZS": CurrencyInformation{834, 2, "Tanzanian shilling", []string{"Tanzania"}},
		"UAH": CurrencyInformation{980, 2, "Ukrainian hryvnia", []string{"Ukraine"}},
		"UGX": CurrencyInformation{800, 0, "Ugandan shilling", []string{"Uganda"}},
		"USD": CurrencyInformation{840, 2, "United States dollar", []string{"United States", "American Samoa (AS)", "Barbados (BB) (as well as Barbados Dollar)", "Bermuda (BM) (as well as Bermudian Dollar)", "British Indian Ocean Territory (IO) (also uses GBP)", "British Virgin Islands (VG)", "Caribbean Netherlands (BQ - Bonaire", "Sint Eustatius and Saba)", "Ecuador (EC)", "El Salvador (SV)", "Guam (GU)", "Haiti (HT)", "Marshall Islands (MH)", "Federated States of Micronesia (FM)", "Northern Mariana Islands (MP)", "Palau (PW)", "Panama (PA)", "Puerto Rico (PR)", "Timor-Leste (TL)", "Turks and Caicos Islands (TC)", "U.S. Virgin Islands (VI)", "Zimbabwe (ZW)"}},
		"USN": CurrencyInformation{997, 2, "United States dollar (next day) (funds code)", []string{"United States"}},
		"UYI": CurrencyInformation{940, 0, "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)", []string{"Uruguay"}},
		"UYU": CurrencyInformation{858, 2, "Uruguayan peso", []string{"Uruguay"}},
		"UZS": CurrencyInformation{860, 2, "Uzbekistan som", []string{"Uzbekistan"}},
		"VEF": CurrencyInformation{937, 2, "Venezuelan bolívar", []string{"Venezuela"}},
		"VND": CurrencyInformation{704, 0, "Vietnamese dong", []string{"Vietnam"}},
		"VUV": CurrencyInformation{548, 0, "Vanuatu vatu", []string{"Vanuatu"}},
		"WST": CurrencyInformation{882, 2, "Samoan tala", []string{"Samoa"}},
		"XAF": CurrencyInformation{950, 0, "CFA franc BEAC", []string{"Cameroon (CM)", "Central African Republic (CF)", "Republic of the Congo (CG)", "Chad (TD)", "Equatorial Guinea (GQ)", "Gabon (GA)"}},
		"XAG": CurrencyInformation{961, -1, "Silver (one troy ounce)", []string{}},
		"XAU": CurrencyInformation{959, -1, "Gold (one troy ounce)", []string{}},
		"XBA": CurrencyInformation{955, -1, "European Composite Unit (EURCO) (bond market unit)", []string{}},
		"XBB": CurrencyInformation{956, -1, "European Monetary Unit (E.M.U.-6) (bond market unit)", []string{}},
		"XBC": CurrencyInformation{957, -1, "European Unit of Account 9 (E.U.A.-9) (bond market unit)", []string{}},
		"XBD": CurrencyInformation{958, -1, "European Unit of Account 17 (E.U.A.-17) (bond market unit)", []string{}},
		"XCD": CurrencyInformation{951, 2, "East Caribbean dollar", []string{"Anguilla (AI)", "Antigua and Barbuda (AG)", "Dominica (DM)", "Grenada (GD)", "Montserrat (MS)", "Saint Kitts and Nevis (KN)", "Saint Lucia (LC)", "Saint Vincent and the Grenadines (VC)"}},
		"XDR": CurrencyInformation{960, -1, "Special drawing rights", []string{"International Monetary Fund"}},
		"XOF": CurrencyInformation{952, 0, "CFA franc BCEAO", []string{"Benin (BJ)", "Burkina Faso (BF)", "Côte d'Ivoire (CI)", "Guinea-Bissau (GW)", "Mali (ML)", "Niger (NE)", "Senegal (SN)", "Togo (TG)"}},
		"XPD": CurrencyInformation{964, -1, "Palladium (one troy ounce)", []string{}},
		"XPF": CurrencyInformation{953, 0, "CFP franc (franc Pacifique)", []string{"French territories of the Pacific Ocean: French Polynesia (PF)", "New Caledonia (NC)", "Wallis and Futuna (WF)"}},
		"XPT": CurrencyInformation{962, -1, "Platinum (one troy ounce)", []string{}},
		"XSU": CurrencyInformation{994, -1, "SUCRE", []string{"Unified System for Regional Compensation (SUCRE)"}},
		"XTS": CurrencyInformation{963, -1, "Code reserved for testing purposes", []string{}},
		"XUA": CurrencyInformation{965, -1, "ADB Unit of Account", []string{"African Development Bank"}},
		"XXX": CurrencyInformation{999, -1, "No currency", []string{}},
		"YER": CurrencyInformation{886, 2, "Yemeni rial", []string{"Yemen"}},
		"ZAR": CurrencyInformation{710, 2, "South African rand", []string{"South Africa"}},
		"ZMW": CurrencyInformation{967, 2, "Zambian kwacha", []string{"Zambia"}},
		"ZWL": CurrencyInformation{932, 2, "Zimbabwean dollar A/10", []string{"Zimbabwe"}},
	}
)
