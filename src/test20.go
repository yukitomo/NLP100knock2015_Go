/*Wikipediaの記事を以下のフォーマットで書き出したファイルjawiki-country.json.gzがある．
1行に1記事の情報がJSON形式で格納される
各行には記事名が"title"キーに，記事本文が"text"キーの辞書オブジェクトに格納され，そのオブジェクトがJSON形式で書き出される
ファイル全体はgzipで圧縮される
以下の処理を行うプログラムを作成せよ．*/

/*20. JSONデータの読み込み
Wikipedia記事のJSONファイルを読み込み，「イギリス」に関する記事本文を表示せよ．
問題21-29では，ここで抽出した記事本文に対して実行せよ．
*/

/* data/jawiki-country.json
ファイル形式
{"text": "~~~~~~~~~", "title": "~~~~~~~~~~~~~~"}
*/

/* ref.
http://qiita.com/oahiroaki/items/ff21d9adfb843161d1d8
http://blog.golang.org/json-and-go
*/

// 実行　$ go run src/test20.go data/jawiki-country.json

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TextTitle struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "not enough args\n")
		os.Exit(1)
	}

	//input data/hightemp.txt
	input_file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	file := bufio.NewReader(input_file)

	for {
		str, err := file.ReadString('\n')

		//ファイルの読み込み時に最後の行に来たらbreak
		if err == io.EOF {
			break
		}

		//各行をttにパース
		strByte := []byte(str)
		var tt TextTitle
		_ = json.Unmarshal(strByte, &tt)

		//Titleが"イギリス"の場合本文を出力
		if tt.Title == "イギリス" {
			fmt.Printf("%+v\n", tt.Text)
		}
	}

}
