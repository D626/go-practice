package main
import(
	"fmt"

)
/*
func main() {
var  arr1 [3]int=[3]int{1,2,3}
arr2:=[3]int{1,2,3}
arr3:=[4]int{1,2,3}
fmt.Println(arr1,arr2,arr3)
arr4:=[...]int{1,2,3}
fmt.Println(arr4)
fmt.Println(len(arr4))}
*/
func main()  {
	var arr [3][3]int=[3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
fmt.Println(arr)
arr0:=arr[0]
fmt.Println(arr0[0],arr0[1],arr0[2])
}