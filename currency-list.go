package trader

import "strings"

// Verify returns whether a currency is valid according to the ISO 4217
func Verify(cur string) bool {
	cur = strings.ToUpper(cur)
	if ValidCurrencies[cur] != "" {
		return true
	}
	return false
}

// FullName returns the name of the currency in plain letters. In the case
// that the currency doesn't exist, an empty string will be returned.
func FullName(cur string) string {
	return ValidCurrencies[strings.ToUpper(cur)]
}

var (
	// ValidCurrencies according to the ISO 4217, though it is recommended
	// to pass through Verify() or FullName() to access this map.
	ValidCurrencies = map[string]string{
		"AFA": "Afghan Afghani",
		"AWG": "Aruban Florin",
		"AUD": "Australian Dollars",
		"ARS": "Argentine Pes",
		"AZN": "Azerbaijanian Manat",
		"BSD": "Bahamian Dollar",
		"BDT": "Bangladeshi Taka",
		"BBD": "Barbados Dollar",
		"BYR": "Belarussian Rouble",
		"BOB": "Bolivian Boliviano",
		"BRL": "Brazilian Real",
		"GBP": "British Pounds Sterling",
		"BGN": "Bulgarian Lev",
		"KHR": "Cambodia Riel",
		"CAD": "Canadian Dollars",
		"KYD": "Cayman Islands Dollar",
		"CLP": "Chilean Peso",
		"CNY": "Chinese Renminbi Yuan",
		"COP": "Colombian Peso",
		"CRC": "Costa Rican Colon",
		"HRK": "Croatia Kuna",
		"CPY": "Cypriot Pounds",
		"CZK": "Czech Koruna",
		"DKK": "Danish Krone",
		"DOP": "Dominican Republic Peso",
		"XCD": "East Caribbean Dollar",
		"EGP": "Egyptian Pound",
		"ERN": "Eritrean Nakfa",
		"EEK": "Estonia Kroon",
		"EUR": "Euro",
		"GEL": "Georgian Lari",
		"GHC": "Ghana Cedi",
		"GIP": "Gibraltar Pound",
		"GTQ": "Guatemala Quetzal",
		"HNL": "Honduras Lempira",
		"HKD": "Hong Kong Dollars",
		"HUF": "Hungary Forint",
		"ISK": "Icelandic Krona",
		"INR": "Indian Rupee",
		"IDR": "Indonesia Rupiah",
		"ILS": "Israel Shekel",
		"JMD": "Jamaican Dollar",
		"JPY": "Japanese yen",
		"KZT": "Kazakhstan Tenge",
		"KES": "Kenyan Shilling",
		"KWD": "Kuwaiti Dinar",
		"LVL": "Latvia Lat",
		"LBP": "Lebanese Pound",
		"LTL": "Lithuania Litas",
		"MOP": "Macau Pataca",
		"MKD": "Macedonian Denar",
		"MGA": "Malagascy Ariary",
		"MYR": "Malaysian Ringgit",
		"MTL": "Maltese Lira",
		"BAM": "Marka",
		"MUR": "Mauritius Rupee",
		"MXN": "Mexican Pesos",
		"MZM": "Mozambique Metical",
		"NPR": "Nepalese Rupee",
		"ANG": "Netherlands Antilles Guilder",
		"TWD": "New Taiwanese Dollars",
		"NZD": "New Zealand Dollars",
		"NIO": "Nicaragua Cordoba",
		"NGN": "Nigeria Naira",
		"KPW": "North Korean Won",
		"NOK": "Norwegian Krone",
		"OMR": "Omani Riyal",
		"PKR": "Pakistani Rupee",
		"PYG": "Paraguay Guarani",
		"PEN": "Peru New Sol",
		"PHP": "Philippine Pesos",
		"QAR": "Qatari Riyal",
		"RON": "Romanian New Leu",
		"RUB": "Russian Federation Ruble",
		"SAR": "Saudi Riyal",
		"CSD": "Serbian Dinar",
		"SCR": "Seychelles Rupee",
		"SGD": "Singapore Dollars",
		"SKK": "Slovak Koruna",
		"SIT": "Slovenia Tolar",
		"ZAR": "South African Rand",
		"KRW": "South Korean Won",
		"LKR": "Sri Lankan Rupee",
		"SRD": "Surinam Dollar",
		"SEK": "Swedish Krona",
		"CHF": "Swiss Francs",
		"TZS": "Tanzanian Shilling",
		"THB": "Thai Baht",
		"TTD": "Trinidad and Tobago Dollar",
		"TRY": "Turkish New Lira",
		"AED": "UAE Dirham",
		"USD": "US Dollars",
		"UGX": "Ugandian Shilling",
		"UAH": "Ukraine Hryvna",
		"UYU": "Uruguayan Peso",
		"UZS": "Uzbekistani Som",
		"VEB": "Venezuela Bolivar",
		"VND": "Vietnam Dong",
		"AMK": "Zambian Kwacha",
		"ZWD": "Zimbabwe Dollar",
	}
)
