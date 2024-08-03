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
		"Johny": {Name: "Johny", LastName: "Cage", Age: 33, Gender: "male", Crimes: 3},
		"Clark": {Name: "Clark", LastName: "Kent", Age: 29, Gender: "male", Crimes: 0},
		"Tony":  {Name: "Tony", LastName: "Soprano", Age: 45, Gender: "male", Crimes: 5},
		"Nina":  {Name: "Nina", LastName: "Williams", Age: 25, Gender: "female", Crimes: 2},
		"Eren":  {Name: "Eren", LastName: "Yeager", Age: 16, Gender: "male", Crimes: 1},
	}

	suspects := []string{"Tony", "Johny", "Eren"}

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
