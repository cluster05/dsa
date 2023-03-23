package stack

import (
	"reflect"
	"testing"
)

type fields struct {
	items  []int
	length int
}

var (
	emptyStack = fields{items: []int{}, length: 0}
	fillStack  = fields{items: []int{10, 20, 30}, length: 3}
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{
			name: "Init Stack",
			want: NewStack(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_ElementAt(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      int
		isPresent bool
	}{
		{
			name:      "Element At when stack is empty",
			fields:    fields{},
			args:      args{pos: 1},
			want:      InvalidIndex,
			isPresent: false,
		},
		{
			name:      "Element At when stack is not empty",
			fields:    fillStack,
			args:      args{pos: 1},
			want:      30,
			isPresent: true,
		},
		{
			name:      "Element At when stack is not empty",
			fields:    fillStack,
			args:      args{pos: 2},
			want:      20,
			isPresent: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items:  tt.fields.items,
				length: tt.fields.length,
			}
			got, got1 := s.ElementAt(tt.args.pos)
			if got != tt.want {
				t.Errorf("ElementAt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.isPresent {
				t.Errorf("ElementAt() got1 = %v, want %v", got1, tt.isPresent)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		{
			name:   "Pop Element when stack is Empty",
			fields: emptyStack,
			want:   InvalidIndex,
			want1:  false,
		},
		{
			name:   "Pop Element when stack contain elements",
			fields: fillStack,
			want:   30,
			want1:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items:  tt.fields.items,
				length: tt.fields.length,
			}
			got, got1 := s.Pop()
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Push element when stack empty",
			fields: emptyStack,
			args: args{
				item: 10,
			},
		},
		{
			name:   "Push element when stack is not empty",
			fields: fillStack,
			args: args{
				item: 40,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items:  tt.fields.items,
				length: tt.fields.length,
			}
			beforeLength := s.length
			s.Push(tt.args.item)

			got, _ := s.Top()
			expected := tt.args.item
			if got != expected {
				t.Errorf("got %q expected %q", got, expected)
			}

			afterLength := s.length

			if beforeLength+1 != afterLength {
				t.Errorf("got before length + 1 %q expeted after length %q", beforeLength, afterLength)
			}

		})
	}
}

func TestStack_Top(t *testing.T) {
	tests := []struct {
		name      string
		fields    fields
		want      int
		isPresent bool
	}{
		{
			name:      "Top when stack is empty",
			fields:    emptyStack,
			want:      InvalidIndex,
			isPresent: false,
		},
		{
			name:      "Top when stack is not empty",
			fields:    fillStack,
			want:      30,
			isPresent: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items:  tt.fields.items,
				length: tt.fields.length,
			}
			got, got1 := s.Top()
			if got != tt.want {
				t.Errorf("Top() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.isPresent {
				t.Errorf("Top() got1 = %v, want %v", got1, tt.isPresent)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "IsEmpty when stack is Empty",
			fields: emptyStack,
			want:   true,
		},
		{
			name:   "IsEmpty when stack is not Empty",
			fields: fillStack,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				items:  tt.fields.items,
				length: tt.fields.length,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
