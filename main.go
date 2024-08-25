// calculator/main.go
package main

import (
    "fmt"
    "github.com/avishek-nayal-sony/calculator"
)

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)
    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    switch operator {
    case "+":
        result := calculator.Add(num1, num2)
        fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
    case "-":
        result := calculator.Subtract(num1, num2)
        fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
    case "*":
        result := calculator.Multiply(num1, num2)
        fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
    case "/":
        result, err := calculator.Divide(num1, num2)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
        }
    default:
        fmt.Println("Invalid operator with sign")
    }
}

