# conversion
**Mercari spring intern Gophers 2022**

ジェネリクスを使った関数を実装しているので、このレポジトリは go 1.18 以降を対象としています。

OSS作りの練習として作成したものです。
小さいサイズの文字列を手元で暗号化したいときにでも使ってください。

## 実装している暗号
`import "conversion/pkg"`していれば`conversion.pkg.cipher`は`cipher`となります。
### シーザー暗号（古典暗号）
`Ceaser(m string) []string`
```
// 引数にアルファベットの平文を入力するとROT1からROT13までの暗号文を生成します。
var c []string = conversion.pkg.cipher.Ceaser("abc")
fmt.Println(c)
// [bcd cde def efg fgh ghi hij ijk jkl klm lmn mno nop]
var d []string = conversion.pkg.cipher.Ceaser(c[12])
// [opq pqr qrs rst stu tuv uvw vwx wxy xyz yza zab abc]
```

### 転置式暗号（古典暗号）
`Transpose(m string, n int) string`
```
// 引数に平文と転置を適用する文字ブロックのサイズを渡すと暗号文を生成します。
var c string = conversion.pkg.cipher.Transpose("abcdefghijklmnop", 5)
fmt.Println(c)
// e.g. bedacgjifhlonkmp
fmt.Println(conversion.pkg.cipher.Transpose(c, 5))
// e.g. bdecagijhflnomkp
```

### バーナム暗号(共通鍵暗号)
`GenBurnhamKey() int // 鍵生成`
`Burnham(m string, key int) // 暗号化&復号化`
```
// 共通鍵を生成します。
key := conversion.pkg.cipher.GenBurnhamKey()
// e.g. key: 52929
// 引数に平文と共通鍵を渡すと暗号文を生成します。
var c string = conversion.pkg.cipher.Burnham("京都", key)
fmt.Println(c)
// e.g 聳帢
// 引数に暗号文と共通鍵を渡すと平文を生成します。
fmt.Println(conversion.pkg.cipher.Burnham(c, key))
// e.g. 京都
```

### RSA暗号(公開鍵暗号)
公開鍵と秘密鍵は以下の構造体です。
```
type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	D *big.Int
	N *big.Int
}
```
`GenRsaKey() (*Publickey, *PrivateKey) // 鍵生成`

`RsaEnc(pk *Publickey, m string) []big.Int // 暗号化`
```
// 公開鍵と秘密鍵を生成します。
pk, sk := conversion.pkg.cipher.GenRsaKey()
fmt.Println(*pk)
// e.g. {76571165813 7}
fmt.Println(*sk)
// e.g. {10938615223 76571165813}
// 平文と公開鍵で暗号文を生成します。
fmt.Println(conversion.pkg.cipher.RsaEnc(pk, m))
```
