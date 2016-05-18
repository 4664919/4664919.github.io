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
	SelectSort()
}
func SelectSort() {
	var data []byte = []byte("SORTEXAMPLE")
	//fmt.Println("once sort:", string(data))
	for i, d := range data {
		tmp_d := d
		tmp_i := i
		for ii, dd := range data {
			if i > ii {
				continue //已经排序好的就跳过
			}
			//if dd < tmp_d {
			if less(dd, tmp_d) {
				tmp_d = dd
				tmp_i = ii
			}
		}
		//交换位置
		//data[i] = tmp_d
		//data[tmp_i] = d
		exch(data, i, tmp_i)
		//fmt.Println("once sort:", string(data))
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
