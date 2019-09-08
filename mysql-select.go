package main

import (
"database/sql"
"fmt"
_"github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.133:6033)/mydb")
	if err!=nil{
		fmt.Print("连接失败")
	return
	}
	defer func(){
		if db !=nil{
			db.Close()
		}
	}()
	stmt,err:=db.Prepare("select * from people")
	if err!=nil{
		fmt.Print("预处理失败")
		return
	}
	defer func(){
		if stmt !=nil{
			stmt.Close()
		}
	}()
	rows,err:=stmt.Query()
	if err!=nil{
		fmt.Print("获取结果失败")
		return
	}
	defer func() {
		if rows!=nil{
			rows.Close()
		}
	}()
	for rows.Next(){
		var id int
		var name string
		var address string
		rows.Scan(&id,&name,&address)
		fmt.Println(id,name,address)
	}
	}
