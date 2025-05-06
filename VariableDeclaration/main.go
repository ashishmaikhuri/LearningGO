package main

import "fmt"

var isLoggedIn1 bool = true
var intVar uint = 255

// Variable and Func with Capital Name can be exported out of the Package
var FloatVar2 float32 = 255.432423423

// Implicit declaration. Lexer will auto detect the type of variable.
var username2 = "Hello New String Variable"

// Declaring constants
const (
	Username    = "Global String Variable" // Public Constant Variable
	isLoggedIn2 = true
)

func main() {
	// No var style. Only works inside the function. It's (:=) called walrus //operator.
	username3 := "Hello New String Variable"
	fmt.Println(username3)

}
