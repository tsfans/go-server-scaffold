package utils

func MapElement[S, T any](source []S, fn func(*S) *T) []T {
	target := make([]T, len(source))
	for i, v := range source {
		target[i] = *fn(&v)
	}
	return target
}

func ToAnySlice[T any](input []T) []any {
	result := make([]any, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}
