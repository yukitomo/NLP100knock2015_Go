/*25. テンプレートの抽出
記事中に含まれる「基礎情報」テンプレートのフィールド名と値を抽出し，
辞書オブジェクトとして格納せよ．*/

//ref. http://golang.jp/effective_go

/*
{{基礎情報 国
|略名 = イギリス
|日本語国名 = グレートブリテン及び北アイルランド連合王国
|公式国名 = {{lang|en|United Kingdom of Great Britain and Northern Ireland}}<ref>英語以外での正式国名:<br/>
*{{lang|gd|An Rìoghachd Aonaichte na Breatainn Mhòr agus Eirinn mu Thuath}}（[[スコットランド・ゲール語]]）<br/>
*{{lang|cy|Teyrnas Gyfunol Prydain Fawr a Gogledd Iwerddon}}（[[ウェールズ語]]）<br/>
*/

// 実行 $ go run src/test20.go data/jawiki-country.json | go run src/test25.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strings"
)

func main() {

	//各値の初期化
	var text string
	var name string
	var value string
	//辞書の定義
	dict := make(map[string]string)

	infoStart := false
	infoEnd := false

	re1, _ := regexp.Compile("^\\||=\\s.*$")
	re2, _ := regexp.Compile("^\\|.* = ")

	//ファイルの読み込み
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()

		//テキスト処理

		//infoEnd = true のときスキャンを止める
		if infoEnd {
			break
		}

		//"{{基礎情報" 以下に入った場合
		if infoStart {
			//基礎情報」テンプレートかどうか
			isTemp, _ := regexp.MatchString("^\\|.* = .*$", text)
			infoEnd, _ = regexp.MatchString("^\\}\\}$", text)

			if isTemp {
				//name, value の抽出
				name = re1.ReplaceAllString(text, "")  //フィールド名の抽出
				value = re2.ReplaceAllString(text, "") //値の抽出

				//辞書への格納
				dict[name] = value
				//fmt.Printf("%v = %v\n", name, value)

			} else {
				// "|公式国名 =" のフォーマットになってないものは直前の value に足しこむ
				dict[name] = dict[name] + text
			}

		} else {
			// "{{基礎情報" か判断
			infoStart, _ = regexp.MatchString("^\\{\\{基礎情報 .*$", text)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//格納した辞書の出力
	for name, value := range dict {
		fmt.Printf("%v = %v\n", name, value)
	}

}
