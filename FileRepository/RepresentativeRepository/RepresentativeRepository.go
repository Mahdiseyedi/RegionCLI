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

func EditRepresentative(id int, newName, newAddress, newPhoneNumber string, newEmployeeCount, newRegionId int) error {
	var updatedRepresentatives []Representatives.Representative
	var err error

	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_APPEND|os.O_CREATE, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var Representative Representatives.Representative
		err := json.Unmarshal([]byte(line), &Representative)
		if err != nil {
			log.Fatal(err)
		}
		if Representative.Id == id {
			lastRepresentative, _ := GetRepresentativeByID(id)
			if newName != "" {
				Representative.Name = newName
			} else {
				Representative.Name = lastRepresentative.Name
			}
			if newAddress != "" {
				Representative.Address = newAddress
			} else {
				Representative.Address = lastRepresentative.Address
			}
			if newPhoneNumber != "" {
				Representative.PhoneNumber = newPhoneNumber
			} else {
				Representative.PhoneNumber = lastRepresentative.PhoneNumber
			}
			Representative.EmployeeCount = newEmployeeCount
			Representative.RegionId = newRegionId

			Representative.CreatedDate = lastRepresentative.CreatedDate
		}
		updatedRepresentatives = append(updatedRepresentatives, Representative)
	}

	f2, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f2.Close()

	for _, Representative := range updatedRepresentatives {
		newData, err := json.Marshal(Representative)
		if err != nil {
			log.Fatal(err)
		}
		f2.WriteString(string(newData) + "\n")
	}
	return err
}

func DeleteRepresentative(id int) (int, error) {
	var updatedRepresentatives []Representatives.Representative
	var err error

	if _, sErr := os.Stat(Representatives.RepresentativeFilePath); os.IsNotExist(err) {
		err = sErr
		return -1, err
	}
	f, oErr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_RDONLY, 0644)
	if oErr != nil {
		err = oErr
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var representative Representatives.Representative
		jErr := json.Unmarshal([]byte(line), &representative)
		if jErr != nil {
			err = jErr
			return -1, err
		}
		if representative.Id != id {
			updatedRepresentatives = append(updatedRepresentatives, representative)
		}
	}

	f2, oFerr := os.OpenFile(Representatives.RepresentativeFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if oFerr != nil {
		err = oFerr
		return -1, err
	}
	defer f2.Close()

	for _, representative := range updatedRepresentatives {
		newData, jMerr := json.Marshal(representative)
		if jMerr != nil {
			return -1, jMerr
		}
		f2.WriteString(string(newData) + "\n")
	}

	return 0, err
}
