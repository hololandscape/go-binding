package pkg

// #include "greeter.h"
import "C"
import "fmt"

func greeter() {
	var name string = "aisuko"
	var year int = 2023
	var message [1000]C.char
	C.greet(C.CString(name), C.int(year), &message[0])
	fmt.Println(C.GoString(&message[0]))
}
