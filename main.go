package main

import "github.com/natashaCarreao/go-multithreading/cmd"

func main() {
	err := cmd.Initialize()
	if err != nil {
		panic(err)
	}
}
