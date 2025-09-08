package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
	"github.com/thisusami/donation-assignment/controller"
)

func main() {
	defer PanicRecovery()
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, proceeding with environment variables")
	}
	if len(os.Args) < 2 {
		panic("file path is required as an argument")
	}
	err := controller.NewController().Handler(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
func PanicRecovery() {
	if r := recover(); r != nil {
		fmt.Printf("panic recovered: %v stack trace: %s", r, string(debug.Stack()))
		os.Exit(1)
	}
}
