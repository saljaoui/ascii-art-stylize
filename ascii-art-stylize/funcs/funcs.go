package ascii

import (
	"fmt"
	"os"
	"strings"
)

func IsBanners(input string) bool {
	return input == "thinkertoy" || input == "standard" || input == "shadow" || input == "soufian"
}

func ScanInput(Input string) bool {
	Input = strings.ReplaceAll(Input, "\r\n", "")

	for _, char := range Input {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

func GetAsciiArt(char int, fileContent string) (string, error) {
	fileContent = strings.ReplaceAll(fileContent, "\r\n", "\n")

	asciiOffset := char - 32
	totalLines := strings.Count(fileContent, "\n")

	if totalLines != 855 {
		return "", fmt.Errorf("the file is not true")
	}

	fileContent = fileContent[1:]
	lines := strings.Split(fileContent, "\n\n")
	return lines[asciiOffset], nil
}

func ParseInput(input string) (string, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return "", fmt.Errorf("input contains invalid characters")
		}
	}
	return input, nil
}

func RenderAscii(lines [][]string) string {
	str := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				str += (lines[j][i])
			}
		}
		str += ("\n")
	}
	return str
}

func GenerateAsciiArt(input, banner string) (string, error) {
	var strs string
	var nStr [][]string

	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		line, err = ParseInput(line)
		if err != nil {
			return "", err
		}
		for _, char := range line {

			asciiStr, err := GetAsciiArt(int(char), string(content))
			if err != nil {
				return "", err
			}
			nStr = append(nStr, strings.Split(asciiStr, "\n"))
		}
		strs += RenderAscii(nStr)
		nStr = nil
	}
	return strs, nil
}
