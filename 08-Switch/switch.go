package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Basic Switch
	i := 2
	fmt.Print("1. Basic: ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	// 2. Multiple Expressions
	day := "Saturday"
	fmt.Print("2. Multiple: ")
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a workday")
	}

	// 3. Tagless Switch (Conditions)
	t := time.Now()
	fmt.Print("3. Tagless: ")
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 4. Switch with Function/Expression
	fmt.Print("4. Expression: ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend (Time check)")
	default:
		fmt.Println("It's a workday (Time check)")
	}

	// 5. Fallthrough
	num := 1
	fmt.Print("5. Fallthrough: ")
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two (triggered by fallthrough)")
	case 3:
		fmt.Println("Three")
	}

	// 6. Short Declaration in Switch
	fmt.Print("6. Short Declaration: ")
	switch os := "macOS"; os {
	case "macOS":
		fmt.Println("Apple")
	case "Linux":
		fmt.Println("Penguin")
	default:
		fmt.Println("Windows?")
	}

	// 7. Type Switch (Value Assignment)
	fmt.Print("7. Type Switch (With Value): ")
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool with value: %v\n", t)
		case int:
			fmt.Printf("I'm an int with value: %d\n", t)
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)

	// 8. Type Switch (No Assignment)
	// Useful when you only need to check the type, not use the value as that type.
	fmt.Print("8. Type Switch (No Value): ")
	checkType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Unknown type")
		}
	}
	checkType(100)
}
