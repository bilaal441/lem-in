package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetData(filepath string) (int, []string, []string) {
	data, err := os.ReadFile(filepath)
	links := []string{}
	rooms := []string{}
	if err != nil {
		log.Fatal(data)
	}
	// fileData:= strings.Split(string(data), "\n")

	for _, curr := range strings.Split(string(data), "\n") {

		if strings.Contains(curr, "-") {
			links = append(links, curr)
		} else {
			rooms = append(rooms, curr)
		}

	}
	numAnt, _ := strconv.Atoi(strings.Split(string(data), "\n")[0])
	return numAnt, rooms, links
}
