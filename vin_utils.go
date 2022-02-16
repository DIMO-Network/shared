package shared

// VINs are not allowed to contain I, O, or Q.
// In the year digit we additionally disallow U, Z, and 0.
// https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Model_year
var years = map[byte]int{
	'R': 1994,
	'S': 1995,
	'T': 1996,
	'V': 1997,
	'W': 1998,
	'X': 1999,
	'Y': 2000,
	'1': 2001,
	'2': 2002,
	'3': 2003,
	'4': 2004,
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
}

type VIN string

// Year will decode the year portion from the VIN from 1994 to 2023. returns 0 if out of range. VIN nomenclature only allows for 30 year timespans, then repeats, ie. 2023 = 1993
func (v VIN) Year() int {
	return years[v[9]]
}

func (v VIN) String() string {
	return string(v)
}
