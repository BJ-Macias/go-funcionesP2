package main

import "fmt"

func fibonacci(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func mayor(num ...int) int {
	var grande int = num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > grande {
			grande = num[i]
		}
	}
	return grande
}

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var inp = i
		i += 2
		return inp
	}
}

func intercambia(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func main() {
	
	fmt.Println(fibonacci(10))

	arr := []int{2, 10, 6, 14, 11, 7, 20, 35}
	fmt.Println(mayor(arr...))

	impares := generadorImpares()
	fmt.Println(impares())
	fmt.Println(impares())
	fmt.Println(impares())

	var a int
	var b int
	fmt.Print("Ingresa valor de 'a': ")
	fmt.Scan(&a)
	fmt.Print("Ingresa valor de 'b': ")
	fmt.Scan(&b)
	fmt.Println(a, b)
	intercambia(&a, &b)
	fmt.Println(a, b)
}