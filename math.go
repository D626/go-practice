package main
import (
	"fmt"
	"math"
)

func main() {
	var a,b float64 =12.3,13.5
fmt.Println(a,b)
fmt.Println(math.Floor(a))
fmt.Println(math.Ceil(a))
fmt.Println(math.Abs(a))
fmt.Println(math.Modf(a))
 fmt.Println(math.Max(a,b))//小数部分精度问题
fmt.Println(math.Pow(2,3))//x的y次方
fmt.Println(math.Round(a))//四舍五入
}
