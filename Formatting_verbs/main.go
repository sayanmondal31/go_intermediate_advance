package main

import "fmt"

func main() {
	// -- General Formatting verbs --
	// %v Prints the value in the default format
	// %#v Prints the value in Go-syntax format
	// %T Prints the type of the value
	// %% Prints the % sign

	i := 111_505.5
	string := "Hello World"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	// fmt.Printf("Add: %v \nSubtract: %v \nMultiply: %v \nDivide: %v", i+i, i-i, i*i, i/i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	// -- Integer Formatting Verbs
	// %b Base 2
	// %d Base 10
	// %+d Base 10 and always show sign
	// %o Base 8
	// %O Base 8 with leading 0o
	// %x Base 16, lowercase
	// %X Base 16, Uppercase
	// %#x Base 16, with leading 0x
	// %4d Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeros (width 4)

	fmt.Printf("\n\n ----------- INT verb check ----------\n\n")

	int := 255

	fmt.Printf("%b\n", int)   // binary - output: 11111111
	fmt.Printf("%d\n", int)   // mostly use - output: 255
	fmt.Printf("%+d\n", int)  // + output: +255 (tells integer with positive or negetive)
	fmt.Printf("%o\n", int)   // octal Base 8 output: 377
	fmt.Printf("%O\n", int)   // octal Base 8 with leading 0o output: 0o377
	fmt.Printf("%x\n", int)   // output: ff hex value lowercase
	fmt.Printf("%X\n", int)   // output: FF hex value capital case
	fmt.Printf("%#X\n", int)  // output: 0XFF
	fmt.Printf("%4d\n", int)  // output: [tab]255
	fmt.Printf("%-4d\n", int) // output: 255
	fmt.Printf("%04d\n", int) // output: 0255 (this fills the gap with 0)
	fmt.Printf("%05d\n", int) // output: 00255 (this fills the gap with 0)

	// -- string Formatting Verbs
	// %s Prints the value as plain string
	// %q Prints the value as a double-qouted string
	// %8s Prints the value as plain string (width 8, right justified)
	// %-8s Prints the value ad plain string (width 8, left justified)
	// %x Prints the value as hex dump of byte values
	// % x prints the value ad hex dump with spaces

	fmt.Printf("\n\n ----------- STRING verb check ----------\n\n")
	txt := "World"

	fmt.Printf("%s\n", txt)   // World
	fmt.Printf("%q\n", txt)   // "World"
	fmt.Printf("%8s\n", txt)  //   World
	fmt.Printf("%-8s\n", txt) //World
	fmt.Printf("%x\n", txt)   //576f726c64 <- hex value
	fmt.Printf("% x\n", txt)  // 57 6f 72 6c 64

	// -- Boolean Formatting Verbs
	// %t Value of the boolean operator in true or false format (same as using %v)

	fmt.Printf("\n\n ----------- Boolean verb check ----------\n\n")
	t := true
	f := false
	fmt.Printf("%t", t)
	fmt.Printf("%t", f)

	// -- Float Formatting Verbs
	// Verb Description
	// %e Scientific notation with 'e' ad exponent
	// %f Decimal point, no exponent
	// %.2f Default width, precision 2
	// %6.2f width 6 precision 2
	// %g Exponent as needed, only necessory digits
	fmt.Printf("\n\n ----------- Float verb check ----------\n\n")
	fit := 9.18

	fmt.Printf("%e\n", fit)
	fmt.Printf("%f\n", fit)
	fmt.Printf("%.2f\n", fit)
	fmt.Printf("%6.2f\n", fit)
	fmt.Printf("%g\n", fit)

}
