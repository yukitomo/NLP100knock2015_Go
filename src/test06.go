//06. 集合
//"paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ,
//XとYとして求め，XとYの和集合，積集合，差集合を求めよ．
//さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．

//ref. http://qiita.com/megu_ma/items/f5e66293b2934e117731
// http://www.geocities.jp/m_hiroi/golang/yagp02.html
// http://golang.jp/effective_go

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

func MakeIntersection(s1, s2 []string) []string {
	set := []string{}
	for _, c := range s1 {
		if member(c, s2) {
			set = append(set, c)
		}
	}
	return set
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"
	//fmt.Printf("%v\n", s1)
	//fmt.Printf("%v\n", s2)

	X := MakeNgram(strings.Split(s1, ""), 2, "")
	Y := MakeNgram(strings.Split(s2, ""), 2, "")
	fmt.Printf("X : %v\n", X)
	fmt.Printf("Y : %v\n", Y)
	//PrintSlice(X)
	//PrintSlice(Y)

	Union := MakeUnion(X, Y)
	fmt.Printf("Union(X,Y) : %v\n", Union)

	Intersection := MakeIntersection(X, Y)
	fmt.Printf("Intersection(X,Y) : %v\n", Intersection)

}
