package gateways

import (
	"bufio"
	"os"
)

type File struct {
}

func (f File) GetLines() ([]string, error) {
	quakeLog, err := os.Open("./quake.log")
	if err != nil {
		panic(err)
	}

	defer quakeLog.Close()

	scanner := bufio.NewScanner(quakeLog)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func NewFile() File {
	return File{}
}
