package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func urlCheck(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	var _, err = http.Head(url)
	if err != nil {
		fmt.Println("[FAIL]", url)
	} else {
		fmt.Println("[OK]", url)
	}
}
func main() {
	var start = time.Now()
	var wg sync.WaitGroup

	fmt.Println("#### URL TEST ####")

	var file []byte
	var err error
	file, err = os.ReadFile("urls.txt")

	if err != nil {
		fmt.Println(err)
	}
	var lines = strings.Split(string(file), "\n")
	wg.Add(len(lines))
	for _, line := range lines {
		go urlCheck(line, &wg)
	}
	wg.Wait()
	fmt.Printf("Elapsed Time - %v", time.Since(start).String())
}
