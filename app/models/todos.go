//Go言語によるSQLの実行(使用するメソッドの説明など)に関して以下のサイトがわかりやすい
//https://blog.suganoo.net/entry/2019/01/25/190200

package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

//「todos」というtableのカラムに値を挿入するコマンドを用意している関数
func (u *User) CreateTodo(content string) (err error) {

	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	//Execメソッドは単純にクエリを実行し、結果行を戻さないメソッド
	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//todos tableにて指定されたidの値で絞り込んでクエリを実行
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
			where id = ?`

	//前のセクション（コミット）で値は挿入済み
	todo = Todo{}

	//クエリの実行
	//一行のSQL結果が返ることが想定されるのでQueryRowメソッドとScanメソッドをメソッドチェーンにして使う
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

//todos tableにて複数の値をセレクトするクエリを実行する
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`

	//SQLの結果が複数行であることが想定される場合はQueryメソッドを使う
	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	//Nextメソッドを使って次の行に移動する
	for rows.Next() {

		var todo Todo

		//todos tableにて複数の値を取得する
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		//todosに入れ込む
		todos = append(todos, todo)

	}

	rows.Close()

	return todos, err
}
