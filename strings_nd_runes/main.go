package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!"  // This won't discards escape sequence \n
	message1 := "hello, \tGo!" // this adds tab
	message2 := "hello, \rGo!" // this return cursor  to first position
	rawMessage := `Hello\nGo`  // for backticks it discards escape sequence \n, treats as string

	fmt.Println("message->", message)
	fmt.Println("message1->", message1)
	fmt.Println("message2->", message2)
	fmt.Println("rawmessage->", rawMessage)

	fmt.Println("length if rawmessage variable is", len(rawMessage))
	fmt.Println("The first charecter in message var is", message[0]) // returns ASCII value

	// string concatination
	greeting := "Hello"
	name := "Sayan"
	fmt.Println(greeting + " " + name)

	// string compare
	str1 := "Apple" // A has ASCII value of 65
	str := "apple"  // a has ASCII value of 97
	str2 := "banana"
	fmt.Println(str1 < str2)
	fmt.Println(str > str1, "<- str") // compiler checks

	// if i have to iterate string
	for idx, ch := range message {
		fmt.Printf("Charecter at index %d is %c of HEX value: %x \n", idx, ch, ch) // formatting verbs %d , %c
	}

	fmt.Println("Rume count: ", utf8.RuneCountInString(message))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a' // rune is single qouted
	jch := '音'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語"
	fmt.Println(NIHONGO)

	for _, runeVal := range NIHONGO {
		fmt.Printf("%v -> %c\n", runeVal, runeVal)
	}
}
