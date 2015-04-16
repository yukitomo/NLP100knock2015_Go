//01. 「パタトクカシーー」
//「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．
// ref. http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37

package main

import "fmt" // 入出力フォーマットを実装したパッケージ

func main() {

	s := "パタトクカシーー"

	fmt.Printf("%v\n",s)
	runes := []rune(s)
    fmt.Println( string(runes[1]) + string(runes[3]) + string(runes[5]) + string(runes[7]))

}


