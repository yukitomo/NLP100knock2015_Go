//04. 元素記号
//"Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
//という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．

//実行 $ go run src/test04.go

package main

import (
	"fmt"     // 入出力フォーマットを実装したパッケージ
	"strings" //string処理のメソッド
)

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	s = strings.Trim(s, ".") // "."の除去

	s = strings.Replace(s, ",", "", -1)
	words := strings.Split(s, " ") //slicesに単語が格納

	words_index := make(map[string]int)

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		switch i {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19: //1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字
			words_index[string(runes[0])] = i
		default: //上記以外の単語は先頭の2文字
			words_index[string(runes[0])+string(runes[1])] = i
		}
	}
	fmt.Println(words_index)
}
