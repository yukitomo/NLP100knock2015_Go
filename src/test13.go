/*13. col1.txtとcol2.txtをマージ
12で作ったcol1.txtとcol2.txtを結合し，
元のファイルの1列目と2列目をタブ区切りで並べたテキストファイルを作成せよ．
確認にはpasteコマンドを用いよ．
*/

//ref. http://www.geocities.jp/m_hiroi/golang/abcgo11.html

//実行 % go run src/test13.go data/col1.txt data/col2.txt
//確認 $ paste data/col1.txt data/col2.txt

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func join(file1 *bufio.Reader, file2 *bufio.Reader) {
	for {
		s1, err1 := file1.ReadString('\n')
		if err1 == io.EOF {
			outputFile(file2)
			break
		}
		s2, err := file2.ReadString('\n')
		if err == io.EOF {
			fmt.Print(s1)
			outputFile(file1)
			break
		}
		//s1の改行をtabに変更
		fmt.Print(strings.Replace(s1, "\n", "\t", -1) + s2)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "not enough args\n")
		os.Exit(1)
	}
	file1, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Fprint(os.Stderr, err1)
		os.Exit(1)
	}
	file2, err2 := os.Open(os.Args[2])
	if err2 != nil {
		fmt.Fprint(os.Stderr, err2)
		os.Exit(1)
	}
	join(bufio.NewReader(file1), bufio.NewReader(file2))
	file1.Close()
	file2.Close()
}
