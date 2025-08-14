package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("#### URL TEST ####")
	var file, err = os.ReadFile("urls.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))

}
