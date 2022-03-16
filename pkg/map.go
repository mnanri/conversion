package pkg

func Map[X, Y any](xs []X, f func(x X) Y) []Y {
	ys := make([]Y, len(xs))
	for i := range xs {
		ys[i] = f(xs[i])
	}
	return ys
}
