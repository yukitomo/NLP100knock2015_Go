//00. 文字列の逆順
//文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
//% go run test00.go

//ref. http://d.hatena.ne.jp/yuheiomori0718/20131008/1381244919

package main

import "fmt" // 入出力フォーマットを実装したパッケージ

func Reverse(s string) string {
	runes := []rune(s)  //rune配列に変換
	for i, j := 0, len(runes)-1; i<j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	s := "stressed"

	fmt.Printf("%v\n",s)
    fmt.Println(Reverse(s))

}