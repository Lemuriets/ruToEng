package io

import (
	"log"
	"os"
)

func WriteFile(name, text string) {
	file, err := os.Create(name)
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(text))
	if err != nil {
		log.Fatal(err)
	}
}
