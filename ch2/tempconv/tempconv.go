package tempconv

import "fmt"

// @program:   go-basic-exercises
// @file:      temp conv.go
// @author:    bowen
// @time:      2023-02-07 17:38
// @description:

//	func sum(a,b int) int{
//		return int(a)
//	}
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.5g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.5g°F", f)
}
