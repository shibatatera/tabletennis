package main

import (
	"fmt"
	"strconv"
)

//数値を一桁ずつにする方法 gemini

func strRune(n int) []int {
	num := n
	s := strconv.Itoa(num) // 数値を文字列に変換

	var digits []int
	for _, r := range s {
		// 各文字を数値に変換
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			// エラーハンドリング（ここでは単純にスキップまたはログ出力など）
			fmt.Printf("Error converting character '%c' to int: %v\n", r, err)
			continue
		}
		digits = append(digits, digit)
	}

	return digits
}

func main() {
	digits := strRune(30219)
	for i := 0; i < len(digits); i++ {
		fmt.Printf("rune Unicode: %v\n", rune('0'+digits[i]))
	}

	// num := 12345
	// s := strconv.Itoa(num) // 数値を文字列に変換
	//
	// var digits []int
	// for _, r := range s {
	// // 各文字を数値に変換
	// digit, err := strconv.Atoi(string(r))
	// if err != nil {
	// // エラーハンドリング（ここでは単純にスキップまたはログ出力など）
	// fmt.Printf("Error converting character '%c' to int: %v\n", r, err)
	// continue
	// }
	// digits = append(digits, digit)
	// }
	//
	// fmt.Printf("元の数値: %d\n", num)
	// fmt.Printf("一桁ずつのスライス: %v\n", digits)

	// または、rune から直接数値に変換（ASCIIコード利用）
	// var digits2 []int
	// for _, r := range s {
	// digit := int(r - '0') // runeから'0'のASCII値を引く
	// digits2 = append(digits2, digit)
	// }
	// fmt.Printf("一桁ずつのスライス (rune-'0'): %v\n", digits2)

	//fmt.Printf("rune: %v\n", rune('0'+digits[0]))

}
