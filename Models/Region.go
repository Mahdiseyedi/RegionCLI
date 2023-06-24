package Models

import (
	"bufio"
	"encoding/json"
	"os"
)

const (
	regionFilePath = "./Regions.txt"
)

type region struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func CreateRegion(regionName string) (region, error) {
	r := region{
		Id:   CountLines() + 1,
		Name: regionName,
	}
	var err error
	f, oErr := os.OpenFile(regionFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	eErr := enc.Encode(r)
	if eErr != nil {
		err = eErr
	}
	return r, err
}

func CountLines() int {
	f, err := os.Open(regionFilePath)
	if err != nil {
		return 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}

	return count
}
