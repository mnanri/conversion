package main

import (
	"conversion/pkg/cipher"
	"fmt"
)

func main() {
	//動作テスト用
	//var v string = "abcdefghijklmn"
	//fmt.Println(cipher.Ceaser(v))

	//v := "abcdefghijklmnop"
	//fmt.Println(cipher.Transpose(v, 5))

	//n := cipher.GenBunrhamKey()
	//fmt.Println(n)
	//fmt.Println(cipher.Burnham("Hello,world!", n))
	//fmt.Println(cipher.Burnham(cipher.Burnham("Hello,world!", n), n))

	s := "abc"
	pk, sk := cipher.GenRsaKey()
	fmt.Println(*pk)
	fmt.Println(*sk)
	fmt.Println(cipher.RsaEnc(pk, s))
}
