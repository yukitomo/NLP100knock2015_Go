//03. 円周率
//"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
//という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．

//ref. https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/07.6.md

package main

import (
	"fmt"     // 入出力フォーマットを実装したパッケージ
	"strings" //string処理のメソッド
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	//fmt.Printf("%v\n", s)

	s2 := strings.Trim(s, ".") // "."の除去
	//fmt.Printf("%v\n", s2)

	//s3 := strings.Trim(s2, ",") Trimは文字列の前後しか除去できない
	//func Replace(s, old, new string, n int) string
	//Replaceは、文字列sのコピーに対し、最初のn個、oldの部分をnewに置換(重複はなし)したものを返します。n < 0のとき、置換数は無制限になります。
	s3 := strings.Replace(s2, ",", "", -1)
	//fmt.Printf("%v\n", s3)

	words := strings.Split(s3, " ") //slicesに単語が格納
	//fmt.Printf("%v\n", words)

	for i := 0; i < len(words); i++ {
		fmt.Printf("%d\n", len(words[i]))
	}

}
