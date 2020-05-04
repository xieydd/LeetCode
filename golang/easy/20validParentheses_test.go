package easy

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Add test cases.
		{name: "test-true", args: args{s: "{([])}"}, want: true},
		{name: "test-false", args: args{s: "{[[])}"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
