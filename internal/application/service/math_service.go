package service

import "example/hona/internal/application/dto/math"

type MathService struct{}

func NewMathService() *MathService {
	return &MathService{}
}

func (m *MathService) Adder(p math.AddRequest) math.AddResponse {
	sum := p.Num1 + p.Num2
	return math.AddResponse{
		Sum: sum,
	}
}

func (m *MathService) SayHello(name string) string {
	return "hello" + name
}
