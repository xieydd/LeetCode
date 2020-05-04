package list 

import (
	"reflect"
	"testing"
)

func TestElement_Prev(t *testing.T) {
	type fields struct {
		next  *Element
		prev  *Element
		list  *List
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				next:  tt.fields.next,
				prev:  tt.fields.prev,
				list:  tt.fields.list,
				Value: tt.fields.Value,
			}
			if got := e.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Element.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElement_Next(t *testing.T) {
	type fields struct {
		next  *Element
		prev  *Element
		list  *List
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				next:  tt.fields.next,
				prev:  tt.fields.prev,
				list:  tt.fields.list,
				Value: tt.fields.Value,
			}
			if got := e.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Element.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Init(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("List.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Head(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	tests := []struct {
		name   string
		fields fields
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.Head(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Back(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	tests := []struct {
		name   string
		fields fields
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.Back(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_insert(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e  *Element
		at *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.insert(tt.args.e, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_insertValue(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		v  interface{}
		at *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.insertValue(tt.args.v, tt.args.at)
		})
	}
}

func TestList_remove(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.remove(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_move(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e  *Element
		at *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.move(tt.args.e, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.Remove(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PushFront(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.PushFront(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PushPrev(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.PushPrev(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.PushPrev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertBefore(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		value interface{}
		at    *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.InsertBefore(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.InsertBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertAfter(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		value interface{}
		at    *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			if got := l.InsertAfter(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.InsertAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_MoveToFront(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.MoveToFront(tt.args.e)
		})
	}
}

func TestList_MoveToBack(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.MoveToBack(tt.args.e)
		})
	}
}

func TestList_MoveBefore(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e  *Element
		at *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.MOveBefore(tt.args.e, tt.args.at)
		})
	}
}

func TestList_MoveAfter(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		e  *Element
		at *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.MoveAfter(tt.args.e, tt.args.at)
		})
	}
}

func TestList_PushBackList(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		other *List
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.PushBackList(tt.args.other)
		})
	}
}

func TestList_PushFrontList(t *testing.T) {
	type fields struct {
		length int
		root   Element
	}
	type args struct {
		other *List
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				length: tt.fields.length,
				root:   tt.fields.root,
			}
			l.PushFrontList(tt.args.other)
		})
	}
}

