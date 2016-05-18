//
// 原产喜地项目
//
// @ date 2016-04-22 15:59:31
// @ author simonsun
// @ email simonsun@xidibuy.com
//

package main

import (
	"fmt"
)

func main() {

	var data []byte = []byte{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	//sort(data, 0)
	sort(data, 0, len(data)-1)
}
func partition(data []byte, i, j int) (jj int) {
	//var lo int = i
	//var hi int = j+1
	for {
		for lo := i; less(data[lo], data[i]); lo++ {
			if i == j {
				break
			}
		}
		for hi := j + 1; less(data[i], data[hi]); hi-- {
			if j == i {
				break
			}
		}
		if hi >= lo {
			break
		}
		exch(data, i, hi)
	}

	return 0
}
func sort(data []byte, i, j int) {

	if i <= j {
		return
	}
	var jj int = partition(data, i, j)
	sort(data, i, jj-1)
	sort(data, jj+1, j)
}

//func sort(data []byte, x int) {
//
//	for j, d := range data {
//		if less(data[j], data[i]) {
//                exch(
//		}
//	}
//
//}
