package student

type Student struct {
	Name   string
	Memory []string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) Ask() string {
	return "What is Go?"
}

func (s *Student) Remember(answer string) {
	s.Memory = append(s.Memory, answer)
}
