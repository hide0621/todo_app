//html/templeteパッケージについてはこのサイトが分かりやすい
//https://code-database.com/knowledges/87

package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

/*
//ハンドラ関数の定義
//引数はパターン
func top(w http.ResponseWriter, r *http.Request) {

	//指定されたパスにあるhtmlファイルを読み込んで変数tに代入
	//詳しい処理は「https://qiita.com/ryuji0123/items/be0a0d09fa432dab1462」が詳しい
	t, err := template.ParseFiles("app/views/templates/top.html")

	//上記のParseFiles関数は戻り値にerrを返すのでエラーハンドリングの実装が推奨されている
	if err != nil {
		log.Fatalln(err)
	}

	//上記のhtmlファイルを表示する
	t.Execute(w, "Hello") //第二引数で渡したdateは,上記パスのhtmlファイルにて{{.}}とすることで、その第二引数のdateを渡すことができる
}
*/

//ハンドラ関数の定義
//引数はパターン
//定義したlayoutテンプレートとtopテンプレートを用いたハンドラ関数の実装
func top(w http.ResponseWriter, r *http.Request) {

	//cookieを取得
	_, err := session(w, r)
	if err != nil {
		//ログインしていない(DBにセッションがない)ということなので、topページにアクセスする
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		//ログインしていればtodosのページへアクセスする
		http.Redirect(w, r, "/todos", 302)
	}

	//dataとしてHelloを登録、htmlテンプレートとしてファイル名が「layout」と「public_navbar」と「top」のものを使用
	generateHTML(w, "Hello", "layout", "public_navbar", "top")

}

//index.htmlを表示するハンドラ関数
func index(w http.ResponseWriter, r *http.Request) {
	//ログインしているかどうか判定する
	//セッションを取得してブラウザのcookieと一致するかどうかチェックする
	sess, err := session(w, r)
	if err != nil {
		//ログインしていなければ(セッションがなければ)トップページにリダイレクトされる
		http.Redirect(w, r, "/", 302)
	} else {
		//セッションのユーザーIDを取得して、それと一致するユーザーを変数userへ代入
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//ユーザーが作成したtodoの一覧を取得
		todos, _ := user.GetTodosByUser()

		//todoの一覧を構造体UserのTodosフィールドに入れる
		user.Todos = todos

		//セッションがあればindex.htmlを表示する
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

//ハンドラ関数
func todoNew(w http.ResponseWriter, r *http.Request) {
	//ログインの確認
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		//ログインしていれば
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

//ハンドラ関数
func todoSave(w http.ResponseWriter, r *http.Request) {
	//ログインの確認
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		//ログインしていればTodoCreateのフォームの値を取得する
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//ユーザーの取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Panicln(err)
		}
		//ログインしていればTodoCreateのフォームの値をPOSTする
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		//todoの一覧ページにリダイレクト
		http.Redirect(w, r, "/todos", 302)
	}
}

//関数オブジェクトとして定義
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	//セッションを確認
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		//ユーザーの確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//引数のidからtodoを取得する
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	//セッションの確認
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//ユーザーの取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
