package animal

type Animal struct {
	Name     string
	Logogram string
	Emoji    string
}

func CreateManyFromCSV(data [][]string) []Animal {
	var animals []Animal
	for i, row := range data {
		if i < 1 {
			continue
		}
		var animal Animal
		for j, field := range row {
			switch j {
			case 0:
				animal.Name = field
			case 1:
				animal.Logogram = field
			case 2:
				animal.Emoji = field
			}
		}
		animals = append(animals, animal)
	}
	return animals
}
