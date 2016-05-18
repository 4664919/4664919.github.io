//
// 原产喜地项目
//
// @ date 2016-04-22 15:20:37
// @ author simonsun
// @ email simonsun@xidibuy.com
//

package main

import (
//"fmt"
)

func less(a, b byte) bool {
	return a < b
}

func exch(data []byte, i, j int) {
	v := data[i]
	data[i] = data[j]
	data[j] = v
}
