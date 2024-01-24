package models

type Student struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Courses   []Course
}

type Course struct {
	Id          int
	Name        string
	TeacherName string
	Price       int
	Students    []Student
}
