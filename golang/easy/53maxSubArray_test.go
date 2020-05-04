package easy

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{name: "test", args: args{nums: []int{1,-2,3,4,5}}, want: 12, want1: []int{3,4,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxSubArray(tt.args.nums)
			if got != tt.want {
				t.Errorf("MaxSubArray() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MaxSubArray() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
