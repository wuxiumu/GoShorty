// internal/util/generic.go
package util

// MapReduce 泛型函数用于聚合数据
func MapReduce[T any, R any](list []T, mapper func(T) R, reducer func(R, R) R, zero R) R {
	var result = zero
	for _, item := range list {
		result = reducer(result, mapper(item))
	}
	return result
}
