package singleSortedLinkedList

import (
	"reflect"
	"testing"
)

func TestJosephuLoop(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "josephuLoop", args: args{[]int{1, 3, 5, 2, 4, 6}}, want: []int{3, 6, 4, 2, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JosephuLoop(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JosephuLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
1 2 3 4 5 6
1 2 4 5
1 2 5
1 5
*/
