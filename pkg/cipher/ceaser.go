package cipher

func Map[X, Y any](xs []X, f func(x X, n int) Y) []Y {
	ys := make([]Y, len(xs))
	for i := range xs {
		ys[i] = f(xs[i], i)
	}
	return ys
}

func Ceaser(s string) []string {
	var ss []string = []string{}
	for i := 0; i < 26; i++ {
		ss = append(ss, s)
	}
	var res []string = Map(ss, func(t string, r int) string {
		tmp := ""
		for j := range t {
			switch {
			case ('A' <= t[j] && t[j] <= 'Z'):
				tmp = tmp + string((t[j]-'A'+byte(r))%26+'A')
			case ('a' <= t[j] && t[j] <= 'z'):
				tmp = tmp + string((t[j]-'a'+byte(r))%26+'a')
			default:
				//何もしない
			}
		}

		return tmp
	})
	return res
}
