package trial

import (
    "html/template"
    "log"
    "net/http"
)

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("templates/template000.html.tpl"))

    str := "Sample Message"

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "template000.html.tpl", str); err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/page0", htmlHandler0)

    // サーバーを起動
    http.ListenAndServe(":8989", nil)
}

// http://localhost:8989/page0