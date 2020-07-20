package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type PlatformDiscovery struct {
	XMLName xml.Name `xml:"PlatformDiscovery"`
	Solutions []Solution `xml:"Solution"`
}

type Solution struct {
	XMLName xml.Name `xml:"Solution"`
	Name string `xml:"name,attr"`
	Exist string `xml:"exist,attr"`
	State string `xml:"state,attr"`
}

func ReadSolutions(reader io.Reader) ([]Solution, error) {
	var solutions PlatformDiscovery
	if err := xml.NewDecoder(reader).Decode(&solutions); err != nil {
		return nil, err
	}

	return solutions.Solutions, nil
}

func main() {
	out, err := exec.Command("PlatformDiscovery.exe").Output()
	if err != nil {
		log.Fatal(err)
	}
	platform := strings.NewReader(string(out))
	solutions, err := ReadSolutions(platform)
	fmt.Printf("Name %s, Exist %s, State %s", solutions[0].Name, solutions[0].Exist, solutions[0].State)
}
