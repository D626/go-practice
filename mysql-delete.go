package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.133:6033)/mydb")
	db.Ping()
	if err!=nil{
		fmt.Println("连接失败")
		return
	}
	defer func() {
		if db!=nil{
			db.Close()
		}
	}()
	stmt,err:=	db.Prepare("delete from people where id=?")
	if err!=nil{
		fmt.Println("预处理出错")
		return
	}
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	r,err:=stmt.Exec(6)
	if err!=nil{
		fmt.Println("执行失败")
		return
	}
	count,err:=r.RowsAffected()
	if count>0{
		fmt.Println("删除成功")

	}else {
		fmt.Print("删除失败")
	}
}