package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id      int
	Name    string
	Address string
}

/*
func main() {
	//runinsert()
	runSelect()
}*/

func runinsert() {
	db := getConnection()
	defer db.Close()
	//stmt, err := db.Prepare("insert into mystudent(name,address) values(?,?)")
	stmt, err := db.Prepare("update mystudent set name=?,address=? where id=?")
	if err != nil {
		panic(err.Error())
	}

	//stmt.Exec("Mahesh", "Canada")
	stmt.Exec("Mangesh", "Aus", 3)
}

func runSelect() {
	db := getConnection()
	defer db.Close()

	//selDB, err := db.Query("Select * from mystudent")
	selDB, err := db.Query("Select * from mystudent where id=?", 2)
	res := []Student{}
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id int
		var name, address string
		err := selDB.Scan(&id, &name, &address)
		if err != nil {
			panic(err.Error())
		}

		s := Student{id, name, address}
		res = append(res, s)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func getConnection() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "college"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
