package main

import "fmt"

func main() {
	// --- PART 1: Basic Variables ---

	// 1. String Variables
	var name string = "Golang"
	fmt.Println("String:", name)

	// 2. Integer Variables
	var age int = 15
	fmt.Println("Integer:", age)

	// 3. Boolean Variables
	var isCool bool = true
	fmt.Println("Boolean:", isCool)

	// 4. Float Variables
	var version float64 = 1.23
	fmt.Println("Float:", version)

	// --- PART 2: Advanced Concepts ---

	// 5. Short Variable Declaration (Type Inference)
	// Only works inside functions. No 'var' keyword needed.
	inferredString := "I am inferred"
	inferredInt := 100
	fmt.Println("Inferred:", inferredString, inferredInt)

	// 6. Zero Values
	// Variables declared without a value take their zero value.
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Printf("Zero Values: int: %d, string: %q, bool: %v\n", zeroInt, zeroString, zeroBool)

	// 7. Multiple Variable Declaration
	var width, height int = 100, 50
	fmt.Println("Multiple:", width, height)

	// Multiple short declaration
	x, y := 10, "ten"
	fmt.Println("Multiple Types:", x, y)

	// --- PART 3: Bit-sized Integers & Float Precision ---

	// 8. Bit-sized Integers
	var i8 int8 = 127                   // 8-bit integer (-128 to 127)
	var i16 int16 = 32767               // 16-bit integer
	var i32 int32 = 2147483647          // 32-bit integer
	var i64 int64 = 9223372036854775807 // 64-bit integer
	fmt.Println("Bit-sized Integers:", i8, i16, i32, i64)

	// 9. Floats (float32 vs float64)
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359 // More precision (default for float)
	fmt.Printf("Floats: %f (32-bit), %f (64-bit)\n", f32, f64)

	// 10. Type Inference details
	pi := 3.14 // Infers float64
	fmt.Printf("Inferred Float Type: %f is type %T\n", pi, pi)
}
