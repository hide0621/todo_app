//html/templeteパッケージについてはこのサイトが分かりやすい
//https://code-database.com/knowledges/87

package controllers

import "net/http"

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
	_, err := session(w, r)
	if err != nil {
		//ログインしていなければ(セッションがなければ)トップページにリダイレクトされる
		http.Redirect(w, r, "/", 302)
	} else {
		//セッションがあればindex.htmlを表示する
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
