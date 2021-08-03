package sort

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{[]int{9, 3, 4, 2, 1, 1, 7, 1, 9, 8}, 10}, []int{1, 1, 1, 2, 3, 4, 7, 8, 9, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountingSort(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountingSortNegative(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{[]int{-5, 3, 4, 2, 1, 1, 7, -1, -5, -2}}, []int{-5, -5, -2, -1, 1, 1, 2, 3, 4, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountingSortNegative(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountingSortNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
