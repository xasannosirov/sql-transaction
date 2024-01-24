package crud

import (
	"database/sql"
	mod "sql-transaction/models"
)

// this func update studen and studen.courses data
func UpdateAllStudent(db *sql.DB, students []mod.Student) ([]mod.Student, error) {
	// start the transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var AllStudents []mod.Student
	for _, student := range students {
		// updates student data and return
		updateStudenQuery := `
		UPDATE 
			students
		SET
			age = age+1
		WHERE
			id = $1
		RETURNING
			id,
			first_name,
			last_name,
			age`

		// run the query
		courseRow := tx.QueryRow(updateStudenQuery, student.Id)

		// scan the query result
		var NewStuden mod.Student
		if err := courseRow.Scan(
			&NewStuden.Id,
			&NewStuden.FirstName,
			&NewStuden.LastName,
			&NewStuden.Age,
		); err != nil {
			tx.Rollback()
			panic(err)
		}

		for _, course := range student.Courses {
			// update course data and return
			updateCourseQuery := `
			UPDATE
				courses
			SET
				price = price + price*0.1
			WHERE
				id = $1
			RETURNING
				id,
				course_name,
				teacher,
				price`

			// run the query
			courseRow := tx.QueryRow(updateCourseQuery, course.Id)

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

			NewStuden.Courses = append(NewStuden.Courses, NewCourse)
		}
		AllStudents = append(AllStudents, NewStuden)
	}

	// finish the transaction
	if ComErr := tx.Commit(); ComErr != nil {
		tx.Rollback()
		return []mod.Student{}, ComErr
	}

	return AllStudents, nil
}

// this func update course and course.student data
func UpdateAllCourse(db *sql.DB, courses []mod.Course) ([]mod.Course, error) {
	// start the transaction
	tx, err := db.Begin()
	if err != nil {
		return []mod.Course{}, err
	}

	var AllCourses []mod.Course
	for _, course := range courses {
		// updates course data and return
		updateCourseQuery := `
		UPDATE 
			courses
		SET
			price = price+price*0.1
		WHERE
			id = $1
		RETURNING
			id,
			course_name,
			teacher,
			price`

		// run the query
		courseRow := tx.QueryRow(updateCourseQuery, course.Id)

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
			// updates student data and return
			updateStudentQuery := `
			UPDATE
				students
			SET
				age = age+1
			WHERE
				id = $1
			RETURNING
				id,
				first_name,
				last_name,
				age`

			// run the query
			studentRow := tx.QueryRow(updateStudentQuery, student.Id)

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
