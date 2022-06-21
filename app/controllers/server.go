package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {

	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)
}

//サーバーの立ち上げ
func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//ハンドラ関数を実行するURLの登録
	//パス以下にアクセスしたら「top.html」のwebページを表示するハンドラ関数を実行する
	http.HandleFunc("/", top) //第二引数がハンドラ関数

	//ハンドラ関数を実行するURLの登録
	//パス以下にアクセスしたら「signup.html」のwebページを表示するハンドラ関数を実行する
	http.HandleFunc("/signup", signup)

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/login", login)

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/authenticate", auhenticate)

	//ポート番号を指定してサーバーの立ち上げ
	return http.ListenAndServe(":"+config.Config.Port, nil) //nilとすることでマルチプレクサを使用する。登録されていないURLにアクセスしたらデフォルトで404エラーを返す。

}
