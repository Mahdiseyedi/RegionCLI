package main

import (
	"RegionCLI/FileRepository/RepresentativeRepository"
	"fmt"
)

func main() {
	n, err := RepresentativeRepository.StatusRepresentative(3)
	//n, err := RepresentativeRepository.GetRepresentatives()
	fmt.Println(n, err)
}
