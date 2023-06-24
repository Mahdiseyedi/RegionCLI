package RepresentativeRepository

import (
	"RegionCLI/Models"
	"RegionCLI/Models/Regions"
	"encoding/json"
	"os"
)

func CreateRegion(regionName string) (int, error) {
	r := Regions.Region{
		Id:   Models.CountLines(Regions.RegionFilePath) + 1,
		Name: regionName,
	}
	var err error
	f, oErr := os.OpenFile(Regions.RegionFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
