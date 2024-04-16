package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugString(t *testing.T) {
	tests := []struct {
		term string
		want string
	}{
		{
			term: "Mercedes-Benz",
			want: "mercedes-benz",
		},
		{
			term: "AC",
			want: "ac",
		},
		{
			term: "AM General",
			want: "am-general",
		},
		{
			term: "Alfa Romeo",
			want: "alfa-romeo",
		},
		{
			term: "Citroën",
			want: "citroen",
		},
		{
			term: "SsangYong",
			want: "ssangyong",
		},
		{
			term: "Land Rover",
			want: "land-rover",
		},
	}
	for _, tt := range tests {
		t.Run(tt.term, func(t *testing.T) {
			assert.Equalf(t, tt.want, SlugString(tt.term), "SlugString(%v)", tt.term)
		})
	}
}
