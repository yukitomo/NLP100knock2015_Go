/*16. ファイルをN分割する
自然数Nをコマンドライン引数などの手段で受け取り，入力のファイルを行単位でN分割せよ．
同様の処理をsplitコマンドで実現せよ．*/

//実行 % go run src/test16.go data/hightemp.txt [N] [output_dir]

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func outputFile(file *bufio.Reader) {
	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(s)
	}
}

func split_Nfile(file *bufio.Reader, N int, output_dir string) {
	lines := []string{}

	for {
		s, err := file.ReadString('\n')
		if err == io.EOF {
			break
		}
		lines = append(lines, s)
	}

	sum_num := len(lines)
	sub_num = len(lines) / N

	fmt.Printf("sum_num : %d\n", sum_num)
	fmt.Printf("sub_num : %d\n", sub_num)

	file_num := 0

	for i := 0; i < sum_num; i += sub_num {
		output_str := strings.Join(lines[i:i+sub_num], "")
		//fmt.Printf(string(i))
		//fmt.Printf(output_str)
		//fmt.Printf(output_dir + string(file_num))

		output_file, err := os.Create(output_dir + "/test16_" + strconv.Itoa(file_num) + ".txt")
		check(err)
		defer output_file.Close()

		output_file.WriteString(output_str)
		file_num += 1
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprint(os.Stderr, "not enough args\n")
		os.Exit(1)
	}

	//input data/hightemp.txt
	input_file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	//Nを引数としてとる
	var N int
	N, _ = strconv.Atoi(os.Args[2])

	output_dir := os.Args[3]

	split_Nfile(bufio.NewReader(input_file), N, output_dir)

}
