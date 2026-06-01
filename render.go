package main

func renderLine(line string, bannermap map[rune][]string) string {
	result := ""
	for i := 0; i < 8; i++ {
		for _, ch := range line {
			result += bannermap[ch][i]
		}
		result += "\n"
	}
	return result
}
