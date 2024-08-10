package utils

func ExistsStringInArray(v string, target []string) bool {
	for i := 0; i < len(target); i++ {
		if target[i] == v {
			return true
		}
	}

	return false
}
