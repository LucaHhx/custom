package helper

// ArrayToMap 数组转换Map
func ArrayToMap[T, V any, K comparable](ts []T, f func(t T) (K, V)) map[K]V {
	m := make(map[K]V)
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

// ArrayDistinct 数组去重取属性
func ArrayDistinct[T any, Y comparable](ts []T, f func(t T) (y Y)) []Y {
	m := make(map[Y]bool)
	for _, t := range ts {
		m[f(t)] = true
	}
	var ys []Y
	for k := range m {
		ys = append(ys, k)
	}
	return ys
}
