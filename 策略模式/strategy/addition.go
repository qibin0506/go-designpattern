package strategy

/**
 * 加法
 */
type Addition struct {}

func (p Addition) Compute(num1 ,num2 int) int {
	return num1 + num2
}