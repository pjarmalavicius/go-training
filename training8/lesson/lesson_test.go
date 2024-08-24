package lesson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLesson_Start_Success(t *testing.T) {
	mockedStudent := newMockStudent(t)
	mockedTeacher := newMockTeacher(t)
	question := "What is the meaning of life?"
	answer := "42"

	mockedStudent.EXPECT().Ask().
		Return(question).
		Times(1)
	mockedTeacher.EXPECT().Answer(question).
		Return(answer, nil).
		Times(1)
	mockedStudent.EXPECT().Remember(answer).Times(1)

	l := NewLesson(mockedStudent, mockedTeacher)

	assert.NoError(t, l.Start())
}

func TestLesson_Start_Fail(t *testing.T) {
	mockedStudent := newMockStudent(t)
	mockedTeacher := newMockTeacher(t)
	question := "What is the meaning of life?"

	mockedStudent.EXPECT().Ask().
		Return(question).Times(1)
	mockedTeacher.EXPECT().Answer(question).
		Return("", assert.AnError).Times(1)

	l := NewLesson(mockedStudent, mockedTeacher)

	assert.Error(t, l.Start())
}
