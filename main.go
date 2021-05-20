package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	loop := 0
	for {
		fmt.Printf("Hello %s, loop: %v \n", os.Getenv("MY_VARIABLE"), loop) //MY_VARIABLE is defined in .devcontainer/devcontainer.json
		time.Sleep(time.Second * 2)
		loop++
	}
}
