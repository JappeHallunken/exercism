package resistorcolorduo


var colors = []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
		return 0
	}
	
	firstColor := colors[0]
	secondColor := colors[1]

	firstIndex := indexOf(firstColor)
	secondIndex := indexOf(secondColor)

	if firstIndex == -1 || secondIndex == -1 {
		return 0
	}

	return firstIndex*10 + secondIndex
}

func indexOf(color string) int {
	for i, c := range colors {
		if c == color {
			return i
		}
	}
	return -1
}
	


