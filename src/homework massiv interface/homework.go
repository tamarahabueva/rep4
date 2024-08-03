package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Ivan":    {Name: "Ivan", LastName: "Ilin", Age: 45, Gender: "male", Crimes: 1},
		"Andrey":  {Name: "Andrey", LastName: "Chernykh", Age: 40, Gender: "male", Crimes: 0},
		"Kseniya": {Name: "Kseniya", LastName: "Ivanova", Age: 23, Gender: "female", Crimes: 2},
		"Oleg":    {Name: "Oleg", LastName: "Repin", Age: 25, Gender: "male", Crimes: 0},
		"Mark":    {Name: "Mark", LastName: "Krylov", Age: 35, Gender: "male", Crimes: 3},
	}

	suspects := []string{"Ivan", "Andrey", "Kseniya", "Oleg", "Mark"}

	var mostCriminal Man
	var mostCriminalFound bool
	for _, name := range suspects {
		person, ok := people[name]
		if !ok {
			continue
		}

		if person.Crimes > mostCriminal.Crimes {
			mostCriminal = person
			mostCriminalFound = true
		}
	}

	if mostCriminalFound {
		fmt.Printf("Наиболее криминальная личность: %s %s", mostCriminal.Name, mostCriminal.LastName)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}
