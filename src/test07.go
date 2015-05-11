//07. テンプレートによる文生成
//引数x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．
//さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．

//ref. http://d.hatena.ne.jp/umisama/20100325/1269508269

//% go run src/test07.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scan(sent string) string {
	fmt.Printf("%v : ", sent)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func main() {
	//各x,y,zをターミナルから入力
	x := strings.Trim(scan("x"), "\n")
	y := strings.Trim(scan("y"), "\n")
	z := strings.Trim(scan("z"), "\n")

	fmt.Printf("「%v時の%vは%v」\n", x, y, z)
}
