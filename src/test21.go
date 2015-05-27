/*21. カテゴリ名を含む行を抽出
記事中でカテゴリ名を宣言している行を抽出せよ

[[Category:イギリス|*]]
[[Category:英連邦王国|*]]
[[Category:G8加盟国]]
[[Category:欧州連合加盟国]]
[[Category:海洋国家]]
[[Category:君主国]]
[[Category:島国|くれいとふりてん]]
[[Category:1801年に設立された州・地域]]
*/

/* ref.
http://qiita.com/hnakamur/items/a53b701c8827fe4bfec7
http://golang.org/pkg/bufio/
http://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/07.3.html
*/

// 実行 $ go run src/test20.go data/jawiki-country.json | go run src/test21.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for scanner.Scan() {
		text = scanner.Text()

		//正規表現でカテゴリ部分のみ出力
		m, _ := regexp.MatchString("\\[\\[Category:.*\\]\\]", text)
		if m {
			fmt.Println(text) // Println will add back the final '\n'
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
