package utils

// ArrayContains - tells whether arr contains x.
func ArrayContains(arr []string, x string) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}
