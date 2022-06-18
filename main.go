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

	//base.goのinit関数の呼び出しをしたいので記述している
	//webapp.sqlというファイルを作成したので、何でも良いから簡単なプログラムを実行している
	//このコードの実行がトリガーとなり「users」と「todos」というテーブルがsqlite3上に作成される
	fmt.Println(models.Db)

	/*
		//userの作成のクエリを実行する
		//「users」というtableのカラムに値を入れている（構造体Userのフィールドに紐付いている）
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		//「users」というtableのカラムに値が入った状態（valuesの？の部分を埋める）にしてからコマンドの実行
		u.CreateUser()
	*/

	/*
		//users tableにてidが「1」のユーザーを取得
		u, _ := models.GetUser(1)

		fmt.Println(u)

		//users tableのカラムの値を更新する
		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		//users tableのカラムの値を消去する
		u.DeleteUser()
		//消去できているか確認
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		//前回「users」テーブルを作るセクションにてuserを作成の後、削除した
		//userのIDは自動で増分されるので、今回のIDは2番とする
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		//指定したUser_idのuserに対してtodoを作成するクエリを実行する
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		//usr_idが3番のuserのtodoを取得する
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		//idが1番のTodo構造体を持って来て、そのContentフィールドを書き換えて、更新のクエリを実行する
		//今回はuser_idは更新していない
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	//idが3番のtodoを削除するクエリを実行する
	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
