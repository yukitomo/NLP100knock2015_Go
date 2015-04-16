//05. n-gram
//与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．
//この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．

// ref. http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37
// http://qiita.com/kazuph/items/4bc33162da7e7d00d80c

package main

import (
	"fmt"     // 入出力フォーマットを実装したパッケージ
	"strings" //string処理のメソッド
)

// n がスライスに含まれているか
func member(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

// 重複要素を取り除く
func RemoveDuplicates(xs []string) []string {
	ys := make([]string, 0, len(xs))
	for _, x := range xs {
		if !member(x, ys) {
			ys = append(ys, x)
		}
	}
	return ys
}

func WordsSplit(s string) []string {
	s2 := strings.Trim(s, ".") // "."の除去

	//s3 := strings.Trim(s2, ",") Trimは文字列の前後しか除去できない
	//func Replace(s, old, new string, n int) string
	//Replaceは、文字列sのコピーに対し、最初のn個、oldの部分をnewに置換(重複はなし)したものを返します。n < 0のとき、置換数は無制限になります。
	s3 := strings.Replace(s2, ",", "", -1)

	words := strings.Split(s3, " ") //[]に単語が格納
	return RemoveDuplicates(words)
}

//与えられたslice(単語 or 文字)からngramを返す
func MakeNgram(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; i < (len(objs) - n + 1); i++ {
		//fmt.Println((objs[i : i+n]))
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
	return RemoveDuplicates(Ngrams)
}

func MakeCharactersNgram(words []string, n int) []string {
	ChNgrams := []string{}
	for i := 0; i < len(words); i++ {
		characters := strings.Split(words[i], "")
		ChNgrams = append(ChNgrams, MakeNgram(characters, n, "")...)
	}
	return ChNgrams
}

func PrintSlice(objs []string) {
	for i := 0; i < len(objs); i++ {
		fmt.Printf("%v\n", objs[i])
	}
}

func MakeUnion(s1, s2 []string) []string {
	set := append(s1, s2...)
	return RemoveDuplicates(set)
}

func main() {
	s := "I am an NLPer"
	fmt.Printf("%v\n", s)

	words := WordsSplit(s) //単語の分割
	//fmt.Println(words[1:3])

	//単語bigram
	fmt.Printf("---words bigram---\n")
	Wordsbigram := MakeNgram(words, 2, " ")
	fmt.Printf("%v\n", Wordsbigram)
	PrintSlice(Wordsbigram)

	//文字bigram
	fmt.Printf("---character bigram---\n")
	fmt.Printf("%v\n", words)
	Charactersbigram := MakeCharactersNgram(words, 2)
	PrintSlice(Charactersbigram)
	//fmt.Printf("%v\n", Charactersbigram)
}
