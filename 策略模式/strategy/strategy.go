package strategy

/**
 * 策略接口
 */
type Strategier interface {
	Compute(num1, num2 int) int
}

func NewStrategy(t string) (res Strategier) {
	switch t {
		case "s": // 减法
			res = Subtraction{}
		case "m": // 乘法
			res = Multiplication{}
		case "d": // 除法
			res = Division{}
		case "a": // 加法
			fallthrough
		default:
			res = Addition{}
	}

	return 
}