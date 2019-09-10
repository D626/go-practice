package main

import (
	"fmt"

)

func main() {
	//循环嵌套  每条语句都执行

	score:=65
	if( score>0&&score<60) {
		fmt.Println("不及格")

	}
		if (score>=60&&score<=70) {
			fmt.Println("及格")
		}
		if(score>=70&&score<=80){
			fmt.Println("良好")
		}else {
			fmt.Println("优秀")
		}
//if....elseif.....else跳过不满足条件语句
	if(score>0&&score<60) {
		fmt.Println("不及格")
		}else if (score>=60&&score<=70) {

		fmt.Println("及格")
	} else if(score>=70&&score<=80){
		fmt.Println("良好")
	}else {
		fmt.Println("优秀")
	}
	}


