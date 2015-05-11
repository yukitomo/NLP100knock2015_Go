/*16. ファイルをN分割する
自然数Nをコマンドライン引数などの手段で受け取り，入力のファイルを行単位でN分割せよ．
同様の処理をsplitコマンドで実現せよ．*/

//実行 % go run src/test16.go data/hightemp.txt [N] [output_dir]
//確認 % bash knock016.sh ../Data/hightemp.txt [outdir] [n]

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
	sub_num := sum_num / N
	rest_num := sum_num % N
	set_nums := []int{}

	for i := 0; i < N; i++ {
		if rest_num > 0 {
			set_nums = append(set_nums, sub_num+1)
			rest_num -= 1
		} else {
			set_nums = append(set_nums, sub_num)
		}
	}

	//fmt.Printf("sum_num : %d\n", sum_num)
	//fmt.Printf("sub_num : %d\n", sub_num)
	//fmt.Println(set_nums)

	file_num := 0
	pre_i := 0
	var nex_i int
	for i := 0; i < N; i++ {
		nex_i = pre_i + set_nums[i]
		output_str := strings.Join(lines[pre_i:nex_i], "")

		output_file, err := os.Create(output_dir + "/test16_" + strconv.Itoa(file_num) + ".txt")
		check(err)
		defer output_file.Close()

		output_file.WriteString(output_str)
		file_num += 1
		pre_i = nex_i
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
