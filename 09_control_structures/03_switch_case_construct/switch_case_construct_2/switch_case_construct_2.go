// control_structures.go
package main

import "fmt"

func main() {
	// ----- if / else -----
	n := 10

	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	// if with a short statement
	if value := n * 2; value > 10 {
		fmt.Println("value is greater than 10:", value)
	} else {
		fmt.Println("value is not greater than 10:", value)
	}

	// ----- switch (value switch) -----
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	// ----- switch (tagless / expression switch) -----
	score := 82

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: D or lower")
	}

	// ----- for loop (classic) -----
	fmt.Println("Classic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// ----- for as a while loop -----
	fmt.Println("While-style for loop:")
	counter := 0
	for counter < 3 {
		fmt.Println("counter:", counter)
		counter++
	}

	// ----- infinite for loop with break -----
	fmt.Println("Infinite loop with break:")
	j := 0
	for {
		if j >= 3 {
			break
		}
		fmt.Println("j:", j)
		j++
	}

	// ----- range loop (over slice) -----
	fmt.Println("Range over slice:")
	numbers := []int{10, 20, 30}
	for idx, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", idx, value)
	}

	// range over map (order is not guaranteed)
	fmt.Println("Range over map:")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Printf("color %s -> %s\n", key, value)
	}
}
