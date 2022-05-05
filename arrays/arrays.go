package arrays

//Equal
func Equal[T comparable](a, b []T) bool {
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
