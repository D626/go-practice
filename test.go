package main

import (
	"fmt"
	"math/cmplx"
	"math"
)
var (aa=3
	ss="kkk"
	bb=true)
func variableZeroValue(){
	var a int
	var s string
	fmt.Println(a,s)

}
func variableInitialValue() {
	var a,b int=3,4
	var s string="abc"
	fmt.Println(a,b,s)

}
func variableTypeDeduction(){
	var a,b,c,s=3,4,true,"def"
	fmt.Println(a,b,c,s)
}
func variableShorter(){
	a,b,c,s:=3,4,true,"def"//冒号定义
	b=5
	fmt.Println(a,b,c,s)

}
func euler(){
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
}
func triangle(){
	var a,b int=3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))//强制类型转换
	fmt.Println(c)
}
func consts(){
	const filename string ="abc.txt"
	const a,b int =3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(filename,c)
}
func enums() {
	const(
		cpp=0
		java=1
		python=2
		golang=3
	)
	const(xb=1<<(10*iota)
		kb
		mb
		gb
		tb
		pb

	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(kb,mb,gb, tb,pb)
}
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,ss,bb)
	euler()
	triangle()
	consts()
	enums()
}

