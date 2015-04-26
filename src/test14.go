/*14. 先頭からN行を出力
自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．
確認にはheadコマンドを用いよ．
*/

/* ref.
http://matope.hatenablog.com/entry/2014/04/22/101127
*/

//実行 % go run src/test14.go data/hightemp.txt 3
//確認 % head -n 3 data/hightemp.txt

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func head_Nline(file *bufio.Reader, N int) {
	n := 0
	for {
		s, _ := file.ReadString('\n')
		if n == N {
			break
		}
		n += 1
		fmt.Print(s)
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
	head_Nline(bufio.NewReader(file), N)

}
