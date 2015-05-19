/*19. 各行の1コラム目の文字列の出現頻度を求め，出現頻度の高い順に並べる
各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
確認にはcut, uniq, sortコマンドを用いよ．*/

//実行 $ go run src/test19.go data/hightemp.txt

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type ValSorter struct {
	Keys []string
	Vals []int
}

func NewValSorter(m map[string]int) *ValSorter {
	vs := &ValSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int, 0, len(m)),
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

	col1Counter := make(map[string]int)

	//ファイルの各行を見ていく
	for {
		str, err := file.ReadString('\n')
		cols := strings.Split(str, "\t") //[]string 改行を消さないとfloat64にしても0になる

		col1Counter[cols[0]] += 1

		if err == io.EOF {
			break
		}

	}

	vs := NewValSorter(col1Counter)
	//fmt.Println(col013_col2)
	//fmt.Printf("%v\n", *vs)
	vs.Sort()
	//fmt.Printf("%v\n", *vs)

	for i := range vs.Keys {
		fmt.Printf("%v %v\n", vs.Keys[i], vs.Vals[i])
	}
}
