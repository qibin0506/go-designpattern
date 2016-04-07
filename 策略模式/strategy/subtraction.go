package strategy

/**
 * 减法
 */
type Subtraction struct {}

func (p Subtraction) Compute(num1 ,num2 int) int {
	return num1 - num2
}