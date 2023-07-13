package main

import (
	"fmt"
	"log"
	"simple-golang-auth-api/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)
	_, err := dbCnt.Exec(`CREATE TABLE IF NOT EXISTS user (
							id INT AUTO_INCREMENT PRIMARY KEY,
							email VARCHAR(255) UNIQUE NOT NULL,
							password VARCHAR(255) NOT NULL,
							created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
							updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
						)`)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("User table created successfully")
}

//上記のテーブル構築クエリについて解説（主にタイムスタンプ）
/*
`CREATE TABLE IF NOT EXISTS user (
							id INT AUTO_INCREMENT PRIMARY KEY,
							email VARCHAR(255) UNIQUE NOT NULL,
							password VARCHAR(255) NOT NULL,
							created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,								//作成時の時刻を自動で打刻し、以降明示的に変更されない限りそのまま
							updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP	//作成時の時刻を自動で打刻し、かつ行のカラムがどれか一つでも変更されたら、変更時の時刻が自動で打刻される
						)`

DEFAULT CURRENT_TIMESTAMP	自動で作成時の時刻を打刻

DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP	自動で作成時の時刻を打刻し、かつ同じ行のカラムがどれか一つでも変更されると変更時の時刻を自動で打刻する。

*/
