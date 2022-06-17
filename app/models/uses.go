package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

//「users」というtableのカラムに値を挿入するコマンドを用意している関数
func (u *User) CreateUser() (err error) {

	cmd := `insert into users (
		uuid, 
		name,
		email,
		password,
		created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//users tableにて指定されたidの値で絞り込んでクエリを実行
func GetUser(id int) (user User, err error) {

	//前のセクション（コミット）で値は挿入済み
	user = User{}

	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`

	//クエリの実行
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

//users tableにて指定したidのnameとemailを更新するクエリを実行
func (u *User) UpdateUser() (err error) {

	//指定したidのnameとemailを更新するコマンドの定義
	cmd := `update users set name = ?, email = ? where id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//users tableにてidで絞り込んで、そのidに紐付けられた他のカラムの値を消去するクエリを実行する
func (u *User) DeleteUser() (err error) {

	//idで絞り込んで、そのidに紐付けられた他のカラムの値を消去するコマンドの定義
	cmd := `delete from users where id = ?`

	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
