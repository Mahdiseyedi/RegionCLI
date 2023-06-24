package Regions

import (
	"RegionCLI/Models"
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

func CreateRegion(regionName string) (int, error) {
	r := region{
		Id:   Models.CountLines(regionFilePath) + 1,
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
	return r.Id, err
}
