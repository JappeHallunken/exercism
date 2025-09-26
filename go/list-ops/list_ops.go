package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Calculate length without len()
func (s IntList) Length() (n int) {
	defer func() { recover() }()
	for i := 0; ; i++ {
		_ = s[i]
		n++
	}
	return n
}

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := IntList{}
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s IntList) Map(fn func(int) int) IntList {
	j := 0
	for _, v := range s {
		s[j] = fn(v)
		j++
	}
	return s
}

func (s IntList) Reverse() IntList {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

// Append without append()
func (s IntList) Append(lst IntList) IntList {
	result := make([]int, len(s)+len(lst))
	copy(result, s)
	copy(result[len(s):], lst)
	return result
}

//Also no append() for concatenations
func (s IntList) Concat(lists []IntList) IntList {
	result := make(IntList, len(s))
	copy(result, s)

	for _, v := range lists {
		tmp := make([]int, len(result)+len(v))
		copy(tmp, result)
		copy(tmp[len(result):], v)
  	result = tmp
	}
	return result
}
