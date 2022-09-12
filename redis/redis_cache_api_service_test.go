package redis

import (
	"testing"
)

func Test_cacheService_keyBuilder(t *testing.T) {
	type fields struct {
		redisClient CacheService
		prefix      string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "with prefix",
			fields: fields{
				prefix: "devices_api",
			},
			args: args{key: "123"},
			want: "devices_api_123",
		},
		{
			name: "empty prefix",
			fields: fields{
				prefix: "",
			},
			args: args{key: "123"},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &cacheService{
				redisClient: tt.fields.redisClient,
				prefix:      tt.fields.prefix,
			}
			if got := cs.keyBuilder(tt.args.key); got != tt.want {
				t.Errorf("keyBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
