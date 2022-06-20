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

	//htmlテンプレートとしてファイル名が「layout」と「public_navbar」と「login」のものを使用
	generateHTML(w, nil, "layout", "public_navbar", "login")

}
