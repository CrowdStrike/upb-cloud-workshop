package main

import (
	"math/rand"
)
type Visitor interface {
	visitSchool(school *School)
}	 

type PisaExam struct{
	name string
	ageLimits []int
}

func newPisaExam(n string, l []int) *PisaExam {
	return &PisaExam{name: n, ageLimits: l}
}

func (f *PisaExam) visitSchool(school *School) {
	lowLimit := f.ageLimits[0]
	highLimit := f.ageLimits[1]
	for _, student := range school.students {
		if student.age >= lowLimit && student.age <= highLimit {
			student.grades[f.name] = rand.Intn(10)
		}
	}
}

type BacExam struct {
	name string
	requiredAge int
}

func newBacExam(n string, r int) *BacExam {
	return &BacExam{name: n, requiredAge: r}
}

func (f *BacExam) visitSchool(school *School) {
	for _, student := range school.students {
		if student.age == f.requiredAge {
			student.grades[f.name] = rand.Intn(10)
		}
	}
}



