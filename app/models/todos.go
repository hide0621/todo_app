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
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}
