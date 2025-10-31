package main

import (
	ecc "bitcoin-go/ellipticcurve"
	"fmt"
)

func main() {
	fe44 := ecc.NewFiniteElement(57, 44)
	fe33 := ecc.NewFiniteElement(57, 33)

	sum := fe44.Add(fe33)
	diff := fe44.Sub(fe33)
	neg := fe44.Neg()

	fmt.Println("Finite Element 1:", fe44)
	fmt.Println("Finite Element 2:", fe33)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Negation of Finite Element 44:", neg)
}
