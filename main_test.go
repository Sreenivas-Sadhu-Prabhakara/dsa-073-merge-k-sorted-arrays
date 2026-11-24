package main

import (
	"reflect"
	"testing"
)

func TestMergeK(t *testing.T) {
	cases := []struct {
		in   [][]int
		want []int
	}{
		{[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[][]int{{1, 2}, {}, {0}}, []int{0, 1, 2}},
		{[][]int{}, []int{}},
		{[][]int{{}, {}}, []int{}},
		{[][]int{{5}}, []int{5}},
		{[][]int{{1, 1}, {1}}, []int{1, 1, 1}},
	}
	for _, c := range cases {
		if got := MergeK(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeK(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
