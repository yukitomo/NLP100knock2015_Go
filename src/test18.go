/*18. 各行を3コラム目の数値の降順にソート
各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）．
確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．*/

//実行 $ go run src/test18.go data/hightemp.txt
//確認 $ sort -r -k 3 data/hightemp.txt

/*ref
http://qiita.com/takuan_osho/items/74ef2d970a4888175336 slice to string
http://qiita.com/suin/items/0b2a815c815e23468adc sliceの要素を取り出す
http://matope.hatenablog.com/entry/2014/04/22/101127 文字列、数値変換
https://gist.github.com/kylelemons/1236125 mapのそーと
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ValSorter struct {
	Keys []string
	Vals []float64
}

func NewValSorter(m map[string]float64) *ValSorter {
	vs := &ValSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]float64, 0, len(m)),
	}
	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *ValSorter) Sort() {
	sort.Sort(vs)
}

func (vs *ValSorter) Len() int           { return len(vs.Vals) }
func (vs *ValSorter) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] } //降順なので逆
func (vs *ValSorter) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
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

	col013_col2 := make(map[string]float64)
	var col2 float64 //3カラム目

	//ファイルの各行を見ていく
	for {
		str, err := file.ReadString('\n')
		cols := strings.Split(str, "\t") //[]string 改行を消さないとfloat64にしても0になる
		//fmt.Printf("%v\n", cols[2])

		col013 := strings.Trim(str, "\n")                                        //カラム013を連結し一つの文字列
		col2, _ = strconv.ParseFloat(strings.Replace(cols[2], "\t", "", -1), 64) //string to float64

		//fmt.Printf("col2 : %v\n", col2)
		col013_col2[col013] = col2

		if err == io.EOF {
			break
		}

	}

	vs := NewValSorter(col013_col2)
	//fmt.Println(col013_col2)
	//fmt.Printf("%v\n", *vs)
	vs.Sort()
	//fmt.Printf("%v\n", *vs)

	for i := range vs.Keys {
		fmt.Printf("%v\n", vs.Keys[i])
	}

}
