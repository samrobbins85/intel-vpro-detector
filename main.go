package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("PlatformDiscovery.exe").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}
