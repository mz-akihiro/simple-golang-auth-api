package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Newdb() *sql.DB {
	//envファイルを読み込む、エラーした場合はプログラムを強制的に止める
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	//GO_ENV環境変数が"dev"の場合、.envファイルの変数が読み込まれる
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load() //	エラーしなければerr変数にnilが入る
		if err != nil {        //	nil変数じゃなかった場合（つまりエラーだった場合）、ログを出力
			log.Fatalln(err)
		}
	}

	dbconf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PW"), os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *sql.DB) {
	//dbをclose中にエラーが起きた場合のハンドリング
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Closed")
}
