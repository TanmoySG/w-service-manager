package main

import (
	"fmt"

	"validity/pkg/kafka"
)

func main()  {

	kc := kafka.Client{}
	kc.Producer()
	fmt.Print("test")
}