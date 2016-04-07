package strategy

type Multiplication struct {}

func (p Multiplication) Compute(num1, num2 int) int {
	return num1 * num2
}