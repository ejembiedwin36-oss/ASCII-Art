package main

import (
	"os"
	"strings"
)

func readBannerFile(banner string) (map[rune][]string, error) {
	filename := "banner-files/" + banner + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(data)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	bannermap := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(32 + i)
		start := 1 + i*9
		bannermap[char] = lines[start : start+8]
	}
	return bannermap, nil

}
