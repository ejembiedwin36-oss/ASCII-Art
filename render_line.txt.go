package main

import "testing"

func TestRenderLine(t *testing.T) {
    bannerMap, err := ReadBannerFile("standard")
    if err != nil {
        t.Fatal("could not read banner file")
    }
    result := RenderLine(" ", bannerMap)
    if result == "" {
        t.Errorf("expected output, got empty string")
    }
}

func TestRenderLineA(t *testing.T) {
    bannerMap, err := ReadBannerFile("standard")
    if err != nil {
        t.Fatal("could not read banner file")
    }
    expected := "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           \n"
    result := RenderLine("A", bannerMap)
    if result != expected {
        t.Errorf("got:\n%s\nwant:\n%s", result, expected)
    }
}
