package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var input, banner string
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		return
	} else if len(os.Args) == 2 {
		input = os.Args[1]
		banner = "standard"
	} else if len(os.Args) == 3 {
		input = os.Args[1]
		banner = os.Args[2]
	}

	bannerMap, err := readBannerFile(banner)
	if err != nil {
		fmt.Println("Error: invalid banner name")
		return
	}
	inputLines := strings.Split(input, "\\n")
	for _, line := range inputLines {
		if line == "" {
			fmt.Println()
			continue
		}
		fmt.Print(renderLine(line, bannerMap))
		fmt.Println()
	}
}
