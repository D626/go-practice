package main

import "fmt"

func main() {
	s1:=[]int{1,2}
	s2:=[]int{3,4,5,6,}
	copy(s1,s2[1:3])
	fmt.Println(s1)
	fmt.Println(s2)//使用copy函数可以保证原切片内容不变
	
	s:=[]int{1,2,3,4,5,6}
	n:=3//要删除元素的索引
	newSlice :=make([]int,n)
	copy(newSlice,s[0:n])
newSlice=append(newSlice,s[n+1:]...)
fmt.Println(newSlice)
fmt.Println(s)

}
