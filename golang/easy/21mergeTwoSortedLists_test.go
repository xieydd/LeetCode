package easy

import (
	"reflect"
	"testing"

	"github.com/xieydd/LeetCode/golang/datastruct/list"
)

func TestMerge2SortedLists(t *testing.T) {
	type args struct {
		l1 *list.List
		l2 *list.List
	}

	l1 := list.New()
	e := l1.PushFront(2)
	e2 := l1.PushFront(3)
	l2 := list.New()
	e3 := l2.PushFront(1)
	e4 := l3.PushFront(4)

	result := list.New()
	result.MoveToBack(e3)
	result.MoveToBack (e)
	result.MoveToBack(e2)
	result.MoveToBack(e4)

	tests := []struct {
		name string
		args args
		want *list.List
	}{
		//  Add test cases.
		{name: "test", args: {l1: &l1, l2: &l2}, want: &result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge2SortedLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge2SortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
