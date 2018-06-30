package people

// Person type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// People as slice of Person
type People []Person

// GeneratePeople generates a dummy slice of people
func GeneratePeople() People {
	loris := Person{
		FirstName: "Loris",
		LastName:  "Guerra",
		Age:       26,
	}
	antonio := Person{
		FirstName: "Antonio",
		LastName:  "Antonini",
		Age:       30,
	}

	return People{loris, antonio}
}
