package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=dbuser password=dbuser dbname=dbuser sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("guowenshuai", "PCB&&SI", "2017-10-12")
	checkErr(err)

	var LastInsertId int
	err = db.QueryRow("insert into userinfo(username,departname,created) values($1,$2,$3) returning uid;", "wenshuai", "PCB&&SI", "2016-08-15",).Scan(&LastInsertId)
	checkErr(err)
	fmt.Println("last insert id = ", LastInsertId)

	fmt.Println("res: ", res)
	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println("id: ", id)

	// update data
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("wenshuai.guo", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Print("affect: ", affect)

	// query data
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)

		fmt.Println(uid, username, departname, created)
	}

	// del data
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("affect: ", affect)
	
	db.Close()
}


	


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
