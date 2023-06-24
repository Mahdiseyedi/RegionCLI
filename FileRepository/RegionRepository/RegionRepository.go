package RegionRepository

import (
	"RegionCLI/Models"
	"RegionCLI/Models/Regions"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	RegionFilePath = "./Regions.txt"
)

func CreateRegion(regionName string) (int, error) {
	r := Regions.Region{
		Id:   Models.CountLines(RegionFilePath) + 1,
		Name: regionName,
	}
	var err error
	f, oErr := os.OpenFile(RegionFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func GetRegions() (map[int]Regions.Region, error) {

	var err error
	f, oErr := os.OpenFile(RegionFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()
	fmt.Println("scan started")
	scanner := bufio.NewScanner(f)
	regions := make(map[int]Regions.Region)
	fmt.Println("scan started before for")
	for scanner.Scan() {
		data := scanner.Bytes()
		var r Regions.Region
		err = json.Unmarshal(data, &r)
		if err != nil {
			return nil, err
		}
		fmt.Println(r)
		regions[r.Id] = r
	}
	fmt.Println("scan before return")
	fmt.Println(regions)
	return regions, nil
}
