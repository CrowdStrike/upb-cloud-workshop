package main

import (
	"fmt"
)

// var _ Visitor = &Student{}
type College struct {
	School
	name string
}

func (c *College) accept(visitor Visitor) {
	fmt.Println("Accept college")
}

type School struct {
	name string
	students []*Student
}

func newSchool(n string, studs []*Student) *School {
	return &School{name: n, students: studs}
}

func (s *School) accept(visitor Visitor) {
	fmt.Println("Accept entered")
	visitor.visitSchool(s)
}

type Student struct {
	name string
	age int
	grades map[string]int
}

func newStudent(n string, a int) *Student {
	return &Student{
		name: n,
		age: a,
		grades: map[string]int{},
	}
}

// func changeAge(s *Student, newAge int) {
// 	s.age = newAge
// }
func (s *Student) setAge(newAge int) {
	s.age = newAge
}

func (s *Student) visit(x int, str string) {
	
}


func main() {
	// s1 := Student{name: "Ana", age: 12}
	// s1.setAge(16)
	// fmt.Println(s1.age)

	s2 := Student{name:"Dana", age:12, grades: map[string]int{}}
	fmt.Printf("Student s2: %+v\n", s2)
	
	s1 := newStudent("ana", 15)
	fmt.Printf("Student: %+v \n", *s1)

	s3 := Student{
		name: "Ana",
		age: 18,
		grades: map[string]int{},
	}

	students := []*Student{s1, &s2, &s3}
	school := newSchool("UPB", students)

	pisaExam := newPisaExam("pisa", []int{12,15})
	bacExam := newBacExam("bac", 18)
	fmt.Println(school.students, pisaExam, bacExam)

	school.accept(pisaExam)
	school.accept(bacExam)

	for _, student := range school.students {
		fmt.Println("Name: ", student.name)
		for k, v := range student.grades {
			fmt.Println(k, v)
		}
	}

	for _ = range make([]int, 10){
		fmt.Println("DA")
	}

	c := College{name: "Harvard"}
	c.accept(bacExam)
}