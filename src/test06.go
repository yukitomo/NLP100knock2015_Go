//06. 集合
//"paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ,
//XとYとして求め，XとYの和集合，積集合，差集合を求めよ．
//さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．

//ref. http://qiita.com/megu_ma/items/f5e66293b2934e117731
// http://www.geocities.jp/m_hiroi/golang/yagp02.html
// http://golang.jp/effective_go

//実行 $ go run src/test06.go

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

//与えられたslice(単語 or 文字)からngramを返す
func MakeNgram(objs []string, n int, c string) []string {
	Ngrams := []string{}
	for i := 0; i < (len(objs) - n + 1); i++ {
		//fmt.Println((objs[i : i+n]))
		Ngrams = append(Ngrams, strings.Join(objs[i:i+n], c))
	}
	return RemoveDuplicates(Ngrams)
}

//和集合の導出
func MakeUnion(set1, set2 []string) []string {
	set := append(set1, set2...) //sliceの結合
	return RemoveDuplicates(set) //重複要素の削除
}

//積集合の導出
func MakeIntersection(set1, set2 []string) []string {
	set := []string{}
	//set1の各要素がset2に含まれていた場合setに格納
	for _, c := range set1 {
		if member(c, set2) {
			set = append(set, c)
		}
	}
	return set
}

//差集合の導出 (set1 - set2)
func MakeDifference(set1, set2 []string) []string {
	set := []string{}
	//set1の各要素がset2に含まれていない場合setに格納
	for _, c := range set1 {
		if !member(c, set2) {
			set = append(set, c)
		}
	}
	return set
}

func main() {
	sent1 := "paraparaparadise"
	sent2 := "paragraph"
	fmt.Printf("%v\n", sent1)
	fmt.Printf("%v\n", sent2)

	X := MakeNgram(strings.Split(sent1, ""), 2, "") //文字bigramの作成
	Y := MakeNgram(strings.Split(sent2, ""), 2, "")
	fmt.Printf("X : %v\n", X)
	fmt.Printf("Y : %v\n", Y)

	Union := MakeUnion(X, Y) //和集合
	fmt.Printf("Union(X,Y) : %v\n", Union)

	Intersection := MakeIntersection(X, Y) //積集合
	fmt.Printf("Intersection(X,Y) : %v\n", Intersection)

	Difference := MakeDifference(X, Y) //差集合
	fmt.Printf("Difference(X,Y) : %v\n", Difference)

	//'se'がXに含まれるか
	if member("se", X) {
		fmt.Printf("'se' is member of X\n")
	} else {
		fmt.Printf("'se' isn't member of X\n")
	}

	//'se'がYに含まれるか
	if member("se", Y) {
		fmt.Printf("'se' is member of Y\n")
	} else {
		fmt.Printf("'se' isn't member of Y\n")
	}

}
