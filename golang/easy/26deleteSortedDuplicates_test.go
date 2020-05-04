package easy

import "testing"

func TestDeleteSortedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test", args: args{nums: []int{1,2,3,3,4,4,5,5,5,6}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSortedArray(tt.args.nums); got != tt.want {
				t.Errorf("DeleteSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
