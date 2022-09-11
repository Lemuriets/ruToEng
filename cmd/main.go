package main

import (
	"flag"
	"fmt"

	"github.com/Lemuriets/ruToEng/internal/io"
	"github.com/Lemuriets/ruToEng/internal/stop"
	"github.com/Lemuriets/ruToEng/pkg/ruToEng"
)

func main() {
	var text string

	f := flag.String("f", "", "-f=filename")
	wf := flag.String("wf", "", "-wf=<filename>")
	flag.Parse()

	filePath := *f
	if filePath != "" {
		text = io.ReadFile(filePath)
	} else {
		text = io.ReadString()
	}

	text = ruToEng.RuToEng(text)

	fileToWrite := *wf
	if fileToWrite != "" {
		io.WriteFile(fileToWrite, text)
	} else {
		fmt.Print(text)
		stop.StopBtInterrup()
	}
}
