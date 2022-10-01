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
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

//「users」というtableのカラムに値を挿入するコマンドを用意している関数
func (u *User) CreateUser() (err error) {

	cmd := `insert into users (
		uuid, 
		name,
		email,
		password,
		created_at) values ($1,$2,$3,$4,$5)`

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
	from users where id = $1`

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
	cmd := `update users set name = $1, email = $2 where id = $3`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//users tableにてidで絞り込んで、そのidに紐付けられた他のカラムの値を消去するクエリを実行する
func (u *User) DeleteUser() (err error) {

	//idで絞り込んで、そのidに紐付けられた他のカラムの値を消去するコマンドの定義
	cmd := `delete from users where id = $1`

	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ログイン画面でemailアドレスを受け取り、それを下にしてuserをDBから持ってくる
func GetUserByEmail(email string) (user User, err error) {

	user = User{}

	cmd := `select id, uuid, name, email, password, created_at
			from users where email = $1`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

//セッションの作成
func (u *User) CreateSession() (session Session, err error) {

	session = Session{}

	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values ($1,$2,$3,$4) `

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Println(err)
	}

	//セッションを持ってくる
	cmd2 := `select id, uuid, email, user_id, created_at
			 from sessions where user_id = $1 and email = $2`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err

}

//セッションがDBにあるかどうか確認する
func (session *Session) CheckSession() (valid bool, err error) {

	cmd := `select id, uuid,email, user_id,created_at
			from sessions where uuid = $1`

	err = Db.QueryRow(cmd, session.UUID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt) //Scanメソッドはsessionに渡してあげる役割

	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return valid, err
}

//セッションを削除する
func (sess *Session) DeleteSessionByUUID() (err error) {

	cmd := `delete from sessions where uuid  = $1`

	_, err = Db.Exec(cmd, sess.UUID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {

	user = User{}

	cmd := `select id, uuid, name, email, created_at FROM users
			where id = $1`

	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}
