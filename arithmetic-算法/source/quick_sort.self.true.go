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
	"os"
)

type QuickSort struct {
	data []byte
}

func (qs *QuickSort) Exit() {
	os.Exit(1)
}
func (qs *QuickSort) Sort() {
	//随机选一个数 todo
	//目前手动设置第一个
	//qs.sort(qs.data,0,len(qs.data))
	qs.sort(0, 0, len(qs.data))
}

//  分割
func (qs *QuickSort) partition(base, begin, end int) (jj int) {
	var i int = base
	qs.Point("partition SORT:", qs.data[begin:end], "iiii:", i)
	for tmp := begin; tmp < end; tmp++ {
		//qs.Point("partition for tmp:", tmp, "i:", i, "")
		if tmp == base { //遇到自己，不处理
			//qs.Point("遇到自己，不处理")
			continue
		}
		if qs.data[tmp] < qs.data[base] {
			qs.Point("exch :", "tmp:", tmp, "i:", i, "qs.data[tmp] > qs.data[base]:", qs.data[tmp], qs.data[base])
			exch(qs.data, tmp, i)
			i++
			//break
			//continue
		}
	}
	//最后
	qs.Point("循环完毕，挪动base", "base:", base, "i:", i)
	exch(qs.data, base, i-1)
	//qs.Point("last:", "base:", base, "i:", i)
	qs.Point("last:", "partition返回i：", i)
	qs.Exit()
	return i
}
func (qs *QuickSort) sort(base, begin, end int) {
	if begin >= end {
		return
	}
	var pi int = qs.partition(base, begin, end)
	qs.sort(pi, begin, pi)
	qs.sort(pi, pi, end)
}

func (qs *QuickSort) GetData() []byte {
	return qs.data
}

func (qs *QuickSort) Point(info string, args ...interface{}) {
	fmt.Println(info, ":", (qs.data), "[other args:]", args)
}

func NewQuickSort(data []byte) *QuickSort {
	qs := new(QuickSort)
	qs.data = data
	return qs
}
func main() {

	//var data []byte = []byte{3, 7, 21, 2, 1, 10}
	var data []byte = []byte{3, 8, 5, 2, 4, 9}
	//var data []byte = []byte{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	qs := NewQuickSort(data)
	qs.Point("   source")
	//qs.Point("sort once")
	qs.Sort()
	fmt.Println(qs.GetData())

	//sort(data, 0, len(data)-1)
}
