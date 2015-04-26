/*15. 末尾のN行を出力
自然数Nをコマンドライン引数などの手段で受け取り，
入力のうち末尾のN行だけを表示せよ．確認にはtailコマンドを用いよ．
*/

/* ref.
http://matope.hatenablog.com/entry/2014/04/22/101127
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	//"strings"
)

func outputFile(file *bufio.Reader) {
	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(s)
	}
}

func tail_Nline(file *bufio.Reader, N int) {
	lines := []string{}

	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		lines = append(lines, s)
	}

	sumnumber := len(lines)

	for i := 0; i < N; i++ {
		fmt.Print(lines[sumnumber-N+i])
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "not enough args\n")
		os.Exit(1)
	}

	//input data/hightemp.txt
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	//Nを引数としてとる
	var N int
	N, _ = strconv.Atoi(os.Args[2])
	tail_Nline(bufio.NewReader(file), N)

}
