-- queries for create tables on your desktop laptop

CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    age INT NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY NOT NULL,
    course_name VARCHAR(64) NOT NULL,
    teacher VARCHAR(64) NOT NULL,
    price INT NOT NULL
);

CREATE TABLE IF NOT EXISTS student_course (
    course_id INT REFERENCES courses(id),
    student_id INT REFERENCES students(id)
);