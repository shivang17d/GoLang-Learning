package main

import "fmt"

// 1. Define a custom type based on an integer (usually int)
type OrderStatus int

// 2. Use iota to auto-increment values
// iota resets to 0 whenever the keyword 'const' appears in the source code
// and increments after each Line Item.
const (
	StatusReceived  OrderStatus = iota // 0
	StatusPrepared                     // 1
	StatusShipped                      // 2
	StatusDelivered                    // 3
	StatusCanceled                     // 4
)

// 3. Implementing the Stringer interface (String() method)
// This makes the enum human-readable when printed.
func (s OrderStatus) String() string {
	switch s {
	case StatusReceived:
		return "Received"
	case StatusPrepared:
		return "Prepared"
	case StatusShipped:
		return "Shipped"
	case StatusDelivered:
		return "Delivered"
	case StatusCanceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

// 4. Using enums with other types and iota tricks
type Weekday int

const (
	// We can skip values or start from a specific number
	Sunday Weekday = iota + 1 // Starts at 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("--- Custom Enums with iota ---")

	status := StatusShipped
	fmt.Printf("Current Status: %v (Internal Value: %d)\n", status, status)

	if status == StatusShipped {
		fmt.Println("Package is on the way!")
	}

	fmt.Println("\n--- Enum in Switch Statement ---")
	checkStatus(StatusDelivered)
	checkStatus(StatusCanceled)
	checkStatus(OrderStatus(99)) // Handling unknown values

	fmt.Println("\n--- Starting iota from non-zero ---")
	fmt.Printf("Sunday is day number: %d\n", Sunday)
	fmt.Printf("Monday is day number: %d\n", Monday)
}

func checkStatus(s OrderStatus) {
	fmt.Print("Checking status: ")
	switch s {
	case StatusReceived, StatusPrepared:
		fmt.Println("Order is being processed.")
	case StatusShipped:
		fmt.Println("Order is in transit.")
	case StatusDelivered:
		fmt.Println("Order has been completed!")
	case StatusCanceled:
		fmt.Println("Order was canceled.")
	default:
		fmt.Printf("Invalid order status code: %d\n", s)
	}
}
