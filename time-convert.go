package main
import("fmt"
"time"
)
func main() {
	t:=time.Now()
	fmt.Println(t)
	s:=t.Format("2006-01-02 15:04:05")//格式内容不能更改
	fmt.Println(s )
	s1:="2022-2-4 18:8:9"
     t,_=time.Parse("2006-1-2 15:4:5",s1)
fmt.Println(t)
}
