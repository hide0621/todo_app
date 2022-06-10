package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {

	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		//ログのテスト
		log.Println("test")
	*/

	//base.goのinit関数の呼び出しをしたいので（webapp.sqlというファイルを作成したので、何でも良いから簡単なプログラムを実行している）
	fmt.Println(models.Db)
}
