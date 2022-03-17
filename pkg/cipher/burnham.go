package cipher

import (
	"conversion/pkg"
	"math"
	"math/rand"
	"time"
)

func GenBunrhamKey() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(math.Pow(2, 16)))
}

func Burnham(s string, n int) string {
	ss := []rune(s)
	var ress []int = pkg.Map(ss, func(r rune) int {
		return int(r) ^ n
	})
	res := ""
	for i := 0; i < len(ress); i++ {
		res = res + string(rune(ress[i]))
	}
	return res
}
