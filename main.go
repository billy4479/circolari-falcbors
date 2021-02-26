package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	old, err := scrap()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Hour)
	new, err := scrap()
	if err != nil {
		log.Fatal(err)
	}

	d := getDiff(new, old)
	j, _ := json.Marshal(d)
	fmt.Println(string(j))
}
