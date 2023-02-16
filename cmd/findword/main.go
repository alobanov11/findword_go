package main

import (
	"findword/internal/pkg/app"
	"flag"
	"log"
)

var (
	word = flag.String("w", "go", "the word is necessary to find")
	tops = flag.Int("t", 5, "the number of max concurrent tasks")
)

func main() {
	flag.Parse()

	a, err := app.New(*word, *tops)

	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
