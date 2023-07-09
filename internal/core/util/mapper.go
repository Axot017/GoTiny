package util

func MapSlice[TIn any, TOut any](items []TIn, mapper func(TIn) TOut) []TOut {
	result := make([]TOut, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}

	return result
}
