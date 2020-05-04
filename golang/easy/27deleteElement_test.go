package easy

import "testing"

func TestDeleteElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{nums: []int{0,1,2,2,3,0,4,2}, val: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("DeleteElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
