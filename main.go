package main

import "fmt"
import "go-tutorials/shape"

func main() {
	const a = 5
	fmt.Println(a)
	fmt.Printf("hello, world\n")
	print(8, 9)

	shape.Draw()
}

func print(num, t int) {
	fmt.Println(num,  t)
}

func init() {
	fmt.Println("init func in main package")
}
