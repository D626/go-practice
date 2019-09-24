package main

import "fmt"

func main() {
	//调用不接收
	show()
	//接收
	a,b,c:=show()
	fmt.Println(a,b,c)
	//占位符不接受接收
	_,e,d:=show()
	fmt.Println(e,d)

	w,s,j:=show1()
	fmt.Println(w,s,j)
	//可变参数函数
	/*1.支持可变参数
	fun 函数（参数，参数，名称...类型）{
	//函数体
	}
	2.声明函数时，在函数体把可变参数当切片使用
	3.可变参数必须存在其他参数后面，一个函数不能有多个可变参数
	因为前面普通参数个数确定的 编译器知道哪个给实参哪个给形参
	 */
	demo("张三","看书","写作业","发呆")
	/*匿名函数
	1.正常函数通过名称多次调用，匿名函数没函数名大部分情况都是在当前位置声明并立即调用
	（函数变量除外）
	fun（）{}（）//括号表示调用*/
	func(){
		fmt.Println("无参数无返回值匿名函数")
	}()
	//有参数的匿名函数
	func(name string){
		fmt.Println("名字为",name)
	}("nancy")//括号必须紧挨着
	//有返回值匿名函数

	name:=func() string{
		fmt.Println("有返回值匿名函数")
		return "李明"
	}()//调用
	fmt.Println(name)
}
func show()(string,int,string)  {
	fmt.Println("执行了show")
	return "Lily",13,"武汉"
}
func show1() (name string,age int ,addr string ) {
	fmt.Println("执行了show1")
	name="lucy"
	age=12
	addr="changsha"
	return

}
func  demo(name string,hover ... string){
	fmt.Println(name,"的晚上工作")
	for i,n:=range hover{

		fmt.Println(i,n)//循环遍历

	}
}