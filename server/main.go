package main

import (
	"boilerplate/env"
	"fmt"
	"os"
)

func init() {
	env.LoadEnvs()
}

func main() {

	fmt.Printf("\n\n####################################\n######### Version %s ############\n####################################\n\n", os.Getenv("APP_VERSION"))
	fmt.Println("Hello World!")
}
