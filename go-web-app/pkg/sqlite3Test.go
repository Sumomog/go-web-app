//package main
//
//import (
//    // ビルド時のみ使用する
//    "database/sql"
//    "log"
//
//    _ "github.com/mattn/go-sqlite3"
//)
//
//// DB Path(相対パスでも大丈夫かと思うが、筆者の場合、絶対パスでないと実行できなかった)
//const dbPath = "/Users/hoge/go/src/github.com/hoge/lesson/db/db.sql"
//
//// コネクションプールを作成
//var DbConnection *sql.DB
//
//// データ格納用
//type Blog struct {
//    id int
//    title string
//}
//
//func main() {
//    // Open(driver,  sql 名(任意の名前))
//    DbConnection, _ := sql.Open("sqlite3", dbPath)
//
//    // Connection をクローズする。(defer で閉じるのが Golang の作法)
//    defer DbConnection.Close()
//
//    // blog テーブルの作成
//    cmd := `CREATE TABLE IF NOT EXISTS blog(
//             id INT,
//             title STRING)`
//
//    // cmd を実行
//    // _ -> 受け取った結果に対して何もしないので、_ にする
//    _, err := DbConnection.Exec(cmd)
//
//    // エラーハンドリング(Go だと大体このパターン)
//    if err != nil {
//        // Fatalln は便利
//        // エラーが発生した場合、以降の処理を実行しない
//        log.Fatalln(err)
//    }
//}


package main

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main(){

	var dbfile string = "./test.db"

	os.Remove( dbfile )

//	db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil { panic(err) }

	_, err = db.Exec( `CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))` )
	if err != nil { panic(err) }

	_, err = db.Exec(
			`INSERT INTO "world" ("country", "capital") VALUES (?, ?) `,
			 "日本",
			 "東京",
			)
	if err != nil { panic(err) }


	stmt, err := db.Prepare( `INSERT INTO "world" ("country", "capital") VALUES (?, ?) ` )
	if err != nil { panic(err) }

	if _, err = stmt.Exec("アメリカ", "ワシントンD.C."); err != nil { panic(err) }
	if _, err = stmt.Exec("ロシア", "モスクワ"); err != nil { panic(err) }
	if _, err = stmt.Exec("イギリス", "ロンドン"); err != nil { panic(err) }
	if _, err = stmt.Exec("オーストラリア", "シドニー"); err != nil { panic(err) }
	stmt.Close()

	db.Close()
}