package cipher

import (
	"conversion/pkg"
	"math"
	"math/big"
	"math/rand"
	"time"
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	D *big.Int
	N *big.Int
}

func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int64) int64 {
	return (a * b) / GCD(a, b)
}

func getCoprime(x int64) int64 {
	var res int64
	for i := int64(2); i < x; i++ {
		if GCD(x, i) == 1 {
			res = i
			break
		}
	}
	return res
}

func GenRsaKey() (*PublicKey, *PrivateKey) {
	prime_list := GenPrime()
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(prime_list))
	j := rand.Intn(len(prime_list))
	p := prime_list[i]
	q := prime_list[j]

	N := big.NewInt(p * q)
	s := big.NewInt(LCM(p-1, q-1))         // p-1とq-1の最小公倍数s
	E := big.NewInt(getCoprime(s.Int64())) // sと互いに素な数e
	D := big.NewInt(0)
	D.ModInverse(E, s)                         // d = e^-1 mod s
	return &PublicKey{N, E}, &PrivateKey{D, N} // 公開鍵、秘密鍵
}

func Encrypt(pk *PublicKey, m *big.Int) *big.Int {
	c := big.NewInt(0)
	m.Mod(m, pk.N)              // Nを超えないようにする
	return c.Exp(m, pk.E, pk.N) // c = m ^ e mod n
}

func Decrypt(sk *PrivateKey, c *big.Int) *big.Int {
	m := big.NewInt(0)
	c.Mod(c, sk.N)              // Nを超えないようにする
	return m.Exp(c, sk.D, sk.N) // m = c ^ d mod n
}

func RsaEnc(pk *PublicKey, s string) []big.Int {
	ss := []rune(s)
	var ress []big.Int = pkg.Map(ss, func(r rune) big.Int {
		c := big.NewInt(int64(r))
		return *Encrypt(pk, c)
	})
	return ress
}

func Remove(s_list []int64, index int64) (tmp []int64) {
	tmp = append(s_list[:index], s_list[(index+1):]...)
	return
}

func GenPrime() []int64 {
	number := int64(math.Pow(2, 20)) // 初期化:number=2^20で固定
	prime_list := []int64{}
	search_list := []int64{}
	for i := int64(2); i < number; i++ {
		search_list = append(search_list, i) // 2からnumber-1までの数字の配列を作る
	}
	limit := int64(math.Sqrt(float64(number))) // 探索リストの先頭値が√numberを超えたら終了
	for search_list[0] <= limit {
		p_num := search_list[0] // 探索リストの先頭を素数リストに移動
		prime_list = append(prime_list, p_num)
		search_list = Remove(search_list, 0)   // 探索リストの先頭を削除
		search_list_length := len(search_list) // p_numの倍数を探索リストから篩落とす
		tmp := []int64{}
		for i := 0; i < search_list_length; i++ {
			if search_list[i]%p_num != 0 {
				tmp = append(tmp, search_list[i])
			}
		}
		search_list = tmp
	}
	prime_list = append(prime_list, search_list...) // 探索リストの残りを素数リストに追加
	return prime_list
}

/*
func main() {
	pk, sk := GenKey()

	m := big.NewInt(0)
	cipherText := Encrypt(pk, m)
	result := Decrypt(sk, cipherText)
	fmt.Println(result)
}
*/
