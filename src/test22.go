/*22. カテゴリ名の抽出
記事のカテゴリ名を（行単位ではなく名前で）抽出せよ．*/

// 実行 $ go run src/test20.go data/jawiki-country.json | go run src/test22.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	re1, _ := regexp.Compile("\\[\\[Category:")
	re2, _ := regexp.Compile("\\]\\]")

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for scanner.Scan() {
		text = scanner.Text()

		//正規表現でカテゴリ部分のみ出力
		m, _ := regexp.MatchString("\\[\\[Category:.*\\]\\]", text)
		if m {
			//タグの削除
			text = re1.ReplaceAllString(text, "")
			text = re2.ReplaceAllString(text, "")

			fmt.Println(text)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
