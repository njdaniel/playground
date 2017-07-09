package normal_switch

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Hi Daniel")
	case "Medhi":
		fmt.Println("Hi Medhi")
	case "Jenny":
		fmt.Println("Hi Jenny")
	default:
		fmt.Println("You have no friendsd")
	}
}

/*
no default fallthrough
no break needed, break is automatic
 */
