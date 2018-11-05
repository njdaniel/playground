package _1_constant

import "fmt"

const p string = "death and taxes"

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

}

// Constant is a simple unchanging value
// References: blog.golang.org/constants
// Golang is a statically typed language that does not permit mix numeric types
//   aka. cant add a float64 to an int
// BUT its legal to write 1e6*time.Second OR Exp(1) OR 1<<('\t'+2.0)
//   ....WTF. Ok lets read the post to explain why this is and what it means.

// Background:
//   C and others let you mix and match numeric types(it did? ok). Many mysterious bugs and crashes and portablitiy
//   problems are caused by combining int of dif sizes and "signedness"..
//   // So these are problems caused by adding float and int together? or 0001 + 01? "signedness"?

//   Some code about signed and unsigned.. // vhat?
//   // I google:  'signed and unsigned int' => unsigned cant be neg. signed CAN
//   // OK that makes sense, moving on
