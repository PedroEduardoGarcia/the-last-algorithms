package code

// LinearSearch returns true if the value is found in the array arr, otherwise false
func LinearSearch(arr *[]any, v any) bool {
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == v {
			return true
		}
	}
	return false
}
