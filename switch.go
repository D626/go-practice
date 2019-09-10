package main

import (
	"fmt"

)

func main() {
	//变量取固定值

	/*num:=16
	switch num {
	case 2:
		fmt.Println("二进制 ")
	case 8:
		fmt.Println("十进制数")
	case 10:
		fmt.Println("十进制数")
	case 16:
		fmt.Println("十六进制数")
	default:
		fmt.Println("内容不对")
		fmt.Println("程序结束")
		//switch条件判断

		score:=67
		switch{
		case score<60:
			fmt.Println("不及格")
		case score>=60&&score<70:
			fmt.Println("及格")
		case score>=70&&score<80:
			fmt.Println("良好")
		case score>=100:
			fmt.Println("优秀" )
		}


	*/
	//switchzhong case支持多值用，分开
    month:=5
    switch month{
	case 1,3,5,7,8,10,12:
		fmt.Println("31天")
		fallthrough//穿透和break对立，fallthrough不能在最后
	case 2:
		fmt.Println("29或28天")
	case 4,6,9,11:
		fmt.Println("30天")
	}



	}

