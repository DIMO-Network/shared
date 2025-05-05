package vin

import "testing"

func TestVIN_Year(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want int
	}{
		{
			name: "2006 VIN bmw",
			v:    "WBAVD33556KL51323",
			want: 2006,
		},
		{
			name: "2020 VIN",
			v:    "WBS2U7C05L7E99189",
			want: 2020,
		},
		{
			name: "2022 VIN",
			v:    "2FMPK4J97NBA08224",
			want: 2022,
		},
		{
			name: "invalid year vin",
			v:    "2FMPK4J97IBA08224",
			want: 0,
		},
		{
			name: "japan chasis number",
			v:    "ZWR90-8000186",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Year(); got != tt.want {
				t.Errorf("Year() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVIN_VDS(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want string
	}{
		{
			name: "1FT FX1E5 7 JKE37092 VIN",
			v:    "1FTFX1E57JKE37092",
			want: "FX1E5",
		},
		{
			name: "japan chasis number",
			v:    "ZWR90-8000186",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.VDS(); got != tt.want {
				t.Errorf("VDS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVIN_VIS(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want string
	}{
		{
			name: "1FT FX1E5 7 JKE37092 VIN",
			v:    "1FTFX1E57JKE37092",
			want: "JKE37092",
		},
		{
			name: "japan chasis number",
			v:    "ZWR90-8000186",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.VIS(); got != tt.want {
				t.Errorf("VIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVIN_CheckDigit(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want string
	}{
		{
			name: "1FT FX1E5 7 JKE37092 VIN",
			v:    "1FTFX1E57JKE37092",
			want: "7",
		},
		{
			name: "japan chasis number",
			v:    "ZWR90-8000186",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.CheckDigit(); got != tt.want {
				t.Errorf("CheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVIN_SerialNumber(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want string
	}{
		{
			name: "1HGBH41JXMN109186 VIN",
			v:    "1HGBH41JXMN109186",
			want: "109186",
		},
		{
			name: "japan chasis number",
			v:    "ZWR90-8000186",
			want: "8000186",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.SerialNumber(); got != tt.want {
				t.Errorf("SerialNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVIN_String(t *testing.T) {
	vs := "WBAEW53494PG11352"
	v := VIN(vs)
	if v.String() != vs {
		t.Errorf("did not get expected vin strin")
	}

	vs = "ZWR90-8000186"
	v = VIN(vs)
	if v.String() != vs {
		t.Errorf("did not get expected vin strin")
	}
}

func TestVIN_IsJapanChassis(t *testing.T) {
	tests := []struct {
		name string
		v    VIN
		want bool
	}{
		{
			name: "Japan chassis number with dash",
			v:    "ZWR90-8000186",
			want: true,
		},
		{
			name: "Invalid Japan chassis without dash",
			v:    "ZWR908000186",
			want: false,
		},
		{
			name: "Standard VIN with 17 characters",
			v:    "1FTFX1E57JKE37092",
			want: false,
		},
		{
			name: "Short VIN without dash",
			v:    "WBMWD",
			want: false,
		},
		{
			name: "Short VIN with dash",
			v:    "BMW-123",
			want: true,
		},
		{
			name: "Empty string",
			v:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsJapanChassis(); got != tt.want {
				t.Errorf("IsJapanChassis() = %v, want %v", got, tt.want)
			}
		})
	}
}
