package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//読み書き、ファイルの作成、ファイルへの追記、パーミッションの定義を行なってlogファイルの作成
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln(err)
	}

	//ログの書き込み先を標準出力とlogfileに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	//ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//ログの出力先をmultiLogFilに指定
	log.SetOutput(multiLogFile)
}
