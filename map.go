package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m==nil)
	fmt.Printf("%p\n",m)
	m2:=map[string]string{"name":"Lily","address":"武汉"}
	m1:=map[string]string{
		"name":"Lucy",
		"address":"湖北",

	}
fmt.Println(m1)
	fmt.Println(m2)
	//实例化时赋值操作
	m3:=map[string]string{"张三":"1234456789","李四":"123456784"}
	fmt.Println(m3)
	m3["王五"]="1234543245"//key不存在时表示新增操作
	fmt.Println(m3)
	m3["李四"]="123321123321"//修改操作
	fmt.Println(m3)
	delete(m3,"李四")//删除操作  若删除内容不存在  不会报错  保持不变
	fmt.Println(m3)
	delete(m3,"赵六")
	fmt.Println(m3)
	//查看key 对应的value值，存在，返回value；不存在返回value类型默认值
	fmt.Println(m3["张三"])
	fmt.Println(m3["李雷"])
	value,ok:=m3["张三"]
fmt.Println(value,ok)

	//循环遍历
	for key,value:=range m3{
		fmt.Println(key,value)
	}

}
