package main

func main() {

}

//IsUnique detemine if string contains all characters
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
