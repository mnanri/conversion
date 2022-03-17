package cipher

import (
	"conversion/pkg"
	"math/rand"
	"strings"
	"time"
)

func Transpose(s string, n int) string {
	t := strings.Split(s, "")
	seg := ""
	var ss []string = []string{}
	for i := 1; i <= len(s); i++ {
		seg = seg + t[i-1]
		if i%5 == 0 {
			ss = append(ss, seg)
			seg = ""
		}
	}
	if seg != "" {
		ss = append(ss, seg)
	}

	var order []int = []int{}
	for i := 0; i < n; i++ {
		order = append(order, i)
	}
	rand.Seed(time.Now().UnixNano())
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		order[i], order[j] = order[j], order[i]
	}

	var ress []string = pkg.Map(ss, func(u string) string {
		tmp := ""
		if len(u) == n {
			var a []string = strings.Split(u, "")
			for i := 0; i < n; i++ {
				tmp = tmp + a[order[i]]
			}
		} else {
			tmp = u
		}
		return tmp
	})

	res := ""
	for i := 0; i < len(ress); i++ {
		res = res + ress[i]
	}
	return res
}
