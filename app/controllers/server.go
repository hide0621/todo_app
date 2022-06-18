package controllers

import (
	"net/http"
	"todo_app/config"
)

//サーバーの立ち上げ
func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//ハンドラ関数を実行するURLの登録
	//パス以下にアクセスしたら「top.html」のwebページを表示するハンドラ関数を実行する
	http.HandleFunc("/", top) //第二引数がハンドラ関数

	//ポート番号を指定してサーバーの立ち上げ
	return http.ListenAndServe(":"+config.Config.Port, nil) //nilとすることでマルチプレクサを使用する。登録されていないURLにアクセスしたらデフォルトで404エラーを返す。

}
