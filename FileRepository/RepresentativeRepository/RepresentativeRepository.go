package RepresentativeRepository

import (
	"RegionCLI/Models"
	"RegionCLI/Models/Representatives"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

func CreateRepresentative(name, address, phoneNumber string, employeeCount, regionId int) (int, error) {
	r := Representatives.Representative{
		Id:            Models.CountLines(Representatives.RepresentativeFilePath) + 1,
		Name:          name,
		Address:       address,
		PhoneNumber:   phoneNumber,
		EmployeeCount: employeeCount,
		RegionId:      regionId,
		CreatedDate:   time.Now(),
	}
	var err error
	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func GetRepresentatives() (map[int]Representatives.Representative, error) {

	var err error
	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	representatives := make(map[int]Representatives.Representative)
	for scanner.Scan() {
		data := scanner.Bytes()
		var r Representatives.Representative
		err = json.Unmarshal(data, &r)
		if err != nil {
			return nil, err
		}
		representatives[r.Id] = r
	}
	return representatives, nil
}

func GetRepresentativeByID(id int) (Representatives.Representative, error) {
	var err error
	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var representative Representatives.Representative
		err := json.Unmarshal([]byte(line), &representative)
		if err != nil {
			log.Fatal(err)
		}
		if representative.Id == id {
			return representative, err
		}
	}

	return Representatives.Representative{-1, "", "", "", 0, 0, time.Now()}, err
}
