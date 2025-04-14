package db

import "testing"

func Test_safePrintDSN(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
		want string
	}{
		{
			name: "password gets cleaned",
			dsn:  "user=bob password=esponja dbname=dimo host=localhost port=5432 sslmode=disable",
			want: "user=bob password=e*** dbname=dimo host=localhost port=5432 sslmode=disable",
		},
		{
			name: "empty password",
			dsn:  "user=bob password= dbname=dimo host=localhost port=5432 sslmode=disable",
			want: "user=bob password= dbname=dimo host=localhost port=5432 sslmode=disable",
		},
		{
			name: "no password field",
			dsn:  "user=bob dbname=dimo host=localhost port=5432 sslmode=disable",
			want: "user=bob dbname=dimo host=localhost port=5432 sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safePrintDSN(tt.dsn); got != tt.want {
				t.Errorf("safePrintDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
