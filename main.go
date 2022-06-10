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

	/*
		//「users」というtableのカラムに値を入れている（構造体Userのフィールドに紐付いている）
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		//「users」というtableのカラムに値が入った状態（valuesの？の部分を埋める）にしてからコマンドの実行
		u.CreateUser()
	*/

	//users tableにてidが「1」のユーザーを取得
	u, _ := models.GetUser(1)

	fmt.Println(u)
}
