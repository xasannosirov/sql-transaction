package crud

import (
	"database/sql"
	mod "sql-transaction/models"
)

// this function creates student with connected courses
func CreateAllStudent(db *sql.DB, students []mod.Student) ([]mod.Student, error) {
	// start the transaction
	tx, txErr := db.Begin()
	if txErr != nil {
		panic(txErr)
	}

	var AllStudents []mod.Student
	for _, student := range students {
		// insert student data
		insertStudentQuery := `
		INSERT INTO students(
			first_name,
			last_name,
			age
		)
		VALUES
			($1, $2, $3)
		RETURNING
			id, 
			first_name,
			last_name,
			age`

		// run the query
		studentRow := tx.QueryRow(insertStudentQuery, student.FirstName, student.LastName, student.Age)

		// scan the query result
		var NewStudent mod.Student
		if err := studentRow.Scan(
			&NewStudent.Id,
			&NewStudent.FirstName,
			&NewStudent.LastName,
			&NewStudent.Age,
		); err != nil {
			tx.Rollback()
			panic(err)
		}

		for _, course := range student.Courses {
			// insert course data
			insertCourseQuery := `
			INSERT INTO courses (
				course_name,
				teacher,
				price
			)
			VALUES
				($1, $2, $3)
			RETURNING
				id,
				course_name,
				teacher,
				price`

			// run the query
			courseRow := tx.QueryRow(insertCourseQuery, course.Name, course.TeacherName, course.Price)

			// scan the query result
			var NewCourse mod.Course
			if err := courseRow.Scan(
				&NewCourse.Id,
				&NewCourse.Name,
				&NewCourse.TeacherName,
				&NewCourse.Price,
			); err != nil {
				tx.Rollback()
				panic(err)
			}

			// connect student and course
			connectQuery := `
			INSERT INTO student_course(
				student_id,
				course_id
			)
			VALUES ($1, $2)
			`

			// run the query
			_, err := tx.Exec(connectQuery, NewStudent.Id, NewCourse.Id)
			if err != nil {
				tx.Rollback()
				panic(err)
			}

			NewStudent.Courses = append(NewStudent.Courses, NewCourse)
		}
		AllStudents = append(AllStudents, NewStudent)
	}

	// finish the transaction
	if ComErr := tx.Commit(); ComErr != nil {
		tx.Rollback()
		panic(ComErr)
	}

	return AllStudents, nil
}

// this function creates course with connected students
func CreateAllCourse(db *sql.DB, courses []mod.Course) ([]mod.Course, error) {
	// start the transaction
	tx, txErr := db.Begin()
	if txErr != nil {
		return []mod.Course{}, txErr
	}

	var AllCourses []mod.Course

	for _, course := range courses {
		// insert course data
		insertCourseQuery := `
		INSERT INTO courses(
			course_name,
			teacher,
			price
		)
		VALUES
			($1, $2, $3)
		RETURNING
			id, 
			course_name,
			teacher,
			price`

		// run the query
		courseRow := tx.QueryRow(insertCourseQuery, course.Name, course.TeacherName, course.Price)

		// scan the query result
		var NewCourse mod.Course
		if err := courseRow.Scan(
			&NewCourse.Id,
			&NewCourse.Name,
			&NewCourse.TeacherName,
			&NewCourse.Price,
		); err != nil {
			tx.Rollback()
			panic(err)
		}

		for _, student := range course.Students {
			// insert student data
			insertStudentQuery := `
			INSERT INTO students (
				first_name,
				last_name,
				age
			)
			VALUES
				($1, $2, $3)
			RETURNING
				id,
				first_name,
				last_name,
				age`

			// run the query
			studentRow := tx.QueryRow(insertStudentQuery, student.FirstName, student.LastName, student.Age)

			// scan the query result
			var NewStudent mod.Student
			if err := studentRow.Scan(
				&NewStudent.Id,
				&NewStudent.FirstName,
				&NewStudent.LastName,
				&NewStudent.Age,
			); err != nil {
				tx.Rollback()
				panic(err)
			}

			// connect student and course
			connectQuery := `
			INSERT INTO student_course(
				student_id,
				course_id
			)
			VALUES ($1, $2)
			`
			// run the query
			_, err := tx.Exec(connectQuery, NewStudent.Id, NewCourse.Id)
			if err != nil {
				tx.Rollback()
				panic(err)
			}

			NewCourse.Students = append(NewCourse.Students, NewStudent)
		}
		AllCourses = append(AllCourses, NewCourse)
	}

	// finish the transaction
	if ComErr := tx.Commit(); ComErr != nil {
		tx.Rollback()
		return []mod.Course{}, ComErr
	}

	return AllCourses, nil
}
