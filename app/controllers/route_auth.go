package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

//ハンドラ関数の定義
func signup(w http.ResponseWriter, r *http.Request) {

	//「/signup」へのリクエストの各種メソッド
	//GETメソッドの時
	if r.Method == "GET" {

		//cookieを取得
		_, err := session(w, r)
		if err != nil {
			//ログインしていない(DBにセッションがない)ということなので、signupページにアクセスする
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			//ログインしていればtodosのページへアクセスする
			http.Redirect(w, r, "/todos", 302)
		}

		//htmlテンプレートとしてファイル名が「layout」と「public_navbar」と「signup」のものを使用
		generateHTML(w, nil, "layout", "public_navbar", "signup")

		//POSTメソッドの時
	} else if r.Method == "POST" {

		//入力フォームの解析
		err := r.PostForm
		if err != nil {
			log.Fatalln(err)
		}

		//user登録
		user := models.User{
			//signup.htmlのinputタグにある各種の属性から任意の値を読み込む
			//参考：https://leben.mobi/go/post/practice/web/
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		//上記の内容でuserを作成
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		//上記userが作成されたらtopページにリダイレクト
		//ステータスコードは「302」
		http.Redirect(w, r, "/", 302)
	}
}

//ハンドラ関数の定義
func login(w http.ResponseWriter, r *http.Request) {

	//cookieを取得
	_, err := session(w, r)
	if err != nil {
		//ログインしていない(DBにセッションがない)ということなので、loginページにアクセスする
		//htmlテンプレートとしてファイル名が「layout」と「public_navbar」と「login」のものを使用
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		//ログインしていればtodosのページへアクセスする
		http.Redirect(w, r, "/todos", 302)
	}
}

//ハンドラ関数の定義
func auhenticate(w http.ResponseWriter, r *http.Request) {

	//リクエストパラメータを全て取得する
	err := r.ParseForm()

	//リクエストボディ中のemailアドレスを取得、それを下にDBから該当するuserを探してくる
	//参考：https://leben.mobi/go/post/practice/web/
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		//ログイン失敗によりログイン画面にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}

	//リクエストボディ中のemailアドレスを取得、それを暗号化して一致するかどうか確認する
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {

		//セッションの作成
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		//cookieとして持っておく情報の設定
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		//上記内容でcookieを作成する
		http.SetCookie(w, &cookie)

		//ログイン成功によりtopページにリダイレクト
		http.Redirect(w, r, "/", 302)

	} else {
		//ログイン失敗によりログイン画面にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
}
