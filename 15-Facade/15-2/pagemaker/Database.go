package pagemaker

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getProperties(dbname string) map[string]string {
	filename := dbname + ".txt"
	prop := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("File %v not found\n", filename))
	}
	defer func() {
		if r := recover(); r != nil {
			err := errors.New(fmt.Sprintf("recovered: %v\n", r))
			fmt.Printf("エラー処理をサボりました %v", err)
		}
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, "=")
		if len(slice) != 2 {
			panic(fmt.Sprintf("%v is invalid format", line))
		}

		prop[slice[0]] = slice[1]
	}

	return prop
}
