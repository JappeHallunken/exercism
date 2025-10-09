package allergies

var algs = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
var algMap = make(map[uint]string)

func Allergies(allergies uint) []string {
	// generate the map {1: "eggs", 2: "peanuts" ... 128: "cats",}
	for i, name := range algs {
		algMap[1<<i] = name
	}
	result := []string{}

	for score, a := range algMap {
		if allergies&uint(score) > 0 {
			result = append(result, a)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	valueMap := make(map[string]uint)
	// invert key and value: {"eggs": 1, "peanuts": 2, ...}
	for k, v := range algMap {
		valueMap[v] = k
	}

	if _, ok := valueMap[allergen]; !ok {
		return false //allgern not in map
	}

	if allergies&valueMap[allergen] == 0 {
		return false // not allergic to given allergen
	}
	return true
}
