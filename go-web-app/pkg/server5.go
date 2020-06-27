// https://www.yoheim.net/blog.php?q=20170403

package main

import (
  "net/http"
  "fmt"
)

func main() {

    // basic認証.
    http.HandleFunc("/basicAuth", handleBasicAuth)

    // 8080ポートで起動
    http.ListenAndServe(":8080", nil)
}

func handleBasicAuth(w http.ResponseWriter, r *http.Request) {

    // Basic認証のデータ取得.
    username, password, ok := r.BasicAuth()

    // そもそもそんなヘッダーがないなどのエラー.
    if ok == false {
        w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
        w.WriteHeader(http.StatusUnauthorized) // 401
        fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
        return
    }

    // Basic認証のヘッダーはあるけど、値が不正な場合.
    if username != "my" || password != "secret" {
        w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
        w.WriteHeader(http.StatusUnauthorized) // 401
        fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
        return
    }

    // OK
    fmt.Fprint(w, "OK ------")
}

// http://localhost:8080/basicAuth
