package main

import (
	"database/sql"
	crud "sql-transaction/crud"
	mod "sql-transaction/models"

	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
)

func main() {

	// connect to sql
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// student struct
	CreStudent := []mod.Student{
		{
			FirstName: "Xasan",
			LastName:  "Nosirov",
			Age:       17,
			Courses: []mod.Course{
				{
					Name:        "Java",
					TeacherName: "Alisher Kasimov",
					Price:       750,
				},
				{
					Name:        "Flutter",
					TeacherName: "Samandar Ahadjonov",
					Price:       600,
				},
			},
		},
		{
			FirstName: "Asadbek",
			LastName:  "Faxriddinov",
			Age:       17,
			Courses: []mod.Course{
				{
					Name:        "Node Js",
					TeacherName: "Saud Abdulwahed",
					Price:       650,
				},
				{
					Name:        "Vue Js",
					TeacherName: "Saud Abdulwahed",
					Price:       650,
				},
			},
		},
		{
			FirstName: "Tohirjon",
			LastName:  "Odilov",
			Age:       19,
			Courses: []mod.Course{
				{
					Name:        ".Net",
					TeacherName: "Mukhammadkarim Tukhtaboev",
					Price:       700,
				},
				{
					Name:        "React Js",
					TeacherName: "Saud Abdulwahed",
					Price:       650,
				},
			},
		},
	}

	// create students
	studentsRes, err := crud.CreateAllStudent(db, CreStudent)
	if err != nil {
		panic(err)
	}
	pp.Println(studentsRes)

	// update students
	studentsUp, err := crud.UpdateAllStudent(db, studentsRes)
	if err != nil {
		panic(err)
	}
	pp.Println(studentsUp)

	// course struct
	CreCourse := []mod.Course{
		{
			Name:        "C++",
			TeacherName: "Akbarshox Sattarov",
			Price:       450,
			Students: []mod.Student{
				{
					FirstName: "Ahrorbek",
					LastName:  "Alijonov",
					Age:       17,
					Courses:   []mod.Course{},
				},
				{
					FirstName: "Akramjon",
					LastName:  "Abduvahobov",
					Age:       15,
					Courses:   []mod.Course{},
				},
			},
		},
		{
			Name:        "C#",
			TeacherName: "Khumoyum",
			Price:       500,
			Students: []mod.Student{
				{
					FirstName: "Alisher",
					LastName:  "Botirov",
					Age:       23,
					Courses:   []mod.Course{},
				},
				{
					FirstName: "Botir",
					LastName:  "Alisherov",
					Age:       32,
					Courses:   []mod.Course{},
				},
			},
		},
		{
			Name:        "Go",
			TeacherName: "Nurali Uktamov",
			Price:       900,
			Students: []mod.Student{
				{
					FirstName: "Ahrorbek",
					LastName:  "Olimjonov",
					Age:       21,
					Courses:   []mod.Course{},
				},
				{
					FirstName: "Oybek",
					LastName:  "Atamatov",
					Age:       22,
					Courses:   []mod.Course{},
				},
				{
					FirstName: "Doston",
					LastName:  "Shernazarov",
					Age:       23,
					Courses:   []mod.Course{},
				},
			},
		},
	}

	// create courses
	courseRes, err := crud.CreateAllCourse(db, CreCourse)
	if err != nil {
		panic(err)
	}
	pp.Println(courseRes)

	// update courses
	coursesUp, err := crud.UpdateAllCourse(db, courseRes)
	if err != nil {
		panic(err)
	}
	pp.Println(coursesUp)

}
