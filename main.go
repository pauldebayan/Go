package main

import "fmt"

const secondsInHour = 3600


func main(){
	fmt.Println("Hello World!")
	distance := 60.8 // distance in km
	fmt.Println("The distance in miles", distance*0.621)


	var age1 int = 30;
	var age2 = 27;

	fmt.Println("age1", age1, "age2", age2)

	var name = "Debayan"

	s := "Learning Golang!" // Block scope

	fmt.Println(s, name)


	car, cost := "Audi", 500000

	fmt.Println("car", car, "cost", cost)

	const days int = 7


	price := 100

	if price > 80 {
		fmt.Println("Greater than 80!")
	}else if price == 100 {
		fmt.Println("Price is equal to 100")
	}

	fmt.Println(days)


	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}



}