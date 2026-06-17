package main  


func RenderLine(line string, bannermap map[rune][]string) string {
	result := ""

	for i := 0; i < 8; i++ {
		for _, ch := range line {
			if ch < 32 || ch > 126 {
				return "unformatted string"
			}
			result += bannermap[ch][i]
		}
		result += "\n"
	}
	return  result
}
