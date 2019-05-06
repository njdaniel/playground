package main

func main() {

}

//IsUnique detemine if string contains all unique characters
func IsUnique(s string) bool {
	u := make(map[rune]rune)
	for _, char := range s {
		if _, ok := u[char]; !ok {
			u[char] = char
		} else {
			return false
		}
	}
	return true
}

// IsUnique2 determine if string contains all unique char but no additional data structures
func IsUnique2(s string) bool {
	//assume not empty
	//sort
	//have two pointers compare if same return false otherwise true
	return true
}
