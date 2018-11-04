package main

func main() {
	input := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}


	numUniqueEmails(input)
}

func numUniqueEmails(emails []string) int {
	//create unique list
	u := make([][]rune, 0, len(emails))
	m := make(map[string]bool)
	//separate lead and end
	//take out '.' and everything after '+'
	for _, email := range emails {
		//split name and domain
		email := []rune(email)
		name := make([]rune, 0)
		domain := make([]rune, 0)
		loc := 0
		for i, value := range email {
			if value == '@' {
				loc = i
				continue
			}
		}
		name = email[:loc]
		domain = email[loc:]
		for i := 0; i < len(name); {
			if name[i] == '.' {
				name = append(name[:i], name[i+1:]...)
			} else if name[i] == '+' {
				name = name[:i]
			} else {
				i++
			}
		}
		//combine name and domain
		email = append(name, domain...)
		if m[string(email)] == false {
			m[string(email)] = true
			u = append(u, email)
		}
	}
	return len(u)
	//return len(u)
}