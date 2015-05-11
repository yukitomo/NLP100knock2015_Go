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

// n がスライスに含まれているか
func member(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

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
		if err == io.EOF {
			break
		}
		if !member(str, uniq_strings) {
			uniq_strings = append(uniq_strings, str)
		}

	}

	sort.Strings(uniq_strings)

	for i := range uniq_strings {
		fmt.Printf("%v", uniq_strings[i])
	}

}
