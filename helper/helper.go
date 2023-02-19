package helper

import "fmt"

// nama func dan variable jika diawali huruf besar bersifat Public
// jika diawali dengan huruf kecil bersifat private

var version = "1.1"
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello Mr. " + name)
}

func SayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
