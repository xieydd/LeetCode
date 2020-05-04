package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{nums: []int{1,2,3,4,5}, value: 6}, want: 5},
		{name: "test1", args: args{nums: []int{2,5,7,8,9,11}, value: 3}, want: 1},
		{name: "test2", args: args{nums: []int{1,2,3,45,90,236,345,1110}, value: 236}, want: 5},
		{name: "test3", args: args{nums: []int{1,2,3,45,90,101,120,345,1110}, value: 236}, want: 7},
		{name: "test4", args: args{nums: []int{1,6,7,45,90,101,120,345,1110,12344,90000}, value: 236}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
