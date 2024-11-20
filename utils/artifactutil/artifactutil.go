// Package artifactutil
// @Description: Data structure processing
package artifactutil

// isInSlice Verify if the data is in the slice
func isInSlice[T comparable](target T, targetSlice []T, fn func(v T) T) bool {
	m := make(map[T]struct{})
	for _, v := range targetSlice {
		m[fn(v)] = struct{}{}
	}
	_, exists := m[target]
	return exists
}

// IsInSlice Verify if the data is in the slice
func IsInSlice[T comparable](target T, targetSlice []T) bool {
	return isInSlice(target, targetSlice, func(v T) T {
		return v
	})
}

// reDup Data slicing deduplication
func reDup[T comparable](slice []T, fn func(v T) T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		k := fn(v)
		if _, exists := seen[k]; !exists {
			seen[k] = struct{}{}
			result = append(result, k)
		}
	}
	return result
}

// ReDup Data slicing deduplication
func ReDup[T comparable](slice []T) []T {
	return reDup(slice, func(v T) T {
		return v
	})
}

// getMapValue Retrieve all values of the map
func getMapValue[E, T comparable](m map[E]T, fn func(v T) T) []T {
	res := make([]T, 0, len(m))
	for _, v := range m {
		va := fn(v)
		res = append(res, va)
	}
	return res
}

// GetMapValue Retrieve all values of the map
func GetMapValue[E, T comparable](m map[E]T) []T {
	return getMapValue(m, func(v T) T {
		return v
	})
}

// GetMapValueDup Retrieve all values of the map And not repeat
func GetMapValueDup[E, T comparable](m map[E]T) []T {
	return ReDup(GetMapValue(m))
}
