package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getInput(ask string) string {
	fmt.Print(ask)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func readData() Data {
	file, err := ioutil.ReadFile(".todo/tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	err2 := json.Unmarshal([]byte(file), &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return data
}

func writeData(data Data) {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile(".todo/tasks.json", content, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
}
