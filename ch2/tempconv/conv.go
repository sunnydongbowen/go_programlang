package tempconv

// Celsius @program:   go-basic-exercises
// @file:      conv.go
// @author:    bowen
// @time:      2023-02-08 10:19
// @description:

// CToF 摄氏温度转换成华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC  华氏温度转换成摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
