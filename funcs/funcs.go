package ascii

import (
	"fmt"
	"os"
	"strings"
)
// IsBanners checks if the input string matches any of the predefined banner types.
func IsBanners(input string) bool {
	return input == "thinkertoy" || input == "standard" || input == "shadow" || input == "soufian"
}
// ScanInput checks if the input string contains only printable ASCII characters.
func ScanInput(Input string) bool {
	for _, char := range Input {
		if (char < 32 || char > 126) && char != 13 && char != 10 {
			return false
		}
	}
	return true
}

// GetAsciiArt retrieves ASCII art for a specific character from fileContent.
func GetAsciiArt(char int, fileContent string) (string, error) {
	fileContent = strings.ReplaceAll(fileContent, "\r\n", "\n")
	asciiOffset := char - 32
	fileContent = fileContent[1:]
	lines := strings.Split(fileContent, "\n\n")
	if len(lines) != 95 {
		return "", fmt.Errorf("the file is not true")
	}
	return lines[asciiOffset], nil
}

// RenderAscii renders ASCII art lines into a single string.
func RenderAscii(lines [][]string) string {
	var str strings.Builder
	for i := 0; i < 8; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				str.WriteString(lines[j][i])
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}

// GenerateAsciiArt generates ASCII art for input using the specified banner.
func GenerateAsciiArt(input, banner string) (string, error) {
	var strs strings.Builder
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		var nStr [][]string
		for _, char := range line {
			asciiStr, err := GetAsciiArt(int(char), string(content))
			if err != nil {
				return "", err
			}
			nStr = append(nStr, strings.Split(asciiStr, "\n"))
		}
		strs.WriteString(RenderAscii(nStr))
	}
	return strs.String(), nil
}
