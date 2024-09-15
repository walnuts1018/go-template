package config

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		envs    map[string]string //env
		want    Config
		wantErr bool
	}{
		{
			name: "custom type",
			envs: map[string]string{
				"LOG_LEVEL": "debug",
			},
			want: Config{
				ServerPort:   "8080",
				ServerURL:    "localhost",
				LogLevel:     slog.LevelDebug,
				PSQLDSN:      "",
				PSQLHost:     "localhost",
				PSQLPort:     "5432",
				PSQLDatabase: "tobechanged",
				PSQLUser:     "postgres",
				PSQLPassword: "postgres",
				PSQLSSLMode:  "disable",
				PSQLTimeZone: "Asia/Tokyo",
			},
			wantErr: false,
		},
		// {
		// 	name:    "required",
		// 	envs:    map[string]string{},
		// 	wantErr: true,
		// },
		{
			name: "check default",
			envs: map[string]string{},
			want: Config{
				ServerPort:   "8080",
				ServerURL:    "localhost",
				LogLevel:     slog.LevelInfo,
				PSQLDSN:      "",
				PSQLHost:     "localhost",
				PSQLPort:     "5432",
				PSQLDatabase: "tobechanged",
				PSQLUser:     "postgres",
				PSQLPassword: "postgres",
				PSQLSSLMode:  "disable",
				PSQLTimeZone: "Asia/Tokyo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envs {
				if err := os.Setenv(k, v); err != nil {
					t.Errorf("failed to set env: %v", err)
					return
				}
				defer os.Unsetenv(k)
			}

			got, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
