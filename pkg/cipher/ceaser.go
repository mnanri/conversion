package cipher

import "conversion/pkg"

//rはROT:r文字ずらす
func Ceaser(s []string, r int) []string {
	var res []string = pkg.Map(s, func(t string) string {
		tmp := ""
		sub := "AZaz"
		for i := range t {
			switch {
			case (sub[0] <= t[i] && t[i] <= sub[1]):
				tmp = tmp + string((t[i]-sub[0]+byte(r))%26+sub[0])
			case (sub[2] <= t[i] && t[i] <= sub[3]):
				tmp = tmp + string((t[i]-sub[2]+byte(r))%26+sub[2])
			default:
				//何もしない
			}
		}
		return tmp
	})
	return res
}
