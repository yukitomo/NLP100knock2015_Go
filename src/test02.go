//02. 「パトカー」＋「タクシー」＝「パタトクカシーー」
//「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ

package main

import "fmt" // 入出力フォーマットを実装したパッケージ

func Altsum(s1, s2 string) string {
	runes1 := []rune(s1)  //rune配列に変換
	runes2 := []rune(s2)
	var s3 string
	for i := 0; i < len(runes1); i = i+1 {
		s3 += string(runes1[i]) + string(runes2[i])
	}
	return s3
}

func main() {

	s1 := "パトカー"
	s2 := "タクシー"

	fmt.Printf("%v\n",s1)
	fmt.Printf("%v\n",s2)
	fmt.Println(Altsum(s1,s2))
}