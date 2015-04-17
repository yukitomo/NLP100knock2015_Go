/*
09. Typoglycemia
スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，
それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．
ただし，長さが４以下の単語は並び替えないこととする．
適当な英語の文（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）
を与え，その実行結果を確認せよ．
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
func scan(sent string) string {
	fmt.Printf("%v : ", sent)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}
*/

func WordsSplit(s string) []string {
	words := strings.Split(s, " ") //[]に単語が格納
	return words
}

func RamdomizeWord(word string) string{
	var new_word string
	runes := []rune(word)
	length := len(runes)

	if length =< 4 {

	}else{
		new_word = runes
	}
	return 
}

func RamdomizeSent(sent string) string{
	words := WordsSplit(sent)

	for i := 0; i<len(words); i++{
		words[i] = RamdomizeWord(words[i])
	}
	return words 
}


func main() {
	//sent := strings.Trim(scan("input sentence"), "\n")
	sent := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	fmt.Printf("%v\n", s)

	cipherSent := Cipher(sent)
	fmt.Printf("cipherted sentence : %v\n", cipherSent)
	decodeSent := Cipher(cipherSent)
	fmt.Printf("decoded sentence : %v\n", decodeSent)
}
