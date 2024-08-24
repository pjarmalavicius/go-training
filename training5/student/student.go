package student

type Student struct {
	name   string
	memory []string
}

func NewStudent(name string) *Student {
	return &Student{
		name: name,
	}
}

func (s *Student) Ask() string {
	return "What is Go?"
}

func (s *Student) Remember(answer string) {
	s.memory = append(s.memory, answer)
}
