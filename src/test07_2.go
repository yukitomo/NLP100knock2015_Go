//07. テンプレートによる文生成
//引数x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．
//さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．

//ref. http://d.hatena.ne.jp/umisama/20100325/1269508269

//$ go run src/test07_2.go x y z

package main

import (
	"flag"
	"fmt"
	"os"
)

func format(a []string) string {
	if len(a) < 3 {
		return ""
	}
	return fmt.Sprintf("%s時の%sは%s", a[0], a[1], a[2])
}

func main() {
	//引数に関してのUsage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s x y z\n", os.Args[0])
		flag.PrintDefaults()
	}

	//引数をパース
	flag.Parse()

	//引数の数を調べ、3つでない場合Usageを表示し、終了コードで終了
	if flag.NArg() != 3 {
		//Usageの表示
		flag.Usage()
		//終了コード1で終了
		os.Exit(1)
	}

	//flag.Args():標準入力の引数がstringの配列で渡される
	fmt.Println(format(flag.Args()))
}
