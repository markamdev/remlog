package server

import "testing"

func TestInitServer(t *testing.T) {
	type args struct {
		conf Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"0 auth port", args{conf: Config{AuthPort: 0, LogPort: 9100}}, true},
		{"0 log port", args{conf: Config{AuthPort: 9100, LogPort: 0}}, true},
		{"equal ports", args{conf: Config{AuthPort: 9000, LogPort: 9000}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitServer(tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("InitServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
