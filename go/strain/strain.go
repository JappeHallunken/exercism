package strain

// Implement the "Keep" and "Discard" function in this file.

func Keep[T any](list []T, filterFunc func(T) bool) []T {
	var result []T
	for _, item := range list {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

func Discard[T any](list []T, filterFunc func(T) bool) []T {
	var result []T
	for _, item := range list {
		if !filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

