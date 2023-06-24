package RegionRepository

import (
	"RegionCLI/Models"
	"RegionCLI/Models/Regions"
	"bufio"
	"encoding/json"
	"log"
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
	scanner := bufio.NewScanner(f)
	regions := make(map[int]Regions.Region)
	for scanner.Scan() {
		data := scanner.Bytes()
		var r Regions.Region
		err = json.Unmarshal(data, &r)
		if err != nil {
			return nil, err
		}
		regions[r.Id] = r
	}
	return regions, nil
}

func GetRegionByID(id int) (Regions.Region, error) {
	var err error
	f, oErr := os.OpenFile(RegionFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var region Regions.Region
		err := json.Unmarshal([]byte(line), &region)
		if err != nil {
			log.Fatal(err)
		}
		if region.Id == id {
			return region, err
		}
	}

	return Regions.Region{-1, ""}, err
}

func EditRegion(id int, newName string) error {
	var updatedRegions []Regions.Region
	var err error

	f, oErr := os.OpenFile(RegionFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var region Regions.Region
		err := json.Unmarshal([]byte(line), &region)
		if err != nil {
			log.Fatal(err)
		}
		if region.Id == id {
			region.Name = newName
		}
		updatedRegions = append(updatedRegions, region)
	}

	f2, oErr := os.OpenFile(RegionFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f2.Close()

	for _, region := range updatedRegions {
		newData, err := json.Marshal(region)
		if err != nil {
			log.Fatal(err)
		}
		f2.WriteString(string(newData) + "\n")
	}
	return err
}

func DeleteRegion(id int) (int, error) {
	var updatedRegions []Regions.Region
	var err error

	if _, sErr := os.Stat(RegionFilePath); os.IsNotExist(err) {
		err = sErr
		return -1, err
	}

	f, oErr := os.OpenFile(RegionFilePath, os.O_RDONLY, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var region Regions.Region
		jErr := json.Unmarshal([]byte(line), &region)
		if jErr != nil {
			err = jErr
			return -1, err
		}
		if region.Id != id {
			updatedRegions = append(updatedRegions, region)
		}
	}

	f2, oFerr := os.OpenFile(RegionFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if oFerr != nil {
		err = oFerr
		return -1, err
	}
	defer f2.Close()

	for _, region := range updatedRegions {
		newData, jMerr := json.Marshal(region)
		if jMerr != nil {
			return -1, jMerr
		}
		f2.WriteString(string(newData) + "\n")
	}

	return 0, err
}
