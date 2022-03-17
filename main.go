package main

import (
	"conversion/pkg/cipher"
	"fmt"
)

func main() {
	//動作テスト用
	//var v string = "abcdefghijklmn"
	//fmt.Println(cipher.Ceaser(v))

	v := "abcdefghijklmnop"
	fmt.Println(cipher.Transpose(v, 5))

	//n := cipher.GenBunrhamKey()
	//fmt.Println(n)
	//println(cipher.Burnham("Hello,world!", n))
}
