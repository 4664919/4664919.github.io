//
// 原产喜地项目
//
// @ date 2016-04-27 10:22:17
// @ author simonsun
// @ email simonsun@xidibuy.com
//
// 队列，又称为伫列（queue），是先进先出（FIFO, First-In-First-Out）的线性表。在具体应用中通常用链表或者数组来实现。队列只允许在后端（称为rear）进行插入操作，在前端（称为front）进行删除操作。
//
// 队列的操作方式和堆栈类似，唯一的区别在于队列只允许新数据在后端进行添加。
//
package main

import (
	"fmt"
)

func main() {

	var data []byte = []byte{6, 3, 1, 7, 5, 8, 9, 2, 4}
	var res []byte

	var i int = 0
	for len(data) > 0 {
		i++
		tmp := data[0]
		data = data[1:]
		if i%2 == 1 {
			res = append(res, tmp)
			continue
		}
		data = append(data, tmp)
	}
	fmt.Println(res)

}
