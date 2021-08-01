package main

import "banbo/CS/data_structure/skiplist/skiplist_list"

func main() {
	cases := []int{1, 2, 3, 4}
	list := skiplist_list.NewList()
	for i, v := range cases {
		list.Add(i, v)
	}
}
