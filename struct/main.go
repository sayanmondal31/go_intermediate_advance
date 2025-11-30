package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	// embedding struct
	Address Address
	//embedded anonymous field
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	City    string
	Country string
}

// value receivers
// here p is an instance means copy of struct
func (p Person) fullName() string {
	return p.FirstName + " " + p.LastName
}

// pointer receiver

func (p *Person) incrementAgeOne() {
	p.Age++
}

func main() {

	// instances of Person struct
	p1 := Person{
		FirstName: "Sayan",
		LastName:  "Mondal",
		Age:       26,
		Address: Address{
			City:    "kolkata",
			Country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "9008098",
			cell: "78768768",
		},
	}

	p3 := Person{
		FirstName: "Sayan",
		LastName:  "Mondal",
		Age:       26,
		Address: Address{
			City:    "kolkata",
			Country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "9008098",
			cell: "78768768",
		},
	}

	p2 := Person{
		FirstName: "Sayan",
		LastName:  "Mondal",
		Age:       26,
		Address: Address{
			City:    "kolkata",
			Country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "9008098",
			cell: "78768768",
		},
	}

	fmt.Println(p1.FirstName, p1.LastName, p1.Age, "address: ", p1.Address.City, p1.Address.Country)
	fmt.Println("home phone", p1.home, p1.cell)

	fmt.Println("full name: ", p1.fullName())
	fmt.Println("before increment-<", p1.Age)
	p1.incrementAgeOne()
	fmt.Println("After increment-<", p1.Age)

	// struct comparison
	fmt.Println("p1 and p3 equal?", p2 == p3)

	// Anonymous struct
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "test@email.com",
	}

	fmt.Println(user.username)

	// ---------------- execution of methods ---------

	fmt.Println("---------------- execution of methods ---------")
	rect := Rectangle{
		Length: 10,
		Width:  20,
	}

	area := rect.Area()

	fmt.Println("Area of rectangle: ", area)

	rect.Scale(4)

	fmt.Println("Area of rect after scaling: ", rect.Area())

	num := myInt(5)
	num2 := myInt(-5)

	fmt.Println(num.WelcomeMessage())
	fmt.Println("is 5 positive?", num.IsPositive())           //true
	fmt.Println("is 5 positive for num2?", num2.IsPositive()) //false

	s := Shape{
		Rectangle: Rectangle{
			Length: 20,
			Width:  10,
		},
	}

	fmt.Println("shape area", s.Area())

}

// methods can be assiciated with any type
type myInt int

// in here we are creating instance in for value receiver
func (m myInt) IsPositive() bool {
	return m > 0
}

func (myInt) WelcomeMessage() string {
	return "Welcome to myInt type"
}
