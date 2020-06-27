//C:>go build hello.go
//C:>hello.exe
package main

import (
    "log"
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/01", handler01)
	http.HandleFunc("/02", handler02)
	http.HandleFunc("/03", handler03)
	// DuckTyping的に、ServeHTTP関数があれば良い.
	http.Handle    ("/04", String("Duck Typing!!!"))
	http.HandleFunc("/05", handler05)
	http.HandleFunc("/06", handler06)
	http.HandleFunc("/07", handler07)

    // 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := `
		<a href="http://localhost:8080/01">01</a></br>
		<a href="http://localhost:8080/02">02</a></br>
		<a href="http://localhost:8080/03">03</a></br>
		<a href="http://localhost:8080/04">04</a></br>
		<a href="http://localhost:8080/05">05</a></br>
		<a href="http://localhost:8080/06">06</a></br>
		<a href="http://localhost:8080/07">07</a></br>
	`
	fmt.Fprintf(w, str)
}

func handler01(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type Page struct {
	Title string
	Count int
}

func handler02(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}
	tmpl, err := template.ParseFiles("templates/layout.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func handler03(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}                                       // テンプレート用のデータ
	tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count") // テンプレート文字列
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page) // テンプレートをジェネレート
	if err != nil {
		panic(err)
	}
}

type String string

// String に ServeHTTP 関数を追加
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func handler05(w http.ResponseWriter, r *http.Request) {

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
    fmt.Fprint(w, "OK")
}

func handler06(w http.ResponseWriter, r *http.Request) {

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

func handler07(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("templates/template000.html.tpl"))

    str := "Sample Message"

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "template000.html.tpl", str); err != nil {
        log.Fatal(err)
    }
}

// http://localhost:8080
