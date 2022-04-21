package main

import (
	v1 "github.com/donvito/hellomod"
	v2 "github.com/donvito/hellomod/v2"
)

func main() {
	//Versión 1
	v1.SayHello()
	//Versión 2
	v2.SayHello("Hola")
}
