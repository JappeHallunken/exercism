package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	len1, len2 := len(l1), len(l2)
	switch {
	case len1 == 0 && len2 == 0:
		return RelationEqual

	case len1 == 0 && len2 != 0:
		return RelationSublist

	case len1 != 0 && len2 == 0:
		return RelationSuperlist

	case len1 == len2:
		for i := range l1 {
			if l1[i] == l2[i] {
				continue
			} else {
				return RelationUnequal
			}
		}
		return RelationEqual

	}
	return RelationUnequal
}
