/*
10. 行数のカウント
行数をカウントせよ．確認にはwcコマンドを用いよ．
*/
// ref . http://stackoverflow.com/questions/18002834/file-input-in-golang

// 実行 $ go run src/test10.go < data/hightemp.txt
// 確認 $ wc data/hightemp.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count += 1
	}
	fmt.Println(count)

}
