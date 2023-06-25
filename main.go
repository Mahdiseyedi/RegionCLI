package main

import (
	"RegionCLI/FileRepository/RegionRepository"
	"RegionCLI/FileRepository/RepresentativeRepository"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	commands := flag.Args()

	switch commands[0] {
	case "region":
		res, err := RegionRepository.GetRegions()
		if err == nil {
			for _, reg := range res {
				fmt.Println(reg.Name)
			}
		} else {
			fmt.Println(err)
		}
	case "list":
		if len(commands) < 2 {
			fmt.Println("Something went wrong try again later ...")
			break
		}
		regionId, _ := strconv.Atoi(commands[1])
		res, err := RepresentativeRepository.GetRepresentativesByRegionId(regionId)
		if err == nil {
			for _, r := range res {
				fmt.Println(r.Name)
			}
		} else {
			fmt.Println(err)
		}
	case "get":
		if len(commands) < 2 {
			fmt.Println("Something went wrong try again later ...")
			break
		}
		representativeId, _ := strconv.Atoi(commands[1])
		res, err := RepresentativeRepository.GetRepresentativeByID(representativeId)

		if err == nil {
			data, err1 := json.Marshal(res)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				jsonString := string(data)
				fmt.Println(jsonString)
			}
		} else {
			fmt.Println(err)
		}
	case "create":
		if len(commands) < 6 {
			fmt.Println("Something went wrong try again later ...")
			break
		}
		rName := commands[1]
		rAddress := commands[2]
		rPhoneNumber := commands[3]
		rEmployeeCount, _ := strconv.Atoi(commands[4])
		rRegionId, err1 := RegionRepository.GetRegionIdByName(commands[5])
		res, err2 := RepresentativeRepository.CreateRepresentative(rName, rAddress, rPhoneNumber, rEmployeeCount, rRegionId)
		if err1 == nil && err2 == nil {
			fmt.Println(res)
		} else if err1 != nil {
			fmt.Println(err1)
		} else if err2 != nil {
			fmt.Println(err2)
		}
	case "edit":
		if len(commands) < 7 {
			fmt.Println("Something went wrong try again later ...")
			break
		}

		repId, _ := strconv.Atoi(commands[1])
		rName := commands[2]
		rAddress := commands[3]
		rPhoneNumber := commands[4]
		rEmployeeCount, _ := strconv.Atoi(commands[5])
		rRegionId, err2 := RegionRepository.GetRegionIdByName(commands[6])
		err1 := RepresentativeRepository.EditRepresentative(repId, rName, rAddress, rPhoneNumber, rEmployeeCount, rRegionId)

		if err1 == nil && err2 == nil {
			fmt.Println("Edited Successfully !")
		} else if err2 != nil {
			fmt.Println(err2)
		} else if err1 != nil {
			fmt.Println(err1)
		}
	case "status":
		if len(commands) < 2 {
			fmt.Println("Something went wrong try again later ...")
			break
		}
		regionId, _ := RegionRepository.GetRegionIdByName(commands[1])
		rRepresentativeCount, rEmployeeCount, err := RepresentativeRepository.StatusRepresentative(regionId)
		if err == nil {
			fmt.Println("Count of Employee: ", rEmployeeCount)
			fmt.Println("Count of Representative: ", rRepresentativeCount)
		} else {
			fmt.Println(err)
		}
	default:
		fmt.Println("Wrong input !")
	}
}
