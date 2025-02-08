package main

import "fmt"
import application "github.com/atoaidoocr7/distributed-systems/mapreduce/app"

func main() {
	fmt.Println("Hello World")
	m := application.Master{}
	m.Start()
}
