package people

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type People []Person

func MakePeople() People {
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
