package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {

	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(writer, "layout", data)
}

//cookieを取得する関数
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	//httpリクエストからcookieを取得する
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		sess = models.Session{UUID: cookie.Value}
		//上記で受け取ったセッションがDB上にあるセッションと同じか確認
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("Invalid session")
		}
	}
	//上記のセッションがDB上にあればerrは返ってこない
	return
}

//URLの正規表現のパターンをコンパイルしておく
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+$)")

//これもパターンとして使われる
//ハンドラ関数（無名関数）を返す関数として定義されている
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	//このハンドラ関数にparseURL関数の処理を書いていく
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		//validPathとURLがマッチした部分をスライスで取得する
		q := validPath.FindStringSubmatch(r.URL.Path)
		//何もマッチしない場合
		if q == nil {
			http.NotFound(w, r)
			return
		}
		//スライスqのインデックス番号2番をIDとして受け取り、qiはint型になる
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		//関数呼び出し（関数オブジェクトのtodoEdit関数を呼び出して実行する）
		//parseURL関数の引数を渡して実行している(parseURL関数の引数と関数オブジェクトのtodoEdit関数の引数は一致しなければならない)
		//一番インプルな無名関数の実行のイメージ　fn := func(w http.ResponseWriter, r *http.Request, int) {}　fn()　と同じ意味
		//無名関数については以下のサイトが分かり易い
		//https://qiita.com/elephant_dev/items/64e301a5668d6e593429
		//https://go.dev/tour/moretypes/24
		fn(w, r, qi)
	}
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
	http.HandleFunc("/logout", logout)

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/authenticate", authenticate)

	//ハンドラ関数を実行するURLの登録
	//ログインしているユーザーしかtodosのページにアクセスできない
	http.HandleFunc("/todos", index)

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/todos/new", todoNew)

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/todos/save", todoSave)

	//ハンドラ関数を実行するURLの登録
	//parseURLとtodoEditはハンドラ関数をチェインさせて実行している
	//故にparseURL関数の戻り値がhttp.HandleFunc型となっている
	//ハンドラ関数のチェインについては以下のサイトが分かり易い
	//https://qiita.com/shzawa/items/92279fc06ca3f6aade28
	http.HandleFunc("todos/edit/", parseURL(todoEdit)) //URL末尾にスラッシュがない場合、URLが完全一致することを求められる。URL末尾にスラッシュがあれば、要求されたURLの先頭が登録されたURLと一致するかどうか調べる。つまりスラッシュがあればスラッシュの後に何か来ていても処理をハンドラ関数に渡すことが出来る。

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))

	//ハンドラ関数を実行するURLの登録
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	//Heroku用
	//環境変数PORTの取得
	port := os.Getenv("PORT")

	//ポート番号を指定してサーバーの立ち上げ
	return http.ListenAndServe(":"+port, nil) //nilとすることでマルチプレクサを使用する。登録されていないURLにアクセスしたらデフォルトで404エラーを返す。

}
