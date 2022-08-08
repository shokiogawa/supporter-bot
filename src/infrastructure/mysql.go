package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase() (db *Database, err error) {
	db = new(Database)
	db.DB, err = tryConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func tryConnect() (db *sqlx.DB, err error) {
	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return
}

func (mysql *Database) Connect() (db *sqlx.DB, err error) {
	err = mysql.DB.Ping()
	if err != nil {
		// リトライ処理
		db, connerr := tryConnect()
		if connerr != nil {
			return nil, connerr
		}
		mysql.DB = db
	}
	return mysql.DB, nil
}
