package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 实例化list
	mylist := list.New()
	fmt.Print("初始化后", mylist)
	//添加
	mylist.PushFront("A") //打印【“A”】
	mylist.PushBack("b")  //[“A”,"b"】
	mylist.PushBack("c")  //[“A”,"b","c"】

	//往某个元素前后添加
	mylist.InsertBefore("d", mylist.Back()) //最后一个元素前面插上一个d，[“A”,"b","d","c"】
	mylist.InsertAfter("e", mylist.Front()) //第一个元素后面插上一个e，["A”,"e","b","d","c"】
	//遍历

	for e := mylist.Front(); e != nil; e = e.Next() {
		//e=e.next()	当前元素的下一个
		fmt.Println(e.Value, "") //当前元素内容

	}
	fmt.Println(mylist.Front().Value)
	fmt.Println(mylist.Back().Value)
	n := 5
	var curr *list.Element
	if n > 0 && n <= mylist.Len() {
		if n == 1 {
			curr = mylist.Front()
		} else if n == mylist.Len() {
			curr = mylist.Back()
		} else {
			curr := mylist.Front()
			for i := 1; i < n; i++ {
				curr = curr.Next()
			}
		}

	} else {
		fmt.Println("n的数值不对")
	}

	fmt.Println("取出第5个", curr.Value) //循环遍历找到第n个元素
	//移动元素
	mylist.MoveBefore(mylist.Front(), mylist.Back())//第一个放最后一个前面
	//遍历

	for e := mylist.Front(); e != nil; e = e.Next() {
		//e=e.next()	当前元素的下一个
		fmt.Println(e.Value, "") //当前元素内容

	}
	mylist.MoveAfter(mylist.Front(), mylist.Back())//第一个放最后一个元素后面
	//遍历
fmt.Println("第一个放最后一个元素后面")
	for e := mylist.Front(); e != nil; e = e.Next() {
		//e=e.next()	当前元素的下一个
		fmt.Println(e.Value, "") //当前元素内容

	}
	//删除
	mylist.Remove(mylist.Front())
	fmt.Println("删除第一个")
	for e := mylist.Front(); e != nil; e = e.Next() {
		//e=e.next()	当前元素的下一个
		fmt.Println(e.Value, "") //当前元素内容

	}
}