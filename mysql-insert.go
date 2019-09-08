package main
import (

	_"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)
func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}
func main() {
	//打开连接
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.133:6033)/mydb")
	db.Ping()
	defer func(){
		if db !=nil{
			db.Close()
		}
	}()
	//2.预处理SQL
	stmt,err:=db.Prepare("insert into people values(default,?,?)")
	defer func(){
		if stmt !=nil{
			stmt.Close()
		}
	}()
	if err!=nil{
		fmt.Println("预处理失败！")
		return
	}
	r,err:=stmt.Exec("张三","海淀")
	if err!=nil{
		fmt.Println("sql执行失败！")
		return
	}
	count,err:=	r.RowsAffected()
	if err!=nil{
		fmt.Println("结果获取失败！")
		return
	}
	if count >0{
		fmt.Println("新增成功")
	}else{
		fmt.Println("新增失败")
		}
		//获取新增值的主键值
		id,_:=r.LastInsertId()
		fmt.Println(id)
		//4.关闭
		stmt.Close()
		db.Close()



}