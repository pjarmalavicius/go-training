package lesson

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLesson_Start_Success(t *testing.T) {
	mockedStudent := &studentMock{}
	mockedTeacher := &teacherMock{}
	question := "What is the meaning of life?"
	answer := "42"

	mockedStudent.On("Ask").
		Return(question).
		Once()
	mockedTeacher.On("Answer", question).
		Return(answer, nil).
		Times(1)
	mockedStudent.On("Remember", answer).Times(1)

	t.Cleanup(func() {
		mockedStudent.AssertExpectations(t)
		mockedTeacher.AssertExpectations(t)
	})

	l := NewLesson(mockedStudent, mockedTeacher)

	assert.NoError(t, l.Start())
}

func TestLesson_Start_Fail(t *testing.T) {
	mockedStudent := &studentMock{}
	mockedTeacher := &teacherMock{}
	question := "What is the meaning of life?"

	mockedStudent.On("Ask").
		Return(question).Times(1)
	mockedTeacher.On("Answer", question).
		Return("", assert.AnError).Times(1)

	t.Cleanup(func() {
		mockedStudent.AssertExpectations(t)
		mockedTeacher.AssertExpectations(t)
	})

	l := NewLesson(mockedStudent, mockedTeacher)

	assert.Error(t, l.Start())
}

type (
	studentMock struct {
		mock.Mock
	}
	teacherMock struct {
		mock.Mock
	}
)

func (s *studentMock) Ask() string {
	args := s.Called()
	return args.String(0)
}

func (s *studentMock) Remember(answer string) {
	s.Called(answer)
}

func (t *teacherMock) Answer(question string) (string, error) {
	args := t.Called(question)
	return args.String(0), args.Error(1)
}
