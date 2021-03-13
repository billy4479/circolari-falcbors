package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func fetch() (err error) {
	err = os.MkdirAll("fetches", 0755)
	if err != nil {
		return
	}
	circolari, err := scrap()
	if err != nil {
		return
	}

	t := time.Now().Local()
	name := fmt.Sprintf("%d-%d-%d %d:%d:%d.json", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	file, err := os.Create("fetches/" + name)
	if err != nil {
		return
	}
	defer file.Close()
	j, err := json.Marshal(circolari)
	if err != nil {
		return
	}
	file.Write(j)

	err = os.RemoveAll("fetches/latest")
	if err != nil {
		return
	}

	err = os.Symlink(name, "fetches/latest")
	if err != nil {
		return
	}

	return
}
