package lesson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLesson_Start(t *testing.T) {
	tests := []struct {
		name      string
		mockFn    func(m *lessonMocks)
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name: "successful lesson",
			mockFn: func(m *lessonMocks) {
				question := "What is the meaning of life?"
				answer := "42"

				m.student.EXPECT().Ask().Return(question).Once()
				m.teacher.EXPECT().Answer(question).Return(answer, nil).Once()
				m.student.EXPECT().Remember(answer).Once()
			},
			assertErr: assert.NoError,
		},
		{
			name: "teacher fails to answer",
			mockFn: func(m *lessonMocks) {
				question := "What is the meaning of life?"

				m.student.EXPECT().Ask().Return(question).Once()
				m.teacher.EXPECT().Answer(question).Return("", assert.AnError).Once()
			},
			assertErr: assert.Error,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mocks := newLessonMocks(t)

			tc.mockFn(mocks)

			l := &Lesson{
				student: mocks.student,
				teacher: mocks.teacher,
			}
			tc.assertErr(t, l.Start())
		})
	}
}

type (
	lessonMocks struct {
		student *mockStudent
		teacher *mockTeacher
	}
)

func newLessonMocks(t *testing.T) *lessonMocks {
	return &lessonMocks{
		student: newMockStudent(t),
		teacher: newMockTeacher(t),
	}
}
