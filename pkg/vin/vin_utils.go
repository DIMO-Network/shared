package vin

import (
	"regexp"
	"strings"
)

// VINs are not allowed to contain I, O, or Q. Except if they're Japanese chassis numbers
// In the year digit we additionally disallow U, Z, and 0.
// https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Model_year
var years = map[byte]int{
	'5': 2005,
	'6': 2006,
	'7': 2007,
	'8': 2008,
	'9': 2009,
	'A': 2010,
	'B': 2011,
	'C': 2012,
	'D': 2013,
	'E': 2014,
	'F': 2015,
	'G': 2016,
	'H': 2017,
	'J': 2018,
	'K': 2019,
	'L': 2020,
	'M': 2021,
	'N': 2022,
	'P': 2023,
	'R': 2024,
	'S': 2025,
	'T': 2026,
	'V': 2027,
	'W': 2028,
	'X': 2029,
	'Y': 2030,
	'1': 2031,
	'2': 2032,
	'3': 2033,
	'4': 2034,
}

type VIN string

func (v VIN) IsValidVIN() bool {
	if len(v) != 17 {
		return false
	}
	if AreAllCharactersTheSameAlt(string(v)) {
		return false
	}

	// Check if the string matches the VIN pattern
	// Excluding I, O, Q as per standard
	re := regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)
	return re.MatchString(string(v))
}

func (v VIN) IsValidJapanChassis() bool {
	if AreAllCharactersTheSameAlt(string(v)) {
		return false
	}
	re := regexp.MustCompile(`(?)^[A-Z0-9]{2,7}-?\d{5,7}$`)
	return re.MatchString(string(v))
}

// Year will decode the year portion from the VIN from 2005 to 2034. returns 0 if out of range. VIN nomenclature only allows for 30 year timespans, then repeats, ie. 2023 = 1993
func (v VIN) Year() int {
	if v.IsValidJapanChassis() {
		return 0
	}
	return years[v[9]]
}

func (v VIN) String() string {
	return string(v)
}

var teslaModelByDigit = map[byte]string{
	'C': "Cybertruck",
	'R': "Roadster",
	'S': "Model S",
	'X': "Model X",
	'Y': "Model Y",
	'3': "Model 3",
}

func (v VIN) TeslaModel() string {
	if v.IsValidJapanChassis() {
		return ""
	}
	// Maybe some error handling?
	return teslaModelByDigit[v[3]]
}

func (v VIN) Wmi() string {
	return string(v[0:3])
}

func (v VIN) VDS() string {
	if v.IsValidJapanChassis() {
		return ""
	}
	return string(v[3:8])
}

func (v VIN) VIS() string {
	if v.IsValidJapanChassis() {
		return ""
	}
	return string(v[9:17])
}

func (v VIN) CheckDigit() string {
	if v.IsValidJapanChassis() {
		return ""
	}
	return string(v[8:9])
}

func (v VIN) SerialNumber() string {
	if v.IsValidJapanChassis() {
		parts := strings.Split(string(v), "-")
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return string(v[11:17])
}

// AreAllCharactersTheSameAlt check not getting bogus VIN
func AreAllCharactersTheSameAlt(s string) bool {
	if len(s) <= 1 {
		return true
	}

	firstChar := rune(s[0])
	for _, char := range s[1:] {
		if char != firstChar {
			return false
		}
	}
	return true
}
