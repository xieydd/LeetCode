package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		arrays []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Add test cases.
		{name: "test", args: args{arrays: []int{1,2,3,4,5}, target: 4}, want: []int{0,2}},
		// arrays must be a sorted arrays
		//{name: "test2", args: args{arrays: []int{2,1,3,4,5}, target: 4}, want: []int{0,2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.arrays, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
