/*23. セクション構造
記事中に含まれるセクション名
とそのレベル（例えば"== セクション名 =="なら1）を表示せよ．*/

// 実行 $ go run src/test20.go data/jawiki-country.json | go run src/test23.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	re1, _ := regexp.Compile("^==+")
	re2, _ := regexp.Compile("==+$")

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	var count int
	for scanner.Scan() {
		text = scanner.Text()

		//正規表現でカテゴリ部分のみ出力
		m, _ := regexp.MatchString("^==+.*==+$", text)
		if m {
			count = strings.Count(text, "=")
			text = re1.ReplaceAllString(text, "")
			text = re2.ReplaceAllString(text, "")
			fmt.Printf("%v : %v\n", text, count/2-1)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
