package pa_parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ParseDirectory(path string) {
	fmt.Println(path)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		parseFile(path)
		return nil
	})
}

func parseFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parseLine(scanner.Text())
	}

}

func parseLine(line string) {

}
