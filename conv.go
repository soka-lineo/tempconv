package tempconv

// CToF は摂氏の温度を華氏へ変換する
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏の温度を摂氏へ変換する
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
