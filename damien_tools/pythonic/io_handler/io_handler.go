package io_handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type File struct {
	osFile      *os.File
	bufioReader *bufio.Reader
	bufioWriter *bufio.Writer
}

func Input(promptStrings ...string) (str string) {
	var (
		err error
	)

	for _, promptString := range promptStrings {
		fmt.Print(promptString)
	}

	reader := bufio.NewReader(os.Stdin)

	if str, err = reader.ReadString('\n'); err != nil {
		log.Fatalln("<func Input(promptStrings ...string) (str string)> - \"reader.ReadString('\\n')\" failed in reading string from \"os.Stdin\":", err)
	}

	str = string(bytes.TrimRight([]byte(str), "\n"))
	str = string(bytes.TrimRight([]byte(str), "\r"))

	return
}

func SelectionListInput(selectionList []string, startFrom int) (str string) {
	var (
		err             error
		selectionNumber int
	)

	if len(selectionList) == 0 {
		log.Fatalln("<func SelectionListInput(selectionList []string, startFrom int) (str string)> - The length of \"selectionList []string\" can not be zero!")
	} else if startFrom < 0 {
		log.Fatalln("<func SelectionListInput(selectionList []string, startFrom int) (str string)> - The parameter \"startFrom int\" needs to be non-negative integer!")
	}

	endWith := startFrom
    numberItemMap := make(map[int]string)

    for _, item := range selectionList {
        fmt.Println("[" + strconv.Itoa(endWith) + "]: " + item);
        numberItemMap[endWith] = item;
        endWith++;
    }

	for selectionNumber, err = strconv.Atoi(Input("Please enter your number to choose: ")); selectionNumber < startFrom || selectionNumber >= endWith || err != nil; {
		if err != nil {
			log.Println("<strconv.Atoi(Input(\"...\"))> - Function failed:", err)
		}

		selectionNumber, err = strconv.Atoi(Input("Invalid input! Please enter your number again: "))
	}

    str = numberItemMap[selectionNumber]

	return
}

// Open only supports file mode "r", "w" and "a".
func Open(fileName string, fileMode string) (file *File) {
	switch fileMode {
	case "r":
		if osFile, err := os.Open(fileName); err != nil {
			log.Fatalln("<os.Open(fileName)> - Failed in opening file \"" + fileName + "\" with file mode \"" + fileMode + "\":", err)
		} else {
			file = &File{osFile: osFile, bufioReader: bufio.NewReader(osFile), bufioWriter: nil}
		}
	case "w":
		if osFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0); err != nil {
			log.Fatalln("<os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0)> - Failed in opening file \"" + fileName + "\" with file mode \"" + fileMode + "\":", err)
		} else {
			file = &File{osFile: osFile, bufioReader: nil, bufioWriter: bufio.NewWriter(osFile)}
		}
	case "a":
		if osFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0); err != nil {
			log.Fatalln("<os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0)> - Failed in opening file \"" + fileName + "\" with file mode \"" + fileMode + "\":", err)
		} else {
			file = &File{osFile: osFile, bufioReader: nil, bufioWriter: bufio.NewWriter(osFile)}
		}
	default:
		log.Fatalln("<func Open(fileName string, fileMode string) (file *File)> - Wrong file mode \"" + fileMode + "\": Open only supports file mode \"r\", \"w\" and \"a\".")
	}

	return
}

func (this *File) ReadLine() (str string) {
	if this.bufioReader != nil {
		var (
			err error
		)

		if str, err = this.bufioReader.ReadString('\n'); err != nil && err != io.EOF {
			log.Fatalln("<this.bufioReader.ReadString('\\n')> - Failed in reading one line from file \"" + this.osFile.Name() + "\":", err)
		}
	} else {
		log.Println("<func (this *File) ReadLine() (str string)> - Cannot read from file \"" + this.osFile.Name() + "\": No read permission.")
	}

	return
}

func (this *File) Write(str string) {
	if this.bufioWriter != nil {
		if _, err := this.bufioWriter.WriteString(str); err != nil {
			log.Fatalln("<this.bufioWriter.WriteString(str)> - Failed in writing \"str string\" into file \"" + this.osFile.Name() + "\":", err)
		}
	} else {
		log.Println("<func (this *File) Write(str string)> - Cannot write into file \"" + this.osFile.Name() + "\": No write permission.")
	}
}

func (this *File) Close() {
	if this.bufioWriter != nil {
		if err := this.bufioWriter.Flush(); err != nil {
			log.Fatalln("<this.bufioWriter.Flush()> - Failed in flushing \"bufioWriter\":", err)
		}
	}

	if err := this.osFile.Close(); err != nil {
		log.Fatalln("<this.osFile.Close()> - Failed in closing file \"" + this.osFile.Name() + "\":", err)
	}
}
