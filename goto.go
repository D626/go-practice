package main

import (
	"fmt"

)
//大量使用goto，编译器会不断产生跳跃性变化
func main() {
	for i := 5;i<7;i++{
		if i == 6 {
			goto loop
		} else {
	goto loop1}

	loop:
		fmt.Println("loop")
	loop1:

			fmt.Println("执行loop1")

	}
	fmt.Println("程序结束")
}