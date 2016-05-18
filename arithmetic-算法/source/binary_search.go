//
// 原产喜地项目
//
// @ date 2016-04-21 18:19:58
// @ author simonsun
// @ email simonsun@xidibuy.com
//

package main

import (
	"fmt"
)

func main() {
	var is_find bool
	is_find = BinarySearch(23)
	fmt.Println("------result", is_find)
	is_find = BinarySearch(50)
	fmt.Println("------result", is_find)
}
func BinarySearch(b byte) bool {
	fmt.Println("find:", b)
	var data []byte = []byte{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98}
	for {
		var _len int = len(data)
		_len_child := _len / 2
		v := data[_len_child]
		//fmt.Println(":", v)
		if v == b {
			return true
		}
		if _len_child == 1 {
			return false
		}
		if b < v {
			data = data[:_len_child]
		} else {
			data = data[_len_child:]
		}
	}
	return false
}
