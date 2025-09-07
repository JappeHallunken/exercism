package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {

	t := option1
	if option2 < option1 {
		t = option2
	}
	return t + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age >= 10:
		return originalPrice * .5
	case 3 <= age && age < 10:
		return originalPrice * .7
	case age < 3:
		return originalPrice * .8
	default:
		return originalPrice

	}
}
