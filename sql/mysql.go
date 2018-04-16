package sql

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func init(){

	DB = NewMySQL()
	log.Println("MySQL init")
}

func NewMySQL() *sql.DB{
	db, err := sql.Open("mysql","root:weijunhao@tcp(47.94.216.185:3306)/test?timeout=3s&readTimeout=3s&writeTimeout=3s")
	if nil != err{
		log.Println(err.Error())
	}

	err = db.Ping()
	if nil != err{
		log.Println(err.Error())
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db
}