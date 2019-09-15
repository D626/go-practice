package main

import (
	"sort"
	"fmt"

)

func main() {

	num:=[]int{1,7,3,5,2}
	sort.Ints(num)//升序
	fmt.Println(num)
sort.Sort(sort.IntSlice(num))
    fmt.Println(num)

   sort.Sort(sort.Reverse(sort.IntSlice(num)))//降序
   fmt.Println(num)
   //float类型切片排序
   f:=[]float64{1.2,9.4,5.6}
   sort.Float64s(f)
   fmt.Println(f)
   sort.Sort(sort.Reverse(sort.Float64Slice(f)))
   fmt.Println(f)
   //string类型排序
   s:=[]string{"a","j","a我们","我们" }//升序
   sort.Strings(s)
   fmt.Println(s)
   //切片应该是一个升序排序的切片
  n:= sort.SearchStrings(s,"a我们")//返回所在位置
  fmt.Println(n)
   sort.Sort(sort.Reverse(sort.StringSlice(s)))
   fmt.Println(s)
   //
}
