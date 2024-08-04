package code

// LinearSearch performs a linear search on a list to find a value.
// Returns true if the value is found in the array arr, otherwise false
func LinearSearch(arr *[]any, v any) bool {
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == v {
			return true
		}
	}
	return false
}
