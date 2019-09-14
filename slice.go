package main

import "fmt"

func main() {
	//声明时为空,切片相当于一个指针  赋值时传递的是地址，变量类似指针
	//值类型在变量间赋值时传递的是值的副本
	var slice []string
	fmt.Println(slice)
	fmt.Println(slice==nil)
	fmt.Printf("%p\n",slice)
	names:=[]string{"Lily","Lucy"}

	fmt.Println(names)
	fmt.Println(names[0],names[1])
	names1:=names
	names1[0]="Jim"
	fmt.Println(names,names1)
	fmt.Printf("%p\n%p\n",names,names1)
}
