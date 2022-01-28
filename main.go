package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(filename string, t *string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	buf := make([]byte, 256)

	for {
		n, err := file.Read(buf)
		if n > 0 {
			*t = string(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != io.EOF && err != nil {
			log.Fatalln(err)
		}
	}
}

func writeData(filename string, t *string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	s := strings.Split(*t, " ")

	for _, item := range s {
		if _, err := strconv.Atoi(item); err == nil {
			file.WriteString(item + " ")
		}
	}
	io.Copy(os.Stdout, file)
}

func main() {

	var text string = ""
	readData("in.txt", &text)
	writeData("out.txt", &text)
}
