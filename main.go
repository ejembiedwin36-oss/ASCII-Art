package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	var input, banner string
	if len(os.Args)< 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		return
	}else if len(os.Args)== 2 {
		input = os.Args[1]
		banner = "standard"
	}else if len(os.Args)== 3 {
		input = os.Args[1]
		banner = os.Args[2]
	}

	bannermap, err := ReadBannerFile(banner)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	inputlines := strings.Split(input, "\\n")
	for _, line := range inputlines {
		if line == "" {
			fmt.Println()
			continue
		}
		fmt.Print(RenderLine(line, bannermap))
	}

}
