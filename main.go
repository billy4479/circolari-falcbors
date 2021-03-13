package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	mode := flag.Int("mode", -1, "1 = fetch; 2 = fetch and diff with last fetch")
	flag.Parse()

	if *mode == 1 {
		err := fetch()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if *mode == 2 {
		err := fetchAndDiff()
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				log.Fatal("First run with mode 1")
			}
			log.Fatal(err)
		}
	}

	if *mode == -1 {
		fmt.Println("You must specify a -mode")
	}
	return
}
