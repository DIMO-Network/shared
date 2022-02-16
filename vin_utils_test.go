package shared

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
			name: "2004 VIN",
			v:    "WBAEW53494PG11352",
			want: 2004,
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
