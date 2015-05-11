/*08. 暗号文
与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．

英小文字ならば(219 - 文字コード)の文字に置換
その他の文字はそのまま出力
この関数を用い，英語のメッセージを暗号化・復号化せよ.
*/

// ref. 正規表現 http://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/07.3.html
//http://www.utf8-chartable.de/unicode-utf8-table.pl?number=128&utf8=dec

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func scan(sent string) string {
	fmt.Printf("%v : ", sent)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

//英小文字かどうか正規表現で判定
func IsLowAlphabet(character string) bool {
	if m, _ := regexp.MatchString("[a-z]", character); !m {
		return false
	}
	return true
}

//英小文字ならば(219 - 文字コード)の文字に置換、その他の文字はそのまま出力
func Cipher(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {

		//正規表現で判断（runeをいちいちstringに戻しているので冗長）
		/*if IsLowAlphabet(string(runes[i])) {
			runes[i] = 219 - runes[i]
		}*/

		//英小文字 [a-z] はutf-8の10進数で 97 ~ 123 で表現されるのでそれを用いて判断
		if runes[i] > 96 && runes[i] < 123 {
			runes[i] = 219 - runes[i]
		}
	}
	return string(runes)
}

func main() {
	sent := strings.Trim(scan("input sentence"), "\n")
	cipherSent := Cipher(sent)
	fmt.Printf("cipherted sentence : %v\n", cipherSent)
	decodeSent := Cipher(cipherSent)
	fmt.Printf("decoded sentence : %v\n", decodeSent)
}
