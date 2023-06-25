package main

import (
	"RegionCLI/FileRepository/RepresentativeRepository"
	"fmt"
)

func main() {
	//for {
	//	flag.Parse()
	//	commands := flag.Args()
	//	for i, c := range commands {
	//		switch c {
	//		case "Region":
	//			fmt.Println(RegionRepository.GetRegions())
	//		case "list":
	//			fmt.Println()
	//		}
	//	}
	//}
	n, err := RepresentativeRepository.GetRepresentativesByRegionId(3)
	fmt.Println(n, err)
}
