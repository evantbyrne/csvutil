package lib

func ArraysEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, value := range a {
		if value != b[i] {
			return false
		}
	}
	return true
}

func rowsEqual(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, value := range a {
		if !ArraysEqual(value, b[i]) {
			return false
		}
	}
	return true
}
