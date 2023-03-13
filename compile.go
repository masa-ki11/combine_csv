package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 読み込むCSVファイルのパスを配列に格納
	files, err := filepath.Glob("*.csv")
	if err != nil {
		panic(err)
	}

	// 出力先のCSVファイルを開く
	f, err := os.Create("output.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 出力先のCSVファイルに書き込むためのWriterを作成
	w := csv.NewWriter(f)
	defer w.Flush()

	// 読み込むCSVファイルの数だけループを回す
	for file_count, file := range files {
		// CSVファイルを開く
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// CSVファイルから読み込むためのReaderを作成
		r := csv.NewReader(f)
		rows, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
		}

		// CSVファイルから1行ずつ読み込んで処理
		for i, record := range rows {
			if file_count == 0 {
				w.Write(record)
			} else {
				if i != 0 {
					w.Write(record)
				}
			}
		}
	}
}
