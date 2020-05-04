package easy

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{s: "Hello Worlds  "}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
