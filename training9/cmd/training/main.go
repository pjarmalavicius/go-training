package main

import (
	"log"

	"example.com/training/lesson"
	"example.com/training/student"
	"example.com/training/teacher"
)

func main() {
	l := lesson.NewLesson(
		student.NewStudent(
			"Jonas"),
		teacher.NewTeacher("Paulius"),
	)

	if err := l.Start(); err != nil {
		log.Fatalf("lesson failed: %v", err)
	}

	log.Println("lesson has ended successfully")
}
