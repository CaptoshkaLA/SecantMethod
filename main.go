package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Метод секущих")
	var a, b, eps float64
	k := 0
	yn := 0.0
	ynF := 0.0
	derivativeYn := 0.0
	fmt.Println("Задайте начальный промежуток неопределенности L0 и требуемую точность eps:")
	fmt.Fscan(os.Stdin, &a, &b, &eps)
	//fmt.Println(a, b, eps)
	secantMethod(a, b, eps, &yn, &ynF, &derivativeYn, &k)
	fmt.Println("На данном промежутке генерируется точка: ", yn)
	fmt.Println("Значение производной в этой точке: ", derivativeYn)
	fmt.Println("Количество итераций: ", k)
}
