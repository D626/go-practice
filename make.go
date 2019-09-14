package main

import (
	"fmt"

)

//make 创建slice,map,channel,interface
func main() {
	s:=make([]int,0)
	fmt.Println(s==nil)
	fmt.Printf("%p \n",s )
fmt.Println(len(s),cap(s))
s1:=make([]string,0)
fmt.Println(len(s1),cap(s1))
s1=append(s1,"a","b")
fmt.Println(len(s1),cap(s1))
	s1=append(s1,"e","r","e","r","r","e","r","r","e","r","r","r","e","r","r","e","r","e","r","r","e","r","r","e")//不超出范围时容量翻倍，超出范围len=cap
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	//添加新的切片
	s2:=make([]int,0)
	s3:=[]int{1,2,3,4}
s2=append(s2,s3...)
fmt.Println(s2)
//通过数组产生切片
num:=[5]int{1,2,3,4,5}
s4:=num[0:2]
fmt.Println(s4)
fmt.Printf("%T\n",s4)//切片是指针
fmt.Printf("%p %p\n",s4,&num[0])
s4[0]=6
fmt.Println(num,s4)
s4=append(s4,7)
fmt.Println(num,s4)
	s4=append(s4,7,8,9,2)//切片放不下，数组重新开辟空间
	fmt.Println(num,s4)
	//删除代码

	n:=3//删除元素的索引
	newSlice:=s4[0:n]
	newSlice=append(newSlice,s4[n+1:]...)
	fmt.Println(s4)//原切片内容跟随变化，不要使用原切片
	fmt.Println(newSlice)
	fmt.Printf("%p %p",s4,newSlice)
}

