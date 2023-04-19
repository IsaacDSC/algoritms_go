package reduce

func Execute[T, M any](list []T, function func(M, T) M, initValue M) M {
	acc := initValue
	for _, value := range list {
		acc = function(acc, value)
	}
	return acc
}
