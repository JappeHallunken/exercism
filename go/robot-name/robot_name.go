package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var (
	namePool = generateRobotNames()
	idx      = 0
)

const maxRNames = 26 * 26 * 1000 // two uppercase letters and 000 to 999

func generateRobotNames() []string {
	pos := 0
	names := make([]string, maxRNames)

	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			for k := 0; k < 1000; k++ {
				names[pos] = fmt.Sprintf("%s%s%03d", string(i), string(j), k)
				pos++
			}
		}
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if idx >= maxRNames {
		return "", fmt.Errorf("uniqueness exhausted")
	}

	r.name = namePool[idx]
	idx++

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
