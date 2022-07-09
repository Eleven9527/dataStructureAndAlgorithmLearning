package stack

import "testing"

func TestStackTest(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "stack_test", args: args{vals: []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StackTest(tt.args.vals)
		})
	}
}
