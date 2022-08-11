package utils

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

type targetList []string

// File file struct
type File struct {
	FilePath string
}

// ReadFile readfile and return content
func (f *File) ReadFile() targetList {
	var content targetList
	fileObj, err := os.OpenFile(f.FilePath, os.O_RDONLY, 0666)
	if err != nil {
		gologger.Error().Msg("Open target file error")
		return nil
	}
	defer fileObj.Close()

	buf := bufio.NewReader(fileObj)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		content = append(content, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				gologger.Error().Msg("Read target file error")
				return nil
			}
		}
	}
	return content
}
