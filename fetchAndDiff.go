package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func fetchAndDiff() (err error) {
	oldj, err := os.ReadFile("fetches/latest")
	if err != nil {
		return
	}
	var old []circolare
	err = json.Unmarshal(oldj, &old)
	if err != nil {
		return
	}

	err = fetch()
	if err != nil {
		return
	}

	newj, err := ioutil.ReadFile("fetches/latest")
	if err != nil {
		return
	}
	var new []circolare
	err = json.Unmarshal(newj, &new)
	if err != nil {
		return
	}

	diff := getDiff(new, old)
	err = os.MkdirAll("diffs", 0755)
	if err != nil {
		return err
	}

	t := time.Now().Local()
	name := fmt.Sprintf("%d-%d-%d %d:%d:%d.json", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	file, err := os.Create("diffs/" + name)
	if err != nil {
		return
	}
	defer file.Close()

	diffj, err := json.Marshal(diff)
	if err != nil {
		return
	}

	file.Write(diffj)

	return
}
