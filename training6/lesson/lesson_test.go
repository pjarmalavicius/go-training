package lesson

import (
	"errors"
	"testing"
)

func TestLesson_Start_Success(t *testing.T) {
	l := NewLesson(
		&studentMock{
			question: "What is the meaning of life?",
		},
		&teacherMock{
			err:    nil,
			answer: "this is my answer",
		},
	)

	if err := l.Start(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLesson_Start_Fail(t *testing.T) {
	l := NewLesson(
		&studentMock{
			question: "What is the meaning of life?",
		},
		&teacherMock{
			err: errors.New("I don't know"),
		},
	)

	if err := l.Start(); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

type (
	studentMock struct {
		question string
	}
	teacherMock struct {
		answer string
		err    error
	}
)

func (s *studentMock) Ask() string {
	return s.question
}

func (s *studentMock) Remember(answer string) {}

func (t *teacherMock) Answer(question string) (string, error) {
	return t.answer, t.err
}
