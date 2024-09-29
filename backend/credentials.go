package backend

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (a *App) GetValue(line int, filename string) string {
	path := fmt.Sprintf("/Users/nevinod/.aws/%s", filename)

	readFile, err := os.Open(path)

	if err != nil {
		return ""
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	currentLine := 0
	for fileScanner.Scan() {
		if currentLine == line {
			rawValue := strings.Split(fileScanner.Text(), "=")[1]
			value := strings.TrimSpace(rawValue)
			return value
		}
		currentLine += 1
	}

	readFile.Close()
	return ""
}
