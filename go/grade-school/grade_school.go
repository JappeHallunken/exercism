package school

import (
	"slices"
)

type Grade struct {
	Level    int
	Students []string
}

type School struct {
	grades map[int][]string
}

func New() *School {
	return &School{
		grades: make(map[int][]string),
	}
}

func (s *School) Add(student string, grade int) {
	s.grades[grade] = append(s.grades[grade], student)
}

func (s *School) Grade(level int) []string {
	if students, ok := s.grades[level]; ok {
		return students
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	result := make([]Grade, 0, len(s.grades))

	for level, students := range s.grades {
		// Schüler alphabetisch sortieren
		slices.Sort(students)
		result = append(result, Grade{
			Level:    level,
			Students: students,
		})
	}

	slices.SortFunc(result, func(a, b Grade) int {
		return a.Level - b.Level
	})

	return result
}
