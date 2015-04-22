/*
11. タブをスペースに置換
タブ1文字につきスペース1文字に置換せよ．
確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．
*/
//ref. http://kobegdg.blogspot.jp/2013/04/go.html
//     http://mattintosh.hatenablog.com/entry/2013/01/16/143323

//実行 $ go run src/test11.go < data/hightemp.txt
//確認 $ expand -t 1 data/hightemp.txt
//　　 $ tr '\t' " " < data/hightemp.txt
//　　 $ sed s/$'\t'/$' '/g data/hightemp.txt

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		resStr := strings.Replace(str, "\t", " ", -1)
		fmt.Println(resStr)
	}

}
