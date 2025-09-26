package sublist

import "slices"

// Relation type is defined in relations.go file.
// since go 1.21 you can compare slices with the slices.Equal function
// otherwise you would need 2 loops to iterate through bot slices and compare the int values

func Sublist(l1, l2 []int) Relation {
	len1, len2 := len(l1), len(l2)

	if slices.Equal(l1, l2) {
		return RelationEqual
	}

	if len1 == 0 {
		return RelationSublist
	}

	if len2 == 0 {
		return RelationSuperlist
	}

	if len1 < len2 && findSublist(l1, l2) {
			return RelationSublist
	}

	if len2 < len1  && findSublist(l2, l1) {
			return RelationSuperlist
	}

	return RelationUnequal
}

// findSubList determines, if a is part of b. a needs to be equal or less in length then b
func findSublist(a, b []int) bool {
	lenA, lenB := len(a), len(b)
	if lenA > lenB {
		return false
	}
	for i := 0; i <= lenB-lenA; i++ {
		if slices.Equal(a, b[i:i+lenA]) {
			return true
		}
	}
	return false
}
