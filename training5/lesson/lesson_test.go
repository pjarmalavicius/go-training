package lesson

import "testing"

func TestLesson_Start_Success(t *testing.T) {
	l := NewLesson(
		&studentMock{},
		&teacherMock{},
	)

	if err := l.Start(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

type (
	studentMock struct{}
	teacherMock struct{}
)

func (s *studentMock) Ask() string {
	return "foo bar baz"
}

func (s *studentMock) Remember(answer string) {}

func (t *teacherMock) Answer(question string) (string, error) {
	return "answer", nil
}
