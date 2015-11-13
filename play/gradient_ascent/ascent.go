package main

import (
	"fmt"
	"math"
	)

func main() {
	alpha := 0.005
	delta := 10000.000
	ascent(2.3, 1.7, alpha, delta)
	ascent(-2.2, 0.1, alpha, delta)
	fmt.Println();
}

func ascent(x float64, y float64, alpha float64, delta float64){
	counter := 0
	fmt.Printf("\nStarting from point %.5f, %.5f.\n", x, y)
	for (delta > .00001) {
		xd, yd := gradient(x, y)
		delta = math.Sqrt(math.Pow(xd, 2) + math.Pow(yd, 2))
		x += xd * alpha
		y += yd * alpha
		counter++
	}
	fmt.Printf("Termination took %d iterations.\n", counter)
	fmt.Printf("The stopping point is %.5f, %.5f.\n", x, y)
	fmt.Printf("The max value is approximately %.5f.\n", fxy(x,y))
}

func gradient(x float64, y float64) (xd float64, yd float64){
	xd = (2 * x * math.Cos(math.Pow(x, 2) + math.Pow(y, 2)) * math.Cos(y + math.Exp(x))) - 
				(math.Exp(x) * math.Sin(math.Pow(x, 2) + math.Pow(y, 2)) * math.Sin(y + math.Exp(x)))
	yd = (2 * y * math.Cos(math.Pow(x, 2) + math.Pow(y, 2)) * math.Cos(y + math.Exp(x))) - 
				(math.Sin(math.Pow(x, 2) + math.Pow(y, 2)) * math.Sin(y + math.Exp(x)))
	return
}

func fxy(x float64, y float64) (z float64){
	z = math.Sin(math.Pow(x, 2) + math.Pow(y, 2)) * math.Cos(y + math.Exp(x))
	return
}