package main

import (
	"belajar-golang/helper"
	"belajar-golang/others"
	"fmt"
)

func main() {
	helper.SayHello("Salman")
	others.SayHelloToOthers()
	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// helper.sayGoodbye()
}
