package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetIdiomFromFile(file string) ([]string, error) {
	fi, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error", err)
		return nil, err
	}
	defer fi.Close()

	idioms := make([]string, 0)
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read line error", err)
			return nil, err
		}

		idioms = append(idioms, string(line))
	}

	return idioms, nil
}
