package service

import "example/hona/internal/application/dto/math"

func Adder(p math.AddRequest) math.AddResponse {
	sum := p.Num1 + p.Num2
	return math.AddResponse{
		Sum: sum,
	}
}

func SayHello(name string) string {
	return "hello" + name
}
