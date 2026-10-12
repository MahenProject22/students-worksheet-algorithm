package main

import "fmt"

// TODO 5: pi is a constant. Try writing `pi = 3.15` inside main,
// read the compiler error, then remove that line again.
const pi = 3.14

func main() {
	// TODO 1 (int): integer division and remainder.
	a, b := 7, 2
	quotient := 0  // TODO: replace 0 with a / b
	remainder := 0 // TODO: replace 0 with a % b
	fmt.Printf("int: %d / %d = %d, %d %% %d = %d\n", a, b, quotient, a, b, remainder)

	// TODO 2 (float64): convert BEFORE dividing. Go has no implicit conversion.
	ratio := 0.0 // TODO: replace 0.0 with float64(a) / float64(b)
	fmt.Printf("float64: %.1f\n", ratio)

	// TODO 3 (bool): the student passes only if BOTH conditions are true.
	score, attendance := 72, 70
	passed := false // TODO: score >= 60 && attendance >= 75
	fmt.Println("bool: passed =", passed)

	// TODO 4 (string): join three pieces with + and print the byte length.
	greeting := "" // TODO: "Hello" + ", " + "Go"
	fmt.Printf("string: %s (len %d)\n", greeting, len(greeting))

	// TODO 6 (rune): a rune is a number. Move one letter forward.
	letter := 'A'
	next := letter // TODO: letter + 1
	fmt.Printf("rune: %c -> %d -> %c\n", letter, letter, next)

	// TODO 7 (const + float64): area of a circle with radius 5.
	radius := 5.0
	area := 0.0 // TODO: pi * radius * radius
	fmt.Printf("const: area = %.1f\n", area)

	// Keep the compiler quiet until you use these variables.
	_, _, _ = score, attendance, radius
}
