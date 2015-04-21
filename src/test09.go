/*
09. Typoglycemia
スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，
それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．
ただし，長さが４以下の単語は並び替えないこととする．
適当な英語の文（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）
を与え，その実行結果を確認せよ．
*/

/*ref.
https://groups.google.com/forum/#!topic/golang-nuts/JDVCR__dHbI
http://golang.org/pkg/math/rand/
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func WordsSplit(s string) []string {
	words := strings.Split(s, " ") //[]に単語が格納
	return words
}

func RandomizeWord(word string) string {
	var new_word string
	runes := []rune(word)
	length := len(runes)

	//文字列長が4以下の場合
	/*
		string = "ABCDE"
		length = 5
		index = [0,1,2,3,4]
		rand_idx := rand.Perm(length - 2) = [0,1,2] → 1を足せば [1,2,3] が取れる
	*/

	if length >= 4 {
		new_runes := make([]rune, length)
		rand_idx := rand.Perm(length - 2) //文字列のインデックスを並び替えたもの
		//最初と最後は同じ
		new_runes[0] = runes[0]
		new_runes[length-1] = runes[length-1]
		//
		for i := range rand_idx {
			new_runes[i+1] = runes[rand_idx[i]+1]
		}
		new_word = string(new_runes)
	} else {
		new_word = word
	}
	return new_word
}

func RandomizeSent(sent string) string {
	words := WordsSplit(sent)

	for i := 0; i < len(words); i++ {
		words[i] = RandomizeWord(words[i])
	}
	return strings.Join(words, " ")
}

func main() {
	//sent := strings.Trim(scan("input sentence"), "\n")
	sent := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	random_sent := RandomizeSent(sent)
	fmt.Printf("src  sent : %v\n", sent)
	fmt.Printf("rand sent : %v\n", random_sent)
}
