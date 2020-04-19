package logitech

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	fmt.Println(logitech.Load())
	fmt.Println(pLogiSteeringInitialize.Find())

}

func TestLogiSteeringInitialize(t *testing.T) {
	type args struct {
		ignoreXInput bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "ignore Input",
			args:    args{true},
			want:    true,
			wantErr: false,
		}, {
			name:    "dont ignore Input",
			args:    args{false},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogiSteeringInitialize(tt.args.ignoreXInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogiSteeringInitialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LogiSteeringInitialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
