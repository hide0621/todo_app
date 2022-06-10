//「users」というtableの作成をしたい

package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

//データベース名をグローバルに指定
var Db *sql.DB

var err error

//table名を定数で指定
const (
	tableNameUser = "users"
)

func init() {

	//ドライバーとデータベースファイルの指定
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	//「users」というtableの作成をするコマンドを作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	//「users」というtableの作成
	Db.Exec(cmdU)
}
