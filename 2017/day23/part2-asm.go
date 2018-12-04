// http://adventofcode.com/2017/day/23
// part 2 - assembly
package main

import (
	"fmt"
)

func main() {
	var b, d, e, f, h int

	b = 109300

loop1:
	f = 1
	d = 2

loop2:
	fmt.Printf("b%7v d%7v e%7v f%7v h%7v\n", b, d, e, f, h)
	e = 2

loop3:
	if d*e == b {
		f = 0
	}
	e++
	if e != b {
		goto loop3
	}

	d++
	if d != b {
		goto loop2
	}

	if f == 0 {
		h++
	}
	if b == 126300 {
		goto end
	}
	b += 17
	fmt.Printf("b%7v d%7v e%7v f%7v h%7v\n", b, d, e, f, h)
	goto loop1

end:
	fmt.Println(h)
}
