package main

import (
	"flag"
	"fmt"

	"nullprogram.com/x/uuid"
)

func main() {
	n := flag.Int("n", 1, "number of UUIDs to generate")
	flag.Parse()

	g := uuid.NewGen()
	for i := 0; i < *n; i++ {
		fmt.Println(g.NewV4())
	}
}
