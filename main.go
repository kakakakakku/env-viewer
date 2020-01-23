package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ENV01 : %s\n", os.Getenv("ENV01"))
	fmt.Printf("ENV02 : %s\n", os.Getenv("ENV02"))
	fmt.Printf("ENV03 : %s\n", os.Getenv("ENV03"))
	fmt.Printf("ENV04 : %s\n", os.Getenv("ENV04"))
	fmt.Printf("ENV05 : %s\n", os.Getenv("ENV05"))
	fmt.Printf("ENV06 : %s\n", os.Getenv("ENV06"))
	fmt.Printf("ENV07 : %s\n", os.Getenv("ENV07"))
	fmt.Printf("ENV08 : %s\n", os.Getenv("ENV08"))
	fmt.Printf("ENV09 : %s\n", os.Getenv("ENV09"))
	fmt.Printf("ENV10 : %s\n", os.Getenv("ENV10"))
}
