package main

import (
	"conversion/pkg/cipher"
	"fmt"
)

func main() {
	//動作テスト用
	//var v []string = cipher.Ceaser([]string{"abc", "def", "ghi", "jkl"}, 4)

	var v []string = []string{"abc", "lmn", "xyz", "ABC"}
	for i := 0; i < 26; i++ {
		fmt.Println(cipher.Ceaser(v, i))
	}
}
