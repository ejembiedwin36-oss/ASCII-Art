package main

import "testing"

// func TestRenderLine(t *testing.T) {
// 	//1. load the banner map
// 	bannerMap, err := readBannerFile("standard")
// 	if err != nil {
// 		t.Fatal("could not read banner file")
// 	}

// 	//2.  call renderLine with a space

// 	result := renderLine(" ", bannerMap)

// 	// 3. check the result is not empty
// 	if result == "" {
// 		t.Errorf("expectet output, got empty string")
// 	}
// }

func TestRenderLine(t *testing.T) {
	bannerMap, err := readBannerFile("standard")
	if err != nil {
		t.Fatal("could not read banner file")

	}

	expected := "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           \n"
	result := renderLine("A", bannerMap)

	if result != expected {
		t.Errorf("got:\n%s\nwant\n%s", result, expected)
	}
}
