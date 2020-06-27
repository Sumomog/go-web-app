// https://www.yoheim.net/blog.php?q=20170403

package main

import (
  "net/http"
  "fmt"
)

func main() {

    // パラメータ受け取り
    http.HandleFunc("/params", handleParams)

    // 8080ポートで起動
    http.ListenAndServe(":8080", nil)
}

func handleParams(w http.ResponseWriter, r *http.Request) {

    // クエリパラメータ取得してみる
    fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

    // Bodyデータを扱う場合には、事前にパースを行う
    r.ParseForm()

    // Formデータを取得.
    form := r.PostForm
    fmt.Fprintf(w, "フォーム：\n%v\n", form)

    // または、クエリパラメータも含めて全部.
    params := r.Form
    fmt.Fprintf(w, "フォーム2：\n%v\n", params)
}

// http://localhost:8080/params
