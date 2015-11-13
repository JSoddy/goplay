package main 

import(
	"fmt"
	"time"
	"math/rand"
	)

// a program which uses the extended Euclidean algorith to find for any integers
// A and B, integers s and t which satisfy the property sA + tB = gcd(A,B)

func main() {
	
	// get A and B
	A := pickANumber()
	B := pickANumber()
	// set a = b, b = B, s = 1, t = 0, u = 0 and v = 1
	a := A
	b := B
	s, t, u, v := 1, 0, 0, 1

	// loop: while b != 0
	for b != 0 {

	// set r = a % b, q = a / b
		r := a % b
		q := a / b

	// set a = b, b = r

		a, b = b, r
	// newu = s - u * q, newv = t - v * q

		newu := s - u * q
		newv := t - v * q

	// set s = u and t = v

		s, t = u, v

	// u = newu and v = newv
		u, v = newu, newv
	}
	// set gcd = a

	gcd := a

	// print gcd, s, and t

	fmt.Printf("GCD of %2d and %2d is %2d\n",A,B,gcd)
	fmt.Printf("%2d * %2d + %2d * %2d = %2d\n",A,s,B,t,gcd)
}

// Gives a random number between 1 and 50
func pickANumber() int {
	seedTime := time.Now()
	rand.Seed(int64(seedTime.Nanosecond()))
	return rand.Intn(50) + 1
}