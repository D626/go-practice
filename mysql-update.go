package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main(){

	db,_:=sql.Open("mysql","root:123456@tcp(192.168.1.133:6033)/mydb")
	defer func(){
		if db!=nil{
			db.Close()
		}

	}()
	stmt,_:=db.Prepare("update people set name =?,address=? where id=?")
	defer func(){
		if stmt!=nil{
			stmt.Close()
		}
	}()
	r,_:=stmt.Exec("王五","长沙",4 )
	count,_:=r.RowsAffected()
	if count>0{
		fmt.Println("修改成功")
	}else{
		fmt.Println("操作失败")
	}
}
