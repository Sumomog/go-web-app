package main

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "runtime"
    "sync"
)

var wg sync.WaitGroup
var limit chan struct{}

func init() {
    cpus := runtime.NumCPU()          // CPUの数を取得する
    limit = make(chan struct{}, cpus) // CPUの数 = 並列数
}

func main() {
    inputpath := filepath.Join("a", "b", "c")
    files, err := ioutil.ReadDir(inputpath)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        wg.Add(1) // 並列処理の登録
        go func(path string, filename string) {
            // 並列処理の中身
            defer wg.Done()
            limit <- struct{}{}
            input, err := os.Open(filepath.Join(path, filename))
            if err != nil {
                panic(err)
            }
            defer input.Close()
            ioutil.ReadAll(input) // テキストファイルを読み込む
            // テキスト処理を何かする
            <-limit
        }(inputpath, file.Name())
    }

    wg.Wait() //処理完了まで待機
}
