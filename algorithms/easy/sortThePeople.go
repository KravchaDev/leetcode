package main

// сортировка пузырьком ↓
func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names)-1-i; j++ {
			if heights[j] < heights[j+1] {
				names[j], names[j+1] = names[j+1], names[j]
				heights[j], heights[j+1] = heights[j+1], heights[j]
			}
		}
	}
	return names
}

func main() {
	names := [8]string{"Драган", "Милица", "Никола", "Петар ", "Елена", "Марко", "Лука ", "Джэгода"}
	println("Неотсортированный массив:")
	for i := 0; i < len(names); i++ {
		println(names[i])
	}
	heights := [8]int{155, 168, 160, 195, 178, 153, 161, 193}
	sortPeople(names[:], heights[:])

	println("\nОтсортированный массив:")
	for i := 0; i < len(names); i++ {
		println(names[i])
	}
}
