package main

import "math"

func secantMethod(a, b, eps float64, min, minF, derivative *float64, k *int) {
	var Fa, Fb, y, Fy float64
	flag := false
	for flag == false {
		Fa = Derivative(a)
		Fb = Derivative(b)
		if Fa == 0 { // Если производные на концах отрезков = 0, то точка найдена
			*min = a
			*minF = Derivative(*min)
			*derivative = Fa
			flag = true
		} else if Fb == 0 {
			*min = b
			*minF = Function(*min)
			*derivative = Fb
			flag = true
		} else if Fa*Fb > 0 {
			if Fa > 0 && Fb > 0 {
				*min = a
				*minF = Function(*min)
				*derivative = Fa
				flag = true
			} else if Fa < 0 && Fb < 0 {
				*min = b
				*minF = Function(*min)
				*derivative = Fb
				flag = true
			}
		} else if Fa*Fb < 0 {
			y = a - Fa/(Fa-Fb)*(a-b) // формула касательной
			Fy = Derivative(y)
			if math.Abs(Fy) <= eps {
				*min = y
				*minF = Function(*min)
				*derivative = Fy
				flag = true
			} else {
				if Fy > 0 {
					b = y
				} else if Fy < 0 {
					a = y
				}
			}
		}
		*k++
	}
}
