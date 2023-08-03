package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 3; i++ {
		fmt.Print("Write ", i, " as ")
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before nooon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("i'm an Int")
		default:
			fmt.Printf("IDK that type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

/*
Closures:
A closure is a function value that references variables from outside its body.
In Go, anonymous functions can form closures when they access variables defined outside their own scope.
The anonymous function "closes over" the variables it uses from its surrounding lexical scope.
This means the function retains access to those variables even after the outer function has finished executing.
Closures are useful when you want to capture the state of a function at a specific point in time.

```golang
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // Output: 1
    fmt.Println(c()) // Output: 2
    fmt.Println(c()) // Output: 3
}

```


*/
