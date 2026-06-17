package main

import "testing"

func TestReadBannerFileInvalid(t *testing.T) {
    _, err := ReadBannerFile("wrongname")
    if err == nil {
        t.Errorf("expected error for invalid banner name, got nil")
    }
}

func TestReadBannerFileValid(t *testing.T) {
    bannerMap, err := ReadBannerFile("standard")
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
    if len(bannerMap) == 0 {
        t.Errorf("expected banner map to have entries, got empty map")
    }
}

func TestReadBannerFileMapSize(t *testing.T) {
    bannerMap, err := ReadBannerFile("standard")
    if err != nil {
        t.Fatal("could not read banner file")
    }
    if len(bannerMap) != 95 {
        t.Errorf("expected 95 characters, got %d", len(bannerMap))
    }
}

func TestSpaceCharacterHas8Lines(t *testing.T) {
    bannerMap, err := ReadBannerFile("standard")
    if err != nil {
        t.Fatal("could not read banner file")
    }
    if len(bannerMap[' ']) != 8 {
        t.Errorf("expected 8 lines, got %d", len(bannerMap[' ']))
    }
}

func TestReadBannerFileShadow(t *testing.T) {
    _, err := ReadBannerFile("shadow")
    if err != nil {
        t.Errorf("expected no error for shadow banner, got %v", err)
    }
}

func TestReadBannerFileThinkertoy(t *testing.T) {
    _, err := ReadBannerFile("thinkertoy")
    if err != nil {
        t.Errorf("expected no error for thinkertoy banner, got %v", err)
    }
}
