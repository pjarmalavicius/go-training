package teacher

import "errors"

type Teacher struct {
	name string
}

func NewTeacher(name string) *Teacher {
	return &Teacher{
		name: name,
	}
}

func (t *Teacher) Answer(question string) (string, error) {
	if question == "What is the meaning of life?" {
		return "42", nil
	}

	if question == "What is Go?" {
		return "Go is an awesome programing language. You have to love it!", nil
	}

	return "", errors.New("I don't know")
}
