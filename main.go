package main

import (
	"conversion/pkg/cipher"
	"fmt"
)

func main() {
	//動作テスト用
	//var v []string = cipher.Ceaser([]string{"abc", "def", "ghi", "jkl"}, 4)
	//fmt.Println('あ')

	var v string = "abcdefghijklmn"
	fmt.Println(cipher.Ceaser(v))

	//v := "abcdefghijklmnop"
	//fmt.Println(cipher.Transpose(v, 5))
}
