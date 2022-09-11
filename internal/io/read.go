package io

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)

	scannedText, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return scannedText
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
