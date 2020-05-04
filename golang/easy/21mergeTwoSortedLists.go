package easy

import (
	"github.com/xieydd/LeetCode/golang/datastruct/list"
)

func Merge2SortedLists(l1, l2 *list.List) *list.List {

	var e *list.List
	result := e.Init()
	if l1.Len() == 0 && l2.Len() != 0 {
		return l2
	}

	if l1.Len() != 0 && l2.Len() == 0 {
		return l1
	}


	for (l1.Len() == 0 && l2.Len() == 0){
		if (l1.root.Value > l2.root.Value) {
			result.root.next = l1.root
			l1.root = l1.root.next
			l1.Len()--
		}else {
			result.root.next = l2.root
			l2.root = l2.root.next
			l2.Len()--
		}

		result = result.root.next
	}

	if (l1.Len() != 0) {
		result.root.next = l1.root
		l1.Len()--
	}

	if (l2.Len() != 0) {
		result.root.next = l2.root
		l2.Len()--
	}
	return result
}
