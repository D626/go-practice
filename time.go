package main

import (
	"fmt"
	"time"
)

//事件类型代表现实世界时间
func main() {
	var t time.Time
	t=time.Now()
	fmt.Println(t)
	//通过纳秒时间戳 创建时间变量
	t1:=time.Unix(0,t.UnixNano())
	fmt.Println(t1)
	//创建指定时间
	t2:=time.Date(2019,9,10,17,8,8,0,time.Local)
fmt.Println("滚回学校时间",t2)
fmt.Println("你想得真美😔")
fmt.Println(t2.Year())
fmt.Println(t2.Month())
fmt.Println(t2.Day())
fmt.Println(t2.Date())
fmt.Println(t2.Clock())
fmt.Println(t2.Nanosecond())
fmt.Println(t.Unix())//秒差 距离1970年1月1日
}
