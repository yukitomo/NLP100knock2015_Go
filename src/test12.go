/*
12. 1列目をcol1.txtに，2列目をcol2.txtに保存
各行の1列目だけを抜き出したものをcol1.txtに，
2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．
確認にはcutコマンドを用いよ．
*/

//実行 $ go run src/test12.go < data/hightemp.txt

package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//loading input_file
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	//create output_file
	f1, err1 := os.Create("data/col1.txt")
	check(err1)
	defer f1.Close()

	f2, err2 := os.Create("data/col2.txt")
	check(err2)
	defer f2.Close()

	for scanner.Scan() {
		str := scanner.Text()
		cols := strings.Split(str, "\t")
		//fmt.Println(cols[0])
		//fmt.Println(cols[1])
		f1.WriteString(cols[0] + "\n")
		f2.WriteString(cols[1] + "\n")
	}

}
