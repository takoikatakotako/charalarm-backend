package logger

import "fmt"

func log(err error) {
	fmt.Println("----err--")
	fmt.Println(err)
	fmt.Println("-------")
}
