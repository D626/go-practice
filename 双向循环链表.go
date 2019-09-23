package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//代表整个循环链表，又代表第一个元素
	r:=ring.New(5)
	r.Value=0
	r.Next().Value=1
	r.Next().Next().Value=2
	r.Next().Next().Next().Value=3
	r.Next().Next().Next().Next().Value=4
	r.Prev().Value=4
	r.Prev().Prev().Value=3
	//循环链表有几个元素 func 执行几次，i代表当前执行元素内容

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
	fmt.Println(r.Move(0).Value)
	//增加
	r1:=ring.New(1)
	r1.Value=5
	r.Link(r1)

	//删除
	r.Unlink(2)
	// 循环链表有几个元素 func 执行几次，i代表当前执行元素内容
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
