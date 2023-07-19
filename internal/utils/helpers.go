package utils

// ArrayContains - tells whether arr contains x.
func ArrayContains[T comparable](arr *[]T, x T) bool {
	for _, n := range *arr {
		if x == n {
			return true
		}
	}
	return false
}
