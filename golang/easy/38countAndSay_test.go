package easy

import "testing"

func TestCountAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{n: 4}, want: "1211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountAndSay(tt.args.n); got != tt.want {
				t.Errorf("CountAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
