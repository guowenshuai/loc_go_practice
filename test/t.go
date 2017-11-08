package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
