/*24. ファイル参照の抽出
記事から参照されているメディアファイルをすべて抜き出せ．*/

//[[File:Oil platform in the North SeaPros.jpg|thumb|[[北海油田]]]]
//[[ファイル:Durham Kathedrale Nahaufnahme.jpg|[[ダラム大聖堂]]

// 実行 $ go run src/test20.go data/jawiki-country.json | go run src/test24.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strings"
)

func main() {

	re1, _ := regexp.Compile("^\\[*ファイル:|^\\[*File:")
	re2, _ := regexp.Compile("\\|.*$")

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	//var count int
	for scanner.Scan() {
		text = scanner.Text()

		//正規表現でカテゴリ部分のみ出力
		m, _ := regexp.MatchString("^\\[*ファイル:.*\\|.*$|^\\[*File:.*\\|.*$", text)
		if m {
			text = re1.ReplaceAllString(text, "")
			text = re2.ReplaceAllString(text, "")
			fmt.Printf("%v\n", text)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
