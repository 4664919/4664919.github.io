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
	InsertSort()

}

func InsertSort() {
	var data []byte = []byte("SORTEXAMPLE")
	for i, _ := range data {
		for j := i; j > 0; j-- {
			//交换
			//if data[j] < data[j-1] {
			if less(data[j], data[j-1]) {
				exch(data, j, j-1)
			}
		}
		//fmt.Println("once sort:", string(data))
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
