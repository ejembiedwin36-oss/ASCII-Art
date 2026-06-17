# ASCII Art Generator

A command-line program written in **Go** that converts text input into ASCII art using banner files.

---

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Usage](#usage)
- [Examples](#examples)
- [Banner Styles](#banner-styles)

---

## Overview

ASCII Art Generator takes a string as a command-line argument and outputs it as a graphic representation using ASCII characters. It supports three banner styles and handles special characters, numbers, spaces, and newlines.

---

## Project Structure

```
ascii-art/
├── main.go           # Entry point — handles arguments and coordinates the program
├── readbanner.go     # Loads and parses the banner file into a character map
├── render.go         # Generates ASCII art for a line of text
└── banner-files/
    ├── standard.txt  # Standard banner style
    ├── shadow.txt    # Shadow banner style
    └── thinkertoy.txt # Thinkertoy banner style
```

---

## How It Works

The program is divided into three clean functions:

### `main.go`
- Validates command-line arguments
- Calls `readBannerFile` to load the banner data
- Splits the input by `\n` to handle multi-line input
- Calls `renderLine` for each line and prints the result

### `readbanner.go`
- Builds the file path from the banner name
- Reads the banner file from disk
- Fixes Windows-style line endings (`\r\n` → `\n`)
- Splits the file content into lines
- Builds a `map[rune][]string` where each character maps to its 8 lines of ASCII art
- Returns the map for use by `renderLine`

### `render.go`
- Takes a line of text and the banner map as input
- Loops through 8 rows and all characters simultaneously
- Looks up each character's ASCII art in the map
- Returns the complete ASCII art as a string

---

## Usage

```bash
# No arguments — shows usage message
go run . 

# One argument — uses standard banner by default
go run . "Hello"

# Two arguments — uses specified banner
go run . "Hello" shadow
go run . "Hello" standard
go run . "Hello" thinkertoy
```

### Handling Newlines

Use `\n` inside your string to print on separate lines:

```bash
go run . "Hello\nThere"
go run . "Hello\n\nThere"
```

---

## Examples

### Standard Banner
```bash
go run . "Hello"
```
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

### Shadow Banner
```bash
go run . "Hello" shadow
```
```
                                 
_|    _|          _| _|          
_|    _|   _|_|   _| _|   _|_|   
_|_|_|_| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
                                 
```

### Thinkertoy Banner
```bash
go run . "Hello" thinkertoy
```
```
                 
o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                 
                 
```

---

## Banner Styles

| Banner | Description |
|--------|-------------|
| `standard` | Clean, bold style using `_`, `|`, `/`, `\` characters |
| `shadow` | Shadow effect style using `_\|` characters |
| `thinkertoy` | Rounded, playful style using `o`, `O`, `-` characters |

---

## Requirements

- Go 1.18 or higher

## Author

Edwin
