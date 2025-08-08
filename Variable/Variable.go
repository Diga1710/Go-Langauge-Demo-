package main

func main{
	var n int;
	var s string;
	var d float64;
	fmt.Println("Enter an integer:")
	fmt.Scanln(&n)
	fmt.Println("Enter a string:")
	fmt.Scanln(&s)
	fmt.Println("Enter a double:")
	fmt.Scanln(&d)
	fmt.Println("You entered:", n, s, d)
	fmt.Println("Goodbye!")
}