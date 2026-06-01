package main

import "testing"


// func TeastReadBannerFileInvalid(t *testing.T) {
// 	_, err := readBannerFile("wrongname")
// 	if err != nil {
// 		t.Errorf("expected error for invalid banner name, got nil")
// 	}
// }

func TeastReadBannerFileInvalid(t *testing.T) {
	bannerMap, err := readBannerFile("standard")
	if err != nil {
		t.Errorf("expected no error, got %v", err)

	}
	if len(bannerMap) == 0 {
		t.Errorf("expected banner map to have entries, got map")
	}
	
}