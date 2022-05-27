package misc

import "golang.org/x/exp/constraints"

// InSlice checks to see if a value is in a given slice
func InSlice[T comparable](val T, list []T) bool {
	for _, key := range list {
		if key == val {
			return true
		}
	}
	return false
}

// UniqueSlice takes a slice and removes duplicates
func UniqueSlice[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// SliceEqual checks if two slices are equal
func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Pointer[T constraints.Ordered](v T) *T {
	return &v
}
