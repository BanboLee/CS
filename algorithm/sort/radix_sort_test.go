package sort

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{[]int{6, 4, 2, 3, 1}}, []int{1, 2, 3, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadixSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
