package lesson

import "fmt"

type (
	Lesson struct {
		student student
		teacher teacher
	}

	student interface {
		Ask() string
		Remember(answer string)
	}

	teacher interface {
		Answer(question string) (string, error)
	}
)

func NewLesson(s student, t teacher) *Lesson {
	return &Lesson{
		student: s,
		teacher: t,
	}
}

func (l *Lesson) Start() error {
	question := l.student.Ask()

	answer, err := l.teacher.Answer(question)
	if err != nil {
		return fmt.Errorf("teacher failed to answer: %w", err)
	}

	l.student.Remember(answer)

	return nil
}
