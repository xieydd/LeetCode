package stack 

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		// TODO: Add test cases.
		{name: "stack1", want: NewStack()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "stack", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, wantValue: 1,wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			gotValue, err := stack.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Stack.Pop() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		value   interface{}
	}{
		// TODO: Add test cases.
		{name: "stack", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, value: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			stack.Push(tt.value)
		})
	}
}

func TestStack_Set(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	type args struct {
		idx   int
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "set", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, args: args{idx: 0, value: 1},  wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			if err := stack.Set(tt.args.idx, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Get(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "get", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, args: args{idx: 0}, wantValue: 1 , wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			gotValue, err := stack.Get(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Stack.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{name: "isempty", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			if got := stack.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Print(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "isempty", fields: fields{top: &node{value: 1, prev: nil}, length: 1, lock: sync.Mutex{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			if err := stack.Print(); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Print() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Swap(t *testing.T) {
	type fields struct {
		top    *node
		length int
		lock   sync.Mutex
	}
	type args struct {
		other *Stack
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "swapstack", fields: fields{top: &node{value: 1, prev: nil}, lock: sync.Mutex{}},args: args{other: NewStack()}, wantErr: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top:    tt.fields.top,
				length: tt.fields.length,
				lock:   tt.fields.lock,
			}
			if err := stack.Swap(tt.args.other); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

