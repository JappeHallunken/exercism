package flatten

func Flatten(nested interface{}) []interface{} {
	flat := make([]interface{}, 0)
	switch elem := nested.(type) {
	case []interface{}:
		for _, v := range elem {
			flat = append(flat, Flatten(v)...)
		}
	case interface{}:
		flat = append(flat, elem)
	}
	return flat
}
