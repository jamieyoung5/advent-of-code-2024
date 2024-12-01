package util

func Counts[V comparable](vs []V) map[V]int {
	return CountsFunc(vs, func(v V) V { return v })
}

func CountsFunc[V any, K comparable](vs []V, key func(V) K) map[K]int {
	h := make(map[K]int)
	for _, v := range vs {
		h[key(v)]++
	}
	return h
}
