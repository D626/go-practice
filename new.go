package main

import (
	"fmt"

)

func main() {
a:=new(int)
fmt.Println(a)//0xc000062080内存地址
*a=3
fmt.Print(*a )

var  b *int
fmt.Print(b)

}
