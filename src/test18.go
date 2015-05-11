/*18. 各行を3コラム目の数値の降順にソート
各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）．
確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．*/

/*17. １列目の文字列の異なり
1列目の文字列の種類（異なる文字列の集合）を求めよ．確認にはsort, uniqコマンドを用いよ．
*/

//実行 $ go run src/test17.go data/hightemp.txt
//確認 $ sort data/col1.txt|uniq

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "not enough args\n")
		os.Exit(1)
	}

	//input data/hightemp.txt
	input_file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	file := bufio.NewReader(input_file)

	uniq_strings := []string{}

	for {
		str, err := file.ReadString('\n')
		cols := strings.Split(str, "\t")
		if err == io.EOF {
			break
		}

	}

}
