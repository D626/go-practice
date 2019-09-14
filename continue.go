package main

import "fmt"

func main() {
for i:=0;i<5;i++ {
		if i==2{
			break
		}
		fmt.Println(i)
	}
	//break在双重for循环
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if j==1{
				break
			}
			fmt.Println(i,j)
		}
	}
label:for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if j==1{
				break label
			}
			fmt.Println(i,j)
		}
	}
	fmt.Println("程序结束")
}
