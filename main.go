package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("#### URL TEST ####")

	var file []byte
	var err error
	file, err = os.ReadFile("urls.txt")

	if err != nil {
		fmt.Println(err)
	}
	var lines = strings.Split(string(file), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
