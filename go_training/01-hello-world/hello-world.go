package main

import "fmt"

func main() {
	fmt.Printf("Enter any number from 1-5 \n")
	fmt.Printf("1: Addition \n")
	fmt.Printf("2: Sub  \n")
	fmt.Printf("1: Addition  \n")
	fmt.Printf("1: Addition  \n")
	fmt.Printf("1: Addition  \n")
	var x int
	fmt.Scanln(&x)

	if x == 1 {
		var n1, n2 int
		fmt.Println("Enter 2 numbers :")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
		var sum int
		sum = n1 + n2
		fmt.Printf("Sum of the number: %d\n", sum)
	}

	if x == 2 {
		var n1, n2 int
		fmt.Println("Enter 2 numbers :")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
		var sum int
		sum = n1 - n2
		fmt.Printf("Minus of the number: %d\n", sum)
	}

	if x == 3 {
		var n1, n2 int
		fmt.Println("Enter 2 numbers :")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
		var sum int
		sum = n1 * n2
		fmt.Printf("Multi of the number: %d\n", sum)
	}

	if x == 4 {
		var n1, n2 int
		fmt.Println("Enter 2 numbers :")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
		var sum int
		sum = n1 / n2
		fmt.Printf("Divide of the number: %d\n", sum)
	}

	if x == 5 {
		print("Exit")

	}

	// if

	// var userName string
	// fmt.Println("Enter your name :")
	// fmt.Scanln(&userName)
	// fmt.Printf("Hi %s, Have a nice day!\n", userName)

	// var n1, n2 int
	// fmt.Println("Enter 2 numbers :")
	// fmt.Scanln(&n1, &n2)
	// fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)

	// Accepting data with spaces
	/*
		var scanner bufio.Scanner
		scanner = *bufio.NewScanner(os.Stdin)
		fmt.Println("Enter a sentence:")
		scanner.Scan()
		txt := scanner.Text()
		fmt.Println(strings.ToUpper(txt))
	*/
}
